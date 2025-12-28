package repositories

import (
	"database/sql"

	"github.com/nissi1278/go-api-practice/models"
)

// articleIDで指定された記事についたコメントをデータベースから取得する処理
func SelectCommentList(db *sql.DB, articleID int) ([]models.Comment, error) {
	const sqlStr = `
		SELECT comment_id, article_id, message, created_at
		FROM comments
		WHERE article_id = ?
	`

	rows, err := db.Query(sqlStr, articleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var commentArray []models.Comment

	for rows.Next() {
		var comment models.Comment
		var createdTime sql.NullTime

		err = rows.Scan(
			&comment.CommentID,
			&comment.ArticleID,
			&comment.Message,
			&createdTime,
		)

		if createdTime.Valid {
			comment.CreatedAt = createdTime.Time
		}
		commentArray = append(commentArray, comment)
	}

	return commentArray, nil
}
