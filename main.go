package main

import (
	"database/sql"
	"net/http"
	"os"

	"github.com/Wuvist/goblog/static"
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
	return c.String(http.StatusOK, "bloger page")
}

func cate(c echo.Context) error {
	return c.String(http.StatusOK, "category")
}

func blog(c echo.Context) error {
	return c.String(http.StatusOK, "blog page")
}
