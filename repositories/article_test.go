package repositories_test

import (
	"testing"

	"github.com/nissi1278/go-api-practice/models"
	"github.com/nissi1278/go-api-practice/repositories"
	"github.com/nissi1278/go-api-practice/repositories/testdata"
)

func TestSelectArticleList(t *testing.T) {
	getPageNum := 1
	expectedNum := 3

	got, err := repositories.SelectArticleList(testDB, getPageNum)
	if err != nil {
		t.Fatal(err)
	}

	if num := len(got); num != expectedNum {
		t.Errorf("want %d but got %d articles\n", expectedNum, num)
	}
}

func TestSelectArticleDetail(t *testing.T) {
	tests := []struct {
		testTitle string
		expected  models.Article
	}{
		{
			testTitle: "subTest1",
			expected:  testdata.ArticleTestData[0],
		},
		{
			testTitle: "subTest2",
			expected:  testdata.ArticleTestData[1],
		},
	}

	for _, test := range tests {
		t.Run(test.testTitle, func(t *testing.T) {
			got, err := repositories.SelectArticleDetail(testDB, test.expected.ID)
			if err != nil {
				t.Fatal(err)
			}

			if got.ID != test.expected.ID {
				t.Errorf("ID: get %d but want %d\n", got.ID, test.expected.ID)
			}

			if got.Title != test.expected.Title {
				t.Errorf("Title: get %s but want %s\n", got.Title, test.expected.Title)
			}

			if got.Contents != test.expected.Contents {
				t.Errorf("Contents: get %s but want %s\n", got.Contents, test.expected.Contents)
			}

			if got.UserName != test.expected.UserName {
				t.Errorf("UserName: get %s but want %s\n", got.UserName, test.expected.UserName)
			}

			if got.NiceNum != test.expected.NiceNum {
				t.Errorf("Nice: get %d but want %d\n", got.NiceNum, test.expected.NiceNum)
			}
		})
	}
}

func TestInsertArticle(t *testing.T) {
	article := testdata.InsertArticleTestData

	expectedArticleTitle := "Test title"
	got, err := repositories.InsertArticle(testDB, article)
	if err != nil {
		t.Fatal(err)
	}

	if expectedArticleTitle != got.Title {
		t.Errorf("new article title is expected %s but got %s\n", expectedArticleTitle, got.Title)
	}

	t.Cleanup(func() {
		const sqlStr = `
			DELETE
			FROM articles
			WHERE title = ?
			AND	  contents = ?
			AND   username = ?
		`

		testDB.Exec(sqlStr, article.Title, article.Contents, article.UserName)
	})
}

func TestUpdateNiceNum(t *testing.T) {
	targetArticleID := 1
	expectedArticleNice := 11
	err := repositories.UpdateNiceNum(testDB, targetArticleID)
	if err != nil {
		t.Fatal(err)
	}

	article, err := repositories.SelectArticleDetail(testDB, targetArticleID)

	if expectedArticleNice != article.NiceNum {
		t.Errorf("update Nice is expected %d but got %d\n", expectedArticleNice, article.NiceNum)
	}

	t.Cleanup(func() {
		const sqlStr = `
		UPDATE articles
		SET nice = nice - 1
		WHERE article_id = ?
		`
		testDB.Exec(sqlStr, targetArticleID)
	})

}
