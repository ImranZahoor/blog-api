package controller

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/ImranZahoor/blog-api/internal/models"
	"github.com/ImranZahoor/blog-api/pkg/util"
	"github.com/gorilla/mux"
)

func (c *Controller) ListCategoryHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	categories, err := c.service.ListCategory(ctx)

	if err != nil {
		log.Printf("Controller::Category::ListCategoryHandler : %s", err.Error())
		if err == io.EOF {
			util.JsonResponse(w, http.StatusNotFound, map[string]string{"error": "No Data Found"})
			return
		}
		util.JsonResponse(w, http.StatusNotFound, map[string]string{"error": err.Error()})
		return
	}

	util.JsonResponse(w, http.StatusOK, categories)
}
func (c *Controller) GetCategoryByIDHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	id := vars["id"]
	intId, err := strconv.Atoi(id)

	if err != nil {
		util.JsonResponse(w, http.StatusBadRequest, map[string]string{"error": "invalid id"})
		return
	}
	category, err := c.service.GetCategoryByID(ctx, models.Uuid(intId))
	if err != nil {
		util.JsonResponse(w, http.StatusBadRequest, map[string]string{"error": "articles not found"})
		return
	}
	util.JsonResponse(w, http.StatusOK, category)
}
func (c *Controller) CreateCategoryHandler(w http.ResponseWriter, r *http.Request) {
	var cateory models.Category
	jsonDecoder := json.NewDecoder(r.Body)
	err := jsonDecoder.Decode(&cateory)
	if err != nil {
		log.Printf("Controller::Category::CreateCategoryHandler: %s", err.Error())
		util.JsonResponse(w, http.StatusBadRequest, map[string]string{"error": "invalid payload"})
		return
	}
	err = c.service.CreateCategory(r.Context(), cateory)

	if err != nil {
		log.Printf("Controller::Category::CreateCategoryHandler: %s", err.Error())
		util.JsonResponse(w, http.StatusBadRequest, map[string]string{"error": "failed to create category"})
		return
	}
	util.JsonResponse(w, http.StatusCreated, map[string]string{"message": "category created"})
}
func (c *Controller) DeleteCategoryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	intId, err := util.ToUUID(id)
	if err != nil {
		util.JsonResponse(w, http.StatusBadRequest, map[string]string{"error": "id conversion error"})
		return
	}
	ctx := r.Context()
	err = c.service.DeleteCategory(ctx, intId)
	if err != nil {
		util.JsonResponse(w, http.StatusBadRequest, map[string]string{"error": "deletion error"})
		return
	}
	util.JsonResponse(w, http.StatusOK, map[string]string{"message": "deleted successfully"})

}
func (c *Controller) UpdateCategoryHandler(w http.ResponseWriter, r *http.Request) {

	var category models.Category
	jsonDecoder := json.NewDecoder(r.Body)
	err := jsonDecoder.Decode(&category)
	if err != nil {
		util.JsonResponse(w, http.StatusBadRequest, map[string]string{"error": "invalid payload"})
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]
	intId, err := util.ToUUID(id)
	if err != nil {
		util.JsonResponse(w, http.StatusBadRequest, map[string]string{"error": "id conversion error"})
	}
	ctx := r.Context()
	err = c.service.UpdateCategory(ctx, intId, category)
	if err != nil {
		util.JsonResponse(w, http.StatusBadRequest, map[string]string{"error": "category update failed"})
		return
	}

	util.JsonResponse(w, http.StatusOK, map[string]string{"message": "updated successfully"})
}
