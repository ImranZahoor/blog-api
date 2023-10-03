package controller

import (
	"fmt"
	"net/http"
)

func (c *Controller) ListUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "list user handler")
}
func (c *Controller) GetUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "get user handler")
}
func (c *Controller) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "create user handler")
}
func (c *Controller) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "delete user handler")
}
func (c *Controller) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "update user handler")
}
