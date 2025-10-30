package main

import (
	"archive/zip"
	"fmt"
	"io"
	"mime"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	"github.com/labstack/echo"
)

// zipImageServer serves files stored inside a zip archive without extracting
// them to disk.
type zipImageServer struct {
	file     *os.File
	reader   *zip.Reader
	fileByID map[string]*zip.File
}

// newZipImageServer opens the provided zip archive and indexes its entries.
func newZipImageServer(zipPath string) (*zipImageServer, error) {
	file, err := os.Open(zipPath)
	if err != nil {
		return nil, fmt.Errorf("open zip: %w", err)
	}

	info, err := file.Stat()
	if err != nil {
		file.Close()
		return nil, fmt.Errorf("zip stat: %w", err)
	}

	reader, err := zip.NewReader(file, info.Size())
	if err != nil {
		file.Close()
		return nil, fmt.Errorf("zip read: %w", err)
	}

	index := make(map[string]*zip.File, len(reader.File))
	for _, f := range reader.File {
		name, ok := cleanZipPath(f.Name)
		if !ok {
			continue
		}
		index[name] = f
	}

	return &zipImageServer{
		file:     file,
		reader:   reader,
		fileByID: index,
	}, nil
}

func (s *zipImageServer) Close() error {
	if s == nil || s.file == nil {
		return nil
	}
	return s.file.Close()
}

func (s *zipImageServer) Handler(prefix string) echo.HandlerFunc {
	trimmedPrefix := strings.TrimSuffix(prefix, "/")
	return func(c echo.Context) error {
		requestPath := c.Request().URL.Path
		if !strings.HasPrefix(requestPath, trimmedPrefix) {
			return echo.NewHTTPError(http.StatusNotFound)
		}

		rel := strings.TrimPrefix(requestPath, trimmedPrefix)
		rel = strings.TrimPrefix(rel, "/")
		if rel == "" {
			return echo.NewHTTPError(http.StatusNotFound)
		}

		key, ok := cleanZipPath(rel)
		if !ok {
			return echo.NewHTTPError(http.StatusNotFound)
		}

		entry, ok := s.fileByID[key]
		if !ok {
			return echo.NewHTTPError(http.StatusNotFound)
		}

		modTime := entry.Modified.UTC()
		if !modTime.IsZero() {
			ifModifiedSince := c.Request().Header.Get(echo.HeaderIfModifiedSince)
			if ifModifiedSince != "" {
				if t, err := time.Parse(http.TimeFormat, ifModifiedSince); err == nil && modTime.Before(t.Add(1*time.Second)) {
					return c.NoContent(http.StatusNotModified)
				}
			}
			c.Response().Header().Set(echo.HeaderLastModified, modTime.Format(http.TimeFormat))
		}

		contentType := mime.TypeByExtension(strings.ToLower(path.Ext(entry.Name)))
		if contentType == "" {
			contentType = "application/octet-stream"
		}
		c.Response().Header().Set(echo.HeaderContentType, contentType)
		c.Response().Header().Set("Cache-Control", "public, max-age=604800, immutable")

		if entry.UncompressedSize64 > 0 {
			c.Response().Header().Set(echo.HeaderContentLength, fmt.Sprintf("%d", entry.UncompressedSize64))
		}

		reader, err := entry.Open()
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "failed to open image")
		}
		defer reader.Close()

		c.Response().WriteHeader(http.StatusOK)
		if _, err := io.Copy(c.Response().Writer, reader); err != nil {
			return err
		}
		return nil
	}
}

func cleanZipPath(name string) (string, bool) {
	name = strings.TrimPrefix(name, "Users/wuvist/temp/")
	return name, true
}
