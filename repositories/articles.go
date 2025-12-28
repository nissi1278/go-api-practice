package repositories

import (
	"database/sql"

	"github.com/nissi1278/go-api-practice/models"
)

// 新規投稿をデータベースにinsertする処理
func InsertArticle(db *sql.DB, article models.Article) (models.Article, error) {
	const sqlStr = `
		INSERT INTO articles (title, contents, username, nice, created_at)
		VALUES (?, ?, ?, 0, now());
	`

	var newArticle models.Article

	res, err := db.Exec(sqlStr, article.Title, article.Contents, article.UserName)
	if err != nil {
		return models.Article{}, err
	}
	newID, err := res.LastInsertId()
	newArticle.Title, newArticle.Contents, newArticle.UserName = article.Title, article.Contents, article.UserName

	newArticle.ID = int(newID)
	return newArticle, nil
}

// pageで指定されたページに表示する投稿一覧をデータベースから取得する処理
func SelectArticleList(db *sql.DB, page int) ([]models.Article, error) {
	var articleList []models.Article

	pageLimit := 3
	pageOffset := (page - 1) * pageLimit

	const sqlSTr = `
		SELECT article_id, title, contents, username, nice
		FROM articles
		LIMIT ? OFFSET ?
	`

	rows, err := db.Query(sqlSTr, pageLimit, pageOffset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var article models.Article
		err = rows.Scan(
			&article.ID,
			&article.Title,
			&article.Contents,
			&article.UserName,
			&article.NiceNum,
		)
		if err != nil {
			return nil, err
		}
		articleList = append(articleList, article)
	}
	return articleList, nil
}

// articleIDで指定された記事をデータベースから取得する処理
func SelectArticleDetail(db *sql.DB, articleID int) (models.Article, error) {

	const sqlStr = `
		SELECT *
		FROM articles
		WHERE article_id = ?
	`

	var article models.Article
	var createdTime sql.NullTime

	err := db.QueryRow(sqlStr, articleID).Scan(
		&article.ID,
		&article.Title,
		&article.Contents,
		&article.UserName,
		&article.NiceNum,
		&createdTime,
	)

	if err != nil {
		return models.Article{}, err
	}

	if createdTime.Valid {
		article.CreatedAt = createdTime.Time
	}

	return article, nil
}

// articleIDで指定された記事のいいね数を+1する処理
func UpdateNiceNum(db *sql.DB, articleID int) error {
	const sqlStr = `
		UPDATE articles
		SET nice = nice + 1
		WHERE article_id = ?
	`

	_, err := db.Exec(sqlStr, articleID)
	if err != nil {
		return err
	}
	return nil
}
