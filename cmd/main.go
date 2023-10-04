package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/ImranZahoor/blog-api/internal/controller"
	"github.com/ImranZahoor/blog-api/internal/repository"
	"github.com/ImranZahoor/blog-api/internal/router"
	"github.com/ImranZahoor/blog-api/internal/service"
)

const (
	PORT = 5000
)

func main() {
	repo := repository.NewRepository()
	service := service.NewService(repo)
	controller := controller.NewController(service)
	server := router.NewServer(controller)
	server.RegisterHandlers()
	router := server.GetRouter()
	if err := http.ListenAndServe(fmt.Sprintf(":%s", strconv.Itoa(PORT)), router); err != nil {
		fmt.Printf("server startup failed %v", err)
		os.Exit(0)
	}
}
