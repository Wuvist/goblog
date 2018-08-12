package main

import (
	"database/sql"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/volatiletech/sqlboiler/boil"
)

func main() {
	// db
	connstr := os.Getenv("connstr")
	db, err := sql.Open("mysql", connstr)
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
