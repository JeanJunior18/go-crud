package main

import (
	"log"

	"github.com/JeanJunior18/go-crud/internal/core"
	"github.com/JeanJunior18/go-crud/internal/handler"
	mongodb "github.com/JeanJunior18/go-crud/internal/infrastructure/mongoDb"
	"github.com/gin-gonic/gin"
)

func main() {
	client, dbName := mongodb.InitConnection()

	port := ":3000"
	postRepo := mongodb.New(client, dbName)
	postService := core.New(postRepo)
	postHandler := handler.New(postService)

	router := gin.Default()
	gin.SetMode(gin.ReleaseMode)

	router.POST("/posts", postHandler.CreatePostHandler)
	router.GET("/posts", postHandler.FindPostsHandler)

	log.Printf("🚀 Servidor iniciado e escutando na porta %s", port)
	if err := router.Run(port); err != nil {
		log.Fatalf("❌ Falha ao iniciar o servidor Gin: %v", err)
	}
}
