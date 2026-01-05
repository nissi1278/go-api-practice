package repositories_test

import (
	"testing"

	"github.com/nissi1278/go-api-practice/repositories"
)

func TestSelectCommentList(t *testing.T) {
	expectedID := 1

	got, err := repositories.SelectCommentList(testDB, expectedID)
	if err != nil {
		t.Fatal(err)
	}

	for _, comment := range got {
		if comment.ArticleID != expectedID {
			t.Errorf("get comment articleID expected %d but got %d\n", expectedID, comment.ArticleID)
		}
	}

}
