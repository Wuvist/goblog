package main

import (
	"archive/zip"
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/labstack/echo"
)

func TestZipImageServerHandler(t *testing.T) {
	tmpDir := t.TempDir()
	archivePath := filepath.Join(tmpDir, "images.zip")

	content := []byte("hello world")
	modTime := time.Date(2023, 10, 5, 12, 0, 0, 0, time.UTC)

	zipFile, err := os.Create(archivePath)
	if err != nil {
		t.Fatalf("create zip: %v", err)
	}

	writer := zip.NewWriter(zipFile)
	header := &zip.FileHeader{
		Name:   "640/test.jpg",
		Method: zip.Deflate,
	}
	header.SetModTime(modTime)
	entryWriter, err := writer.CreateHeader(header)
	if err != nil {
		t.Fatalf("create header: %v", err)
	}
	if _, err := entryWriter.Write(content); err != nil {
		t.Fatalf("write entry: %v", err)
	}
	if err := writer.Close(); err != nil {
		t.Fatalf("close writer: %v", err)
	}
	if err := zipFile.Close(); err != nil {
		t.Fatalf("close file: %v", err)
	}

	server, err := newZipImageServer(archivePath)
	if err != nil {
		t.Fatalf("newZipImageServer: %v", err)
	}
	defer server.Close()

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/images/pic/640/test.jpg", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := server.Handler("/images/pic")
	if err := handler(c); err != nil {
		t.Fatalf("handler returned error: %v", err)
	}

	resp := rec.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("unexpected status: got %d", resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("read body: %v", err)
	}
	if !bytes.Equal(body, content) {
		t.Fatalf("unexpected body: %q", body)
	}

	lastModified := resp.Header.Get(echo.HeaderLastModified)
	if lastModified == "" {
		t.Fatalf("expected Last-Modified header to be set")
	}
	if _, err := time.Parse(http.TimeFormat, lastModified); err != nil {
		t.Fatalf("invalid Last-Modified header: %v", err)
	}

	cacheControl := resp.Header.Get("Cache-Control")
	if cacheControl == "" {
		t.Fatalf("expected Cache-Control header")
	}
}

func TestZipImageServerNotFound(t *testing.T) {
	tmpDir := t.TempDir()
	archivePath := filepath.Join(tmpDir, "images.zip")

	zipFile, err := os.Create(archivePath)
	if err != nil {
		t.Fatalf("create zip: %v", err)
	}
	writer := zip.NewWriter(zipFile)
	if err := writer.Close(); err != nil {
		t.Fatalf("close writer: %v", err)
	}
	if err := zipFile.Close(); err != nil {
		t.Fatalf("close file: %v", err)
	}

	server, err := newZipImageServer(archivePath)
	if err != nil {
		t.Fatalf("newZipImageServer: %v", err)
	}
	defer server.Close()

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/images/pic/640/missing.jpg", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := server.Handler("/images/pic")
	err = handler(c)
	if he, ok := err.(*echo.HTTPError); !ok || he.Code != http.StatusNotFound {
		t.Fatalf("expected not found, got %v", err)
	}
}

func TestCleanZipPath(t *testing.T) {
	tests := map[string]struct {
		input string
		want  string
		ok    bool
	}{
		"simple":         {"640/test.jpg", "640/test.jpg", true},
		"leading_slash":  {"/640/test.jpg", "640/test.jpg", true},
		"dot_prefix":     {"./640/./test.jpg", "640/test.jpg", true},
		"windows_slash":  {"640\\test.jpg", "640/test.jpg", true},
		"parent_discard": {"../secret.jpg", "", false},
		"root":           {".", "", false},
	}

	for name, tc := range tests {
		tc := tc
		t.Run(name, func(t *testing.T) {
			got, ok := cleanZipPath(tc.input)
			if ok != tc.ok || got != tc.want {
				t.Fatalf("cleanZipPath(%q) = (%q,%v), want (%q,%v)", tc.input, got, ok, tc.want, tc.ok)
			}
		})
	}
}
