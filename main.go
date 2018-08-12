package main

import (
	"database/sql"
	"net/http"
	"os"

	"github.com/Wuvist/goblog/models"
	"github.com/Wuvist/goblog/static"
	"github.com/Wuvist/goblog/tpl"
	"github.com/Wuvist/goblog/tpl/skins"
	"github.com/volatiletech/sqlboiler/queries/qm"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/volatiletech/sqlboiler/boil"
)

// run manually, not sure why go generate failed
// esc -o static/css.go -pkg static -include="(.+).css" -prefix="oldweb/" oldweb/Template

//go:generate gorazor skins/skin5_comment.gohtml tpl/skins/skin5_comment.go

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
	blogID := c.QueryParam("article_id")
	blogData, err := models.ArticlesG(qm.Where("`index` = ?", blogID)).One()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	bloggerData, _ := models.BloggersG(qm.Where("`index` = ?", blogData.Blogger)).One()
	if bloggerData == nil {
		return c.String(http.StatusNotFound, "找不到博客")
	}

	blogger := tpl.NewBloggerFromDb(bloggerData)
	blog := tpl.NewBlogFromDb(blogData)
	comments := []*tpl.Comment{}

	page := skins.Skin5_comment(blogger, blog, comments)

	return c.HTML(http.StatusOK, page)
}
