package main

import (
	"database/sql"
	"errors"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/Wuvist/goblog/models"
	"github.com/Wuvist/goblog/static"
	"github.com/Wuvist/goblog/tpl"
	"github.com/Wuvist/goblog/tpl/skins"
	"github.com/aarondl/sqlboiler/v4/boil"
	"github.com/aarondl/sqlboiler/v4/queries/qm"
	"github.com/labstack/echo"
	_ "modernc.org/sqlite"
)

// run manually, not sure why go generate failed
// esc -o static/css.go -pkg static -include="(.+).css" -prefix="oldweb/" oldweb/Template

//go:generate gorazor skins tpl/skins

func main() {
	// db
	connstr := os.Getenv("connstr")
	if connstr == "" {
		connstr = "file:blogwind.db?_pragma=foreign_keys(ON)&_pragma=busy_timeout=5000"
	}

	db, err := sql.Open("sqlite", connstr)
	if err != nil {
		panic(err)
	}
	boil.SetDB(db)

	// web
	e := echo.New()
	e.Pre(friendlyURLRewrite())
	e.GET("/", home)
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
	ctx := c.Request().Context()

	objs, err := models.Bloggers(
		qm.Where("Reveal = 1 AND blogs > ?", 0),
		qm.OrderBy("Last_post DESC"),
	).AllG(ctx)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	bloggers := make([]*tpl.Blogger, len(objs))
	for i := 0; i < len(objs); i++ {
		b := &tpl.Blogger{}
		obj := objs[i]
		b.Username = obj.ID
		b.BlogName = obj.Blogname
		b.Nick = obj.Nick.String
		bloggers[i] = b
	}

	page := skins.Home(bloggers)
	return c.HTML(http.StatusOK, page)
}

func blogger(c echo.Context) error {
	blogerUsername := c.QueryParam("blogger")
	ctx := c.Request().Context()

	bloggerData, err := models.Bloggers(qm.Where("id = ?", blogerUsername)).OneG(ctx)
	if errors.Is(err, sql.ErrNoRows) || bloggerData == nil {
		return c.String(http.StatusNotFound, "找不到博客")
	}
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	blogger := tpl.NewBloggerFromDb(ctx, bloggerData)

	bloggerID := int64(0)
	if bloggerData.Index.Valid {
		bloggerID = bloggerData.Index.Int64
	}
	blogs := tpl.GetBlogSummariesFromBlogger(ctx, bloggerID)

	page := skins.Skin5_default(blogger, blogs)

	return c.HTML(http.StatusOK, page)
}

func cate(c echo.Context) error {
	blogerUsername := strings.ToLower(c.QueryParam("blogger"))
	rawCateID := c.QueryParam("cate_id")
	cateID, err := strconv.ParseInt(rawCateID, 10, 64)
	if err != nil {
		return c.String(http.StatusBadRequest, "网志分类参数错误")
	}

	ctx := c.Request().Context()

	cateData, err := models.Userdefinecategories(qm.Where("\"index\" = ?", cateID)).OneG(ctx)
	if errors.Is(err, sql.ErrNoRows) || cateData == nil {
		return c.String(http.StatusNotFound, "找不到网志分类")
	}
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	bloggerData, err := models.Bloggers(qm.Where("\"index\" = ?", cateData.Blogger)).OneG(ctx)
	if errors.Is(err, sql.ErrNoRows) || bloggerData == nil {
		return c.String(http.StatusNotFound, "找不到博客")
	}
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	blogger := tpl.NewBloggerFromDb(ctx, bloggerData)
	if strings.ToLower(blogger.Username) != blogerUsername {
		return c.String(http.StatusNotFound, "找不到分类")
	}

	cate := &tpl.Cate{}
	if cateData.Index.Valid {
		cate.CateID = int(cateData.Index.Int64)
	}
	cate.CateName = cateData.Cate

	blogs := tpl.GetBlogSummariesFromCate(ctx, cateID)

	page := skins.Skin5_UserCate(blogger, cate, blogs)

	return c.HTML(http.StatusOK, page)
}

func blog(c echo.Context) error {
	blogerUsername := strings.ToLower(c.QueryParam("blogger"))
	rawBlogID := c.QueryParam("article_id")
	blogID, err := strconv.ParseInt(rawBlogID, 10, 64)
	if err != nil {
		return c.String(http.StatusBadRequest, "网志参数错误")
	}

	ctx := c.Request().Context()
	blogData, err := models.Articles(qm.Where("\"index\" = ?", blogID)).OneG(ctx)
	if errors.Is(err, sql.ErrNoRows) || blogData == nil {
		return c.String(http.StatusNotFound, "找不到网志")
	}
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	bloggerData, err := models.Bloggers(qm.Where("\"index\" = ?", blogData.Blogger)).OneG(ctx)
	if errors.Is(err, sql.ErrNoRows) || bloggerData == nil {
		return c.String(http.StatusNotFound, "找不到博客")
	}

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	blogger := tpl.NewBloggerFromDb(ctx, bloggerData)
	if strings.ToLower(blogger.Username) != blogerUsername {
		return c.String(http.StatusNotFound, "找不到网志")
	}

	blog := tpl.NewBlogFromDb(ctx, blogData)

	articleID := int64(0)
	if blogData.Index.Valid {
		articleID = blogData.Index.Int64
	}
	comments := tpl.GetBlogComments(ctx, articleID)

	page := skins.Skin5_comment(blogger, blog, comments)

	return c.HTML(http.StatusOK, page)
}

var (
	blogURLPattern    = regexp.MustCompile(`^/([A-Za-z0-9_]+)/([0-9]+)\.shtml$`)
	cateURLPattern    = regexp.MustCompile(`^/([A-Za-z0-9_]+)/cate([0-9]+)\.shtml$`)
	bloggerURLPattern = regexp.MustCompile(`^/([A-Za-z0-9_]+)/?$`)

	reservedFriendlyRoots = map[string]struct{}{
		"template": {},
	}
)

func friendlyURLRewrite() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			path := req.URL.Path

			if path == "/" {
				return next(c)
			}

			if matches := blogURLPattern.FindStringSubmatch(path); matches != nil {
				username := matches[1]
				articleID := matches[2]
				rewriteRequest(req, "/blog.go", map[string]string{
					"blogger":    username,
					"article_id": articleID,
				})
				c.SetPath("/blog.go")
			} else if matches := cateURLPattern.FindStringSubmatch(path); matches != nil {
				username := matches[1]
				cateID := matches[2]
				rewriteRequest(req, "/cate.go", map[string]string{
					"blogger": username,
					"cate_id": cateID,
				})
				c.SetPath("/cate.go")
			} else if matches := bloggerURLPattern.FindStringSubmatch(path); matches != nil {
				username := matches[1]
				if _, ok := reservedFriendlyRoots[strings.ToLower(username)]; ok {
					return next(c)
				}
				rewriteRequest(req, "/blogger.go", map[string]string{
					"blogger": username,
				})
				c.SetPath("/blogger.go")
			}

			return next(c)
		}
	}
}

func rewriteRequest(req *http.Request, targetPath string, params map[string]string) {
	query := req.URL.Query()
	for key, value := range params {
		query.Set(key, value)
	}
	req.URL.RawQuery = query.Encode()
	req.URL.Path = targetPath
	req.URL.RawPath = targetPath
	if req.URL.RawQuery != "" {
		req.RequestURI = targetPath + "?" + req.URL.RawQuery
	} else {
		req.RequestURI = targetPath
	}
}
