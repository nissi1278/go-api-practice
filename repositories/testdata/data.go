package testdata

import "github.com/nissi1278/go-api-practice/models"

var ArticleTestData = []models.Article{
	models.Article{
		ID:       1,
		Title:    "firstPost",
		Contents: "This is my first blog",
		UserName: "nissi",
		NiceNum:  10,
	},
	models.Article{
		ID:       2,
		Title:    "second Post",
		Contents: "This is my second blog",
		UserName: "nissi",
		NiceNum:  2,
	},
}

var InsertArticleTestData = models.Article{
	Title:    "Test title",
	Contents: "Test contents",
	UserName: "Test username",
	NiceNum:  3,
}

var CommentTestData = []models.Comment{}
