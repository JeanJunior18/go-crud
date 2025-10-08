package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/JeanJunior18/go-crud/internal/core"
	"github.com/JeanJunior18/go-crud/internal/handler/dto"
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
	var bodyReq dto.CreatePostRequest

	if err := ctx.ShouldBindJSON(&bodyReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid Payload",
			"details": err.Error(),
		})
		return
	}

	serviceCtx, cancel := context.WithTimeout(ctx.Request.Context(), 5*time.Second)
	defer cancel()

	createdPost, err := c.service.CreatePost(serviceCtx, bodyReq)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to create post",
			"details": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, createdPost)
}

func (c *PostController) FindPostsHandler(ctx *gin.Context) {
	serviceCtx, cancel := context.WithTimeout(ctx.Request.Context(), 5*time.Second)
	defer cancel()

	posts, err := c.service.FindPost(serviceCtx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to list posts",
			"details": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, posts)
}

func (c *PostController) FindPostByIdHandler(ctx *gin.Context) {
	serviceCtx, cancel := context.WithTimeout(ctx.Request.Context(), 5*time.Second)
	defer cancel()

	id := ctx.Param("id")
	post, err := c.service.FindById(serviceCtx, id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Failed to get post by id",
			"detail": err.Error(),
		})
		return
	}

	if post == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "ID not found",
			"id":    id,
		})
	}

	ctx.JSON(http.StatusOK, *post)
}
