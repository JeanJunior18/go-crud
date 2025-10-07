package main

import (
	"log"

	"github.com/JeanJunior18/go-crud/internal/core"
	"github.com/JeanJunior18/go-crud/internal/handler"
	mongodb "github.com/JeanJunior18/go-crud/internal/infrastructure/mongoDb"
	"github.com/gin-gonic/gin"
)

func main() {
	port := "localhost:3000"
	postRepo := mongodb.New()
	postService := core.New(postRepo)
	postHandler := handler.New(postService)

	router := gin.Default()

	router.POST("/posts", postHandler.CreatePostHandler)

	log.Printf("🚀 Servidor iniciado e escutando na porta %s", port)
	if err := router.Run(port); err != nil {
		log.Fatalf("❌ Falha ao iniciar o servidor Gin: %v", err)
	}
}
