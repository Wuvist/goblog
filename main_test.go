package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
)

func TestFriendlyRewriteBlogger(t *testing.T) {
	e := echo.New()
	e.Pre(friendlyURLRewrite())

	var captured string
	e.GET("/blogger.go", func(c echo.Context) error {
		captured = c.QueryParam("blogger")
		if got := c.Request().URL.Path; got != "/blogger.go" {
			return c.String(http.StatusInternalServerError, got)
		}
		return c.String(http.StatusOK, "ok")
	})

	req := httptest.NewRequest(http.MethodGet, "/foo/", nil)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status OK, got %d", rec.Code)
	}
	if captured != "foo" {
		t.Fatalf("expected blogger=foo, got %q", captured)
	}
}

func TestFriendlyRewriteCate(t *testing.T) {
	e := echo.New()
	e.Pre(friendlyURLRewrite())

	var blogger, cateID string
	e.GET("/cate.go", func(c echo.Context) error {
		blogger = c.QueryParam("blogger")
		cateID = c.QueryParam("cate_id")
		if got := c.Request().URL.Path; got != "/cate.go" {
			return c.String(http.StatusInternalServerError, got)
		}
		return c.String(http.StatusOK, "ok")
	})

	req := httptest.NewRequest(http.MethodGet, "/foo/cate42.shtml", nil)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status OK, got %d", rec.Code)
	}
	if blogger != "foo" || cateID != "42" {
		t.Fatalf("unexpected query params blogger=%q cate_id=%q", blogger, cateID)
	}
}

func TestFriendlyRewriteBlog(t *testing.T) {
	e := echo.New()
	e.Pre(friendlyURLRewrite())

	var blogger, articleID string
	e.GET("/blog.go", func(c echo.Context) error {
		blogger = c.QueryParam("blogger")
		articleID = c.QueryParam("article_id")
		if got := c.Request().URL.Path; got != "/blog.go" {
			return c.String(http.StatusInternalServerError, got)
		}
		return c.String(http.StatusOK, "ok")
	})

	req := httptest.NewRequest(http.MethodGet, "/foo/123.shtml", nil)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status OK, got %d", rec.Code)
	}
	if blogger != "foo" || articleID != "123" {
		t.Fatalf("unexpected query params blogger=%q article_id=%q", blogger, articleID)
	}
}

func TestFriendlyRewriteOldImage(t *testing.T) {
	e := echo.New()
	e.Pre(friendlyURLRewrite())

	req := httptest.NewRequest(http.MethodGet, "/foo/bar.jpg", nil)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	if rec.Code != http.StatusMovedPermanently {
		t.Fatalf("expected 301 redirect, got %d", rec.Code)
	}
	location := rec.Header().Get("Location")
	expected := "https://storage.googleapis.com/blogwind/images/old/foo/bar.jpg"
	if location != expected {
		t.Fatalf("unexpected redirect location %q", location)
	}
}
