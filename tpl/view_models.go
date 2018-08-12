package tpl

// Blog is the struture for displaying a blog
type Blog struct {
	Title    string
	CateName string
	CateID   int
	Content  string
	Tags     string
	Adddate  string
}

// Blogger is the struture for blogger info
type Blogger struct {
	Username string
	BlogName string
	Info     string
	Nick     string
}

// Comment is the struture for displaying a comment
type Comment struct {
	Author  string
	Adddate string
	Content string
}
