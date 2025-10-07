package handler

import (
	"fmt"

	"github.com/JeanJunior18/go-crud/internal/core"
	"github.com/gin-gonic/gin"
)

type PostController struct {
	service core.PostServiceContract
}

func New(s core.PostServiceContract) *PostController {
	return &PostController{
		service: s,
	}
}

func (c *PostController) CreatePostHandler(ctx *gin.Context) {
	fmt.Printf("Received request")
}
