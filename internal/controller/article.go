package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ImranZahoor/blog-api/internal/models"
)

func (c *Controller) ListArticleHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	articles, err := c.service.ListArticle(ctx)
	if err != nil {

		jsonResponse(w, http.StatusBadRequest, map[string]string{"error": "articles not found"})
	}
	response, _ := json.Marshal(articles)
	_, _ = w.Write(response)
}
func (c *Controller) GetArticleByIDHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "get article handler")
}
func (c *Controller) CreateArticleHandler(w http.ResponseWriter, r *http.Request) {
	var article models.Article
	jsonDecoder := json.NewDecoder(r.Body)
	err := jsonDecoder.Decode(&article)
	if err != nil {
		jsonResponse(w, http.StatusBadRequest, map[string]string{"error": "invalid payload"})
	}
	err = c.service.CreateArticle(r.Context(), article)

	if err != nil {

		jsonResponse(w, http.StatusBadRequest, err)
	}
	fmt.Fprintf(w, "create article handler testing nodemod\n")
}
func (c *Controller) DeleteArticleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "delete article handler")
}
func (c *Controller) UpdateArticleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "update article handler")
}

func jsonResponse(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, _ = w.Write(response)
}
