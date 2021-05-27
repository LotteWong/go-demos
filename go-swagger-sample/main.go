package main

import (
	"go-swagger-sample/router"
	"log"
)

// @title Title For Go-Swagger-Sample Api Docs
// @description `Markdown` Description For Go-Swagger-Sample Api Docs
// @version 1.0.0
// @host 127.0.0.1:8080
// @BasePath /api/v1
func main() {
	route := router.InitRouter()
	if err := route.Run(); err != nil {
		log.Fatalf("App crashed, err: %v", err)
	}
}
