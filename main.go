package main

import (
	"database/sql"
	"net/http"
	"os"

	"github.com/Wuvist/goblog/models"
	"github.com/Wuvist/goblog/static"
	"github.com/volatiletech/sqlboiler/queries/qm"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/volatiletech/sqlboiler/boil"
)

//go:generate esc -o static/css.go -pkg static -include="(.+).css" -prefix="oldweb/" oldweb/Template

func main() {
	// db
	connstr := os.Getenv("connstr")
	db, err := sql.Open("mysql", connstr+"?parseTime=true")
	if err != nil {
		panic(err)
	}
	boil.SetDB(db)

	// web
	e := echo.New()
	e.GET("/", home)
	e.GET("/blogger", blogger)
	e.GET("/cate", cate)
	e.GET("/blog", blog)

	// static
	e.GET("/Template/*", echo.WrapHandler(http.FileServer(static.FS(false))))

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

// Handler
func home(c echo.Context) error {
	return c.String(http.StatusOK, "homepage")
}

func blogger(c echo.Context) error {
	bloggerID := c.QueryParam("id")
	blogger, _ := models.BloggersG(qm.Where("id = ?", bloggerID)).One()
	if blogger == nil {
		return c.String(http.StatusNotFound, "找不到博客")
	}

	return c.String(http.StatusOK, blogger.Nick.String)
}

func cate(c echo.Context) error {
	return c.String(http.StatusOK, "category")
}

func blog(c echo.Context) error {
	return c.String(http.StatusOK, "blog page")
}
