package tpl

import (
	"github.com/Wuvist/goblog/models"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

var dateFormat = "2006-01-02 15:04:05"

// Blog is the struture for displaying a blog
type Blog struct {
	Title    string
	CateName string
	CateID   int
	Content  string
	Tags     string
	Adddate  string
}

// NewBlogFromDb create blog struct from db model
func NewBlogFromDb(data *models.Article) *Blog {
	b := &Blog{}
	b.Title = data.Title.String
	if b.Title == "" {
		b.Title = "无题"
	}
	b.Content = data.Content.String
	b.Adddate = data.AddDate.Time.Format(dateFormat)

	cateData, _ := models.UserdefinecategoriesG(qm.Where("`index` = ?", data.Cate.Int)).One()
	if cateData != nil {
		b.CateID = cateData.Index
		b.CateName = cateData.Cate
	}
	return b
}

// Cate is blogger's own blog categories
type Cate struct {
	CateName  string
	CateID    int
	BlogCount int
}

// Blogger is the struture for blogger info
type Blogger struct {
	Username     string
	BlogName     string
	Info         string
	Nick         string
	VisitorCount int
	Cates        []*Cate
}

// NewBloggerFromDb create blogger struct from db model
func NewBloggerFromDb(data *models.Blogger) *Blogger {
	b := &Blogger{}
	b.Username = data.ID
	b.Nick = data.Nick.String
	b.Info = data.Intro.String
	b.BlogName = data.Blogname
	b.VisitorCount = data.Visitor

	objs, _ := models.UserdefinecategoriesG(qm.Where("`blogger` = ?", data.Index)).All()
	b.Cates = make([]*Cate, len(objs))
	for i := 0; i < len(objs); i++ {
		cate := &Cate{}
		obj := objs[i]
		cate.CateName = obj.Cate
		cate.CateID = obj.Index
		cate.BlogCount = obj.Files
		b.Cates[i] = cate
	}

	return b
}

// Comment is the struture for displaying a comment
type Comment struct {
	Author  string
	Adddate string
	Content string
}
