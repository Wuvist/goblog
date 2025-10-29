package tpl

import (
	"context"

	"github.com/Wuvist/goblog/models"
	"github.com/aarondl/sqlboiler/v4/boil"
	"github.com/aarondl/sqlboiler/v4/queries/qm"
)

var dateFormat = "2006-01-02 15:04:05"

// BlogSummary represents minimal blog metadata for list views.
type BlogSummary struct {
	ID           int
	Title        string
	Adddate      string
	CommentCount int
}

// Blog represents a full blog post for rendering views.
type Blog struct {
	Title    string
	CateName string
	CateID   int
	Content  string
	Tags     string
	Adddate  string
}

// GetBlogSummariesFromCate returns all blogs in a category ordered by latest first.
func GetBlogSummariesFromCate(ctx context.Context, cateID int64) []*BlogSummary {
	exec := boil.GetContextDB()
	if exec == nil {
		return nil
	}

	records, err := models.Articles(
		qm.Where("cate = ?", cateID),
		qm.OrderBy("\"articles\".\"index\" DESC"),
	).All(ctx, exec)
	if err != nil {
		return nil
	}

	blogs := make([]*BlogSummary, 0, len(records))
	for _, obj := range records {
		summary := &BlogSummary{
			Title:        obj.Title.String,
			CommentCount: int(obj.Comment),
		}
		if obj.Index.Valid {
			summary.ID = int(obj.Index.Int64)
		}
		if obj.AddDate.Valid {
			summary.Adddate = obj.AddDate.Time.Format(dateFormat)
		}
		blogs = append(blogs, summary)
	}

	return blogs
}

// GetBlogSummariesFromBlogger returns the latest 20 blogs for a blogger.
func GetBlogSummariesFromBlogger(ctx context.Context, bloggerID int64) []*BlogSummary {
	exec := boil.GetContextDB()
	if exec == nil {
		return nil
	}

	records, err := models.Articles(
		qm.Where("blogger = ?", bloggerID),
		qm.OrderBy("\"articles\".\"index\" DESC"),
		qm.Limit(20),
	).All(ctx, exec)
	if err != nil {
		return nil
	}

	blogs := make([]*BlogSummary, 0, len(records))
	for _, obj := range records {
		summary := &BlogSummary{
			Title:        obj.Title.String,
			CommentCount: int(obj.Comment),
		}
		if obj.Index.Valid {
			summary.ID = int(obj.Index.Int64)
		}
		if obj.AddDate.Valid {
			summary.Adddate = obj.AddDate.Time.Format(dateFormat)
		}
		blogs = append(blogs, summary)
	}

	return blogs
}

// GetBlogComments fetches all comments for a blog sorted by newest first.
func GetBlogComments(ctx context.Context, blogID int64) []*Comment {
	exec := boil.GetContextDB()
	if exec == nil {
		return nil
	}

	records, err := models.Comments(
		qm.Where("article = ?", blogID),
		qm.OrderBy("\"comment\".\"index\" DESC"),
	).All(ctx, exec)
	if err != nil {
		return nil
	}

	comments := make([]*Comment, 0, len(records))
	for _, obj := range records {
		comment := &Comment{
			Content: obj.Content,
			Adddate: obj.AddDate.Format(dateFormat),
		}
		if obj.Author.Valid {
			comment.Author = obj.Author.String
		}
		comments = append(comments, comment)
	}

	return comments
}

// NewBlogFromDb materialises a Blog from the generated Article model.
func NewBlogFromDb(ctx context.Context, data *models.Article) *Blog {
	if data == nil {
		return nil
	}

	blog := &Blog{
		Title:   data.Title.String,
		Content: data.Content.String,
	}
	if blog.Title == "" {
		blog.Title = "无题"
	}
	if data.AddDate.Valid {
		blog.Adddate = data.AddDate.Time.Format(dateFormat)
	}

	if data.Cate.Valid {
		exec := boil.GetContextDB()
		if exec != nil {
			cate, err := models.Userdefinecategories(
				qm.Where("\"index\" = ?", data.Cate.Int64),
			).One(ctx, exec)
			if err == nil && cate != nil {
				if cate.Index.Valid {
					blog.CateID = int(cate.Index.Int64)
				}
				blog.CateName = cate.Cate
			}
		}
	}

	return blog
}

// Cate represents a blogger's custom category.
type Cate struct {
	CateName  string
	CateID    int
	BlogCount int
}

// Link represents a blogger's custom link.
type Link struct {
	URL  string
	Link string
}

// Blogger aggregates blogger metadata and their related categories/links.
type Blogger struct {
	Username     string
	BlogName     string
	Info         string
	Nick         string
	VisitorCount int
	Cates        []*Cate
	Links        []*Link
}

func getBloggerCates(ctx context.Context, bloggerID int64) []*Cate {
	exec := boil.GetContextDB()
	if exec == nil {
		return nil
	}

	records, err := models.Userdefinecategories(
		qm.Where("blogger = ?", bloggerID),
	).All(ctx, exec)
	if err != nil {
		return nil
	}

	cates := make([]*Cate, 0, len(records))
	for _, obj := range records {
		cate := &Cate{
			CateName:  obj.Cate,
			BlogCount: int(obj.Files),
		}
		if obj.Index.Valid {
			cate.CateID = int(obj.Index.Int64)
		}
		cates = append(cates, cate)
	}

	return cates
}

func getBloggerLinks(ctx context.Context, bloggerID int64) []*Link {
	exec := boil.GetContextDB()
	if exec == nil {
		return nil
	}

	records, err := models.Links(
		qm.Where("blogger = ? AND reveal = 1", bloggerID),
	).All(ctx, exec)
	if err != nil {
		return nil
	}

	links := make([]*Link, 0, len(records))
	for _, obj := range records {
		link := &Link{
			Link: obj.Link,
			URL:  obj.URL,
		}
		links = append(links, link)
	}

	return links
}

// NewBloggerFromDb hydrates a Blogger view model from the generated Blogger model.
func NewBloggerFromDb(ctx context.Context, data *models.Blogger) *Blogger {
	if data == nil {
		return nil
	}

	blogger := &Blogger{
		Username:     data.ID,
		Nick:         data.Nick.String,
		Info:         data.Intro.String,
		BlogName:     data.Blogname,
		VisitorCount: int(data.Visitor),
	}

	bloggerID := int64(0)
	if data.Index.Valid {
		bloggerID = data.Index.Int64
	}

	blogger.Cates = getBloggerCates(ctx, bloggerID)
	blogger.Links = getBloggerLinks(ctx, bloggerID)

	return blogger
}

// Comment is the struture for displaying a comment
type Comment struct {
	Author  string
	Adddate string
	Content string
}
