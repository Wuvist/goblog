package tpl

import (
	"github.com/Wuvist/goblog/models"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

var dateFormat string = "2006-01-02 15:04:05"

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
	b.Content = data.Content.String
	b.Adddate = data.AddDate.Time.Format(dateFormat)

	return b
}

// Blogger is the struture for blogger info
type Blogger struct {
	Username     string
	BlogName     string
	Info         string
	Nick         string
	VisitorCount int
}

// NewBloggerFromDb create blogger struct from db model
func NewBloggerFromDb(data *models.Blogger) *Blogger {
	b := &Blogger{}
	b.Username = data.ID
	b.Nick = data.Nick.String
	b.Info = data.Intro.String
	b.BlogName = data.Blogname
	b.VisitorCount = data.Visitor
	return b
}

// Comment is the struture for displaying a comment
type Comment struct {
	Author  string
	Adddate string
	Content string
}
