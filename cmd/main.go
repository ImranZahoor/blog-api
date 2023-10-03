package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/ImranZahoor/blog-api/internal/router"
)

const (
	PORT = 5000
)

func main() {
	server := router.NewServer()
	server.RegisterHandlers()
	router := server.GetRouter()
	if err := http.ListenAndServe(fmt.Sprintf(":%s", strconv.Itoa(PORT)), router); err != nil {
		fmt.Printf("server startup failed %v", err)
		os.Exit(0)
	}
}
