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

	createZipArchive(t, archivePath, map[string][]byte{
		"640/test.jpg": content,
	}, modTime)

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

	createZipArchive(t, archivePath, map[string][]byte{}, time.Now())

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

func TestZipMountHandler(t *testing.T) {
	tmpDir := t.TempDir()
	modTime := time.Date(2023, 5, 1, 9, 0, 0, 0, time.UTC)

	zip640 := filepath.Join(tmpDir, "640.zip")
	createZipArchive(t, zip640, map[string][]byte{
		"640/a.jpg": []byte("thumb"),
	}, modTime)

	zipFull := filepath.Join(tmpDir, "full.zip")
	createZipArchive(t, zipFull, map[string][]byte{
		"full/a.jpg": []byte("fullsize"),
	}, modTime)

	server640, err := newZipImageServer(zip640)
	if err != nil {
		t.Fatalf("newZipImageServer 640: %v", err)
	}
	defer server640.Close()

	serverFull, err := newZipImageServer(zipFull)
	if err != nil {
		t.Fatalf("newZipImageServer full: %v", err)
	}
	defer serverFull.Close()

	handler := newZipMountHandler("/images/pic", map[string]*zipImageServer{
		"640":  server640,
		"full": serverFull,
	})

	e := echo.New()

	// 640 request
	req := httptest.NewRequest(http.MethodGet, "/images/pic/640/a.jpg", nil)
	rec := httptest.NewRecorder()
	if err := handler(e.NewContext(req, rec)); err != nil {
		t.Fatalf("640 handler error: %v", err)
	}
	resp := rec.Result()
	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		t.Fatalf("read 640 body: %v", err)
	}
	if string(body) != "thumb" {
		t.Fatalf("unexpected 640 body: %q", body)
	}

	// full request
	req = httptest.NewRequest(http.MethodGet, "/images/pic/full/a.jpg", nil)
	rec = httptest.NewRecorder()
	if err := handler(e.NewContext(req, rec)); err != nil {
		t.Fatalf("full handler error: %v", err)
	}
	resp = rec.Result()
	body, err = io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		t.Fatalf("read full body: %v", err)
	}
	if string(body) != "fullsize" {
		t.Fatalf("unexpected full body: %q", body)
	}

	// unknown segment
	req = httptest.NewRequest(http.MethodGet, "/images/pic/other/a.jpg", nil)
	rec = httptest.NewRecorder()
	err = handler(e.NewContext(req, rec))
	if he, ok := err.(*echo.HTTPError); !ok || he.Code != http.StatusNotFound {
		t.Fatalf("expected not found for other segment, got %v", err)
	}
}

func createZipArchive(t *testing.T, path string, files map[string][]byte, modTime time.Time) {
	t.Helper()

	zipFile, err := os.Create(path)
	if err != nil {
		t.Fatalf("create zip: %v", err)
	}
	defer zipFile.Close()

	writer := zip.NewWriter(zipFile)

	for name, content := range files {
		header := &zip.FileHeader{
			Name:   name,
			Method: zip.Deflate,
		}
		header.SetModTime(modTime)
		entryWriter, err := writer.CreateHeader(header)
		if err != nil {
			t.Fatalf("create header %s: %v", name, err)
		}
		if _, err := entryWriter.Write(content); err != nil {
			t.Fatalf("write content %s: %v", name, err)
		}
	}

	if err := writer.Close(); err != nil {
		t.Fatalf("close writer: %v", err)
	}
}
