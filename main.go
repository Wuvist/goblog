package main

import (
	"database/sql"
	"net/http"
	"os"
	"strings"

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
//go:generate gorazor skins/skin5_UserCate.gohtml tpl/skins/skin5_UserCate.go
//go:generate gorazor skins/skin5_default.gohtml tpl/skins/skin5_default.go

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
	e.GET("/blogger.go", blogger)
	e.GET("/cate.go", cate)
	e.GET("/blog.go", blog)

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
	blogerUsername := c.QueryParam("blogger")

	bloggerData, err := models.BloggersG(qm.Where("id = ?", blogerUsername)).One()
	if err != nil {
		println(err.Error())
	}
	if bloggerData == nil {
		return c.String(http.StatusNotFound, "找不到博客")
	}

	blogger := tpl.NewBloggerFromDb(bloggerData)

	blogs := tpl.GetBlogSummariesFromBlogger(bloggerData.Index)

	page := skins.Skin5_default(blogger, blogs)

	return c.HTML(http.StatusOK, page)
}

func cate(c echo.Context) error {
	blogerUsername := strings.ToLower(c.QueryParam("blogger"))
	cateID := c.QueryParam("cate_id")
	cateData, _ := models.UserdefinecategoriesG(qm.Where("`index` = ?", cateID)).One()
	if cateData == nil {
		return c.String(http.StatusNotFound, "找不到网志分类")
	}

	bloggerData, _ := models.BloggersG(qm.Where("`index` = ?", cateData.Blogger)).One()
	if bloggerData == nil {
		return c.String(http.StatusNotFound, "找不到博客")
	}

	blogger := tpl.NewBloggerFromDb(bloggerData)
	if strings.ToLower(blogger.Username) != blogerUsername {
		return c.String(http.StatusNotFound, "找不到分类")
	}

	cate := &tpl.Cate{}
	cate.CateID = cateData.Index
	cate.CateName = cateData.Cate

	blogs := tpl.GetBlogSummariesFromCate(cate.CateID)

	page := skins.Skin5_UserCate(blogger, cate, blogs)

	return c.HTML(http.StatusOK, page)
}

func blog(c echo.Context) error {
	blogerUsername := strings.ToLower(c.QueryParam("blogger"))
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
	if strings.ToLower(blogger.Username) != blogerUsername {
		return c.String(http.StatusNotFound, "找不到网志")
	}

	blog := tpl.NewBlogFromDb(blogData)
	comments := tpl.GetBlogComments(blogData.Index)

	page := skins.Skin5_comment(blogger, blog, comments)

	return c.HTML(http.StatusOK, page)
}
