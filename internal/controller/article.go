package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ImranZahoor/blog-api/internal/models"
	"github.com/ImranZahoor/blog-api/pkg/util"
	"github.com/gorilla/mux"
)

func (c *Controller) ListArticleHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	articles, err := c.service.ListArticle(ctx)
	if err != nil {

		util.JsonResponse(w, http.StatusBadRequest, map[string]string{"error": "articles not found"})
	}
	response, _ := json.Marshal(articles)
	_, _ = w.Write(response)
}
func (c *Controller) GetArticleByIDHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	id := vars["id"]
	intId, err := strconv.Atoi(id)

	if err != nil {
		util.JsonResponse(w, http.StatusBadRequest, map[string]string{"error": "id conversion error"})
	}
	article, err := c.service.GetArticleByID(ctx, models.Uuid(intId))
	if err != nil {
		util.JsonResponse(w, http.StatusBadRequest, map[string]string{"error": "articles not found"})
	}
	response, _ := json.Marshal(article)
	_, _ = w.Write(response)
}
func (c *Controller) CreateArticleHandler(w http.ResponseWriter, r *http.Request) {
	var article models.Article
	jsonDecoder := json.NewDecoder(r.Body)
	err := jsonDecoder.Decode(&article)
	if err != nil {
		util.JsonResponse(w, http.StatusBadRequest, map[string]string{"error": "invalid payload"})
	}
	err = c.service.CreateArticle(r.Context(), article)

	if err != nil {
		util.JsonResponse(w, http.StatusBadRequest, err)
	}
	util.JsonResponse(w, http.StatusCreated, map[string]string{"message": "article created"})
}
func (c *Controller) DeleteArticleHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	intId, err := util.ToUUID(id)
	if err != nil {
		util.JsonResponse(w, http.StatusBadRequest, map[string]string{"error": "id conversion error"})
	}
	ctx := r.Context()
	err = c.service.DeleteArticle(ctx, intId)
	if err != nil {
		util.JsonResponse(w, http.StatusBadRequest, map[string]string{"error": "deletion error"})
	}
	util.JsonResponse(w, http.StatusOK, map[string]string{"message": "deleted successfully"})
}
func (c *Controller) UpdateArticleHandler(w http.ResponseWriter, r *http.Request) {

	var article models.Article
	jsonDecoder := json.NewDecoder(r.Body)
	err := jsonDecoder.Decode(&article)
	if err != nil {
		util.JsonResponse(w, http.StatusBadRequest, map[string]string{"error": "invalid payload"})
	}
	vars := mux.Vars(r)
	id := vars["id"]
	intId, err := util.ToUUID(id)
	if err != nil {
		util.JsonResponse(w, http.StatusBadRequest, map[string]string{"error": "id conversion error"})
	}
	ctx := r.Context()
	err = c.service.UpdateArticle(ctx, intId, article)
	if err != nil {
		util.JsonResponse(w, http.StatusBadRequest, map[string]string{"error": "update failed"})
	}

	util.JsonResponse(w, http.StatusOK, map[string]string{"message": "updated successfully"})
}
