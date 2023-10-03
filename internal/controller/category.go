package controller

import (
	"fmt"
	"net/http"
)

func (c *Controller) ListCategoryHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "list article handler")
}
func (c *Controller) GetCategoryByIDHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "get article handler")
}
func (c *Controller) CreateCategoryHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "create article handler")
}
func (c *Controller) DeleteCategoryHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "delete article handler")
}
func (c *Controller) UpdateCategoryHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "update article handler")
}
