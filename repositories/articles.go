package repositories

import (
	"database/sql"

	"github.com/nissi1278/go-api-practice/models"
)

// 新規投稿をデータベースにinsertする処理
func InsertArticle(db *sql.DB, article models.Article) (models.Article, error) {
	return article, nil
}

// pageで指定されたページに表示する投稿一覧をデータベースから取得する処理
func SelectArticleList(db *sql.DB, page int) ([]models.Article, error) {
	var articleList []models.Article
	return articleList, nil
}

// articleIDで指定された記事をデータベースから取得する処理
func SelectArticleDetail(db *sql.DB, articleID int) (models.Article, error) {
	var getArticle models.Article
	return getArticle, nil
}

// articleIDで指定された記事のいいね数を+1する処理
func UpdateNiceNum(db *sql.DB, articleID int) error {
	return nil
}

// articleIDで指定された記事についたコメントをデータベースから取得する処理
func SelectCommentList(db *sql.DB, articleID int) ([]models.Comment, error) {
	var commentArray []models.Comment
	return commentArray, nil
}
