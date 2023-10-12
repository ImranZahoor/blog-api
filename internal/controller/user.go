package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ImranZahoor/blog-api/internal/models"
	"github.com/ImranZahoor/blog-api/pkg/util"
	"github.com/gorilla/mux"
)

func (c *Controller) ListUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	users, err := c.service.ListUsers(ctx)
	if err != nil {
		util.JsonResponse(w, http.StatusBadRequest, map[string]string{"error": "users not found"})
		return
	}

	response, _ := json.Marshal(users)
	_, _ = w.Write(response)
}

func (c *Controller) GetUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	id := vars["id"]
	intID, err := strconv.Atoi(id)
	if err != nil {
		util.JsonResponse(w, http.StatusBadRequest, map[string]string{"error": "id conversion error"})
		return
	}

	user, err := c.service.GetUserByID(ctx, models.Uuid(intID))
	if err != nil {
		util.JsonResponse(w, http.StatusBadRequest, map[string]string{"error": "user not found"})
		return
	}

	response, _ := json.Marshal(user)
	_, _ = w.Write(response)
}

func (c *Controller) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	jsonDecoder := json.NewDecoder(r.Body)
	err := jsonDecoder.Decode(&user)
	if err != nil {
		util.JsonResponse(w, http.StatusBadRequest, map[string]string{"error": "invalid payload"})
		return
	}

	err = c.service.CreateUser(r.Context(), user)
	if err != nil {
		util.JsonResponse(w, http.StatusBadRequest, map[string]string{"error": "user creation failed"})
		return
	}

	util.JsonResponse(w, http.StatusCreated, map[string]string{"message": "user created"})
}

func (c *Controller) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	intID, err := util.ToUUID(id)
	if err != nil {
		util.JsonResponse(w, http.StatusBadRequest, map[string]string{"error": "id conversion error"})
		return
	}

	ctx := r.Context()
	err = c.service.DeleteUser(ctx, intID)
	if err != nil {
		util.JsonResponse(w, http.StatusBadRequest, map[string]string{"error": "user deletion failed"})
		return
	}

	util.JsonResponse(w, http.StatusOK, map[string]string{"message": "user deleted successfully"})
}

func (c *Controller) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	jsonDecoder := json.NewDecoder(r.Body)
	err := jsonDecoder.Decode(&user)
	if err != nil {
		util.JsonResponse(w, http.StatusBadRequest, map[string]string{"error": "invalid payload"})
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]
	intID, err := util.ToUUID(id)
	if err != nil {
		util.JsonResponse(w, http.StatusBadRequest, map[string]string{"error": "id conversion error"})
		return
	}

	ctx := r.Context()
	err = c.service.UpdateUser(ctx, intID, user)
	if err != nil {
		util.JsonResponse(w, http.StatusBadRequest, map[string]string{"error": "user update failed"})
		return
	}

	util.JsonResponse(w, http.StatusOK, map[string]string{"message": "user updated successfully"})
}
