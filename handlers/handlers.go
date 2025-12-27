package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/nissi1278/go-api-practice/models"
)

func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	var reqArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
		return
	}

	article := reqArticle
	json.NewEncoder(w).Encode(article)
}

func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	// queryMap := req.URL.Query()
	// var page int
	// if p, ok := queryMap["page"]; ok && len(p) > 0 {
	// 	var err error
	// 	page, err = strconv.Atoi(p[0])
	// 	if err != nil {
	// 		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
	// 		return
	// 	}
	// } else {
	// 	page = 1
	// }

	articles := []models.Article{models.Article1, models.Article2}
	json.NewEncoder(w).Encode(articles)
}

func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}

	articles := models.Article1
	jsonData, err := json.Marshal(articles)
	if err != nil {
		errMsg := fmt.Sprintf("fail to encode json (articleID:%d)\n", articleID)
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	var reqArticle models.Article
	err := json.NewDecoder(req.Body).Decode(&reqArticle)

	if err != nil {
		http.Error(w, "failed to decode json\n", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(reqArticle)
}

func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	var reqComment models.Comment
	err := json.NewDecoder(req.Body).Decode(&reqComment)
	if err != nil {
		http.Error(w, "failed to decode json\n", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(reqComment)
}
