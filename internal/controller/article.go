package controller

import (
	"fmt"
	"net/http"
)

func (c *Controller) ListArticleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "list article handler")
}
func (c *Controller) GetArticleByIDHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "get article handler")
}
func (c *Controller) CreateArticleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "create article handler")
}
func (c *Controller) DeleteArticleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "delete article handler")
}
func (c *Controller) UpdateArticleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "update article handler")
}
