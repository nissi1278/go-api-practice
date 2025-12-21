package models

import "time"

var (
	Comment1 = Comment{
		CommentID: 1,
		ArticleID: 1,
		Message:   "first test comment",
		CreatedAt: time.Now(),
	}
	Comment2 = Comment{
		CommentID: 2,
		ArticleID: 1,
		Message:   "second test comment",
		CreatedAt: time.Now(),
	}
)

var (
	Article1 = Article{
		ID:          1,
		Title:       "first test title",
		Contents:    "this is the test article1",
		UserName:    "nissi",
		NiceNum:     1,
		CommentList: []Comment{Comment1, Comment2},
		CreatedAt:   time.Now(),
	}

	Article2 = Article{
		ID:          2,
		Title:       "second test title",
		Contents:    "this is the test article2",
		UserName:    "nissi",
		NiceNum:     5,
		CommentList: []Comment{},
		CreatedAt:   time.Now(),
	}
)
