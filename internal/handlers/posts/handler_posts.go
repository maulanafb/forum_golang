package posts

import (
	"github.com/gin-gonic/gin"
	"github.com/maulanafb/forum_golang/internal/middleware"
	"github.com/maulanafb/forum_golang/internal/model/posts"
	"golang.org/x/net/context"
)

type postService interface {
	CreatePost(ctx context.Context, UserID int64, req posts.CreatePostRequest) error
	CreateComment(ctx context.Context, postID, userId int64, request posts.CreateCommentRequest) error
}

type Handler struct {
	*gin.Engine
	postSvc postService
}

func NewHandler(api *gin.Engine, postSvc postService) *Handler {
	return &Handler{
		Engine:  api,
		postSvc: postSvc,
	}
}

func (h *Handler) RegisterRoutes() {
	route := h.Group("/posts")
	route.Use(middleware.AuthMiddleware())
	route.POST("/create-post", h.CreatePost)
	route.POST("/create-comment/:postID", h.CreateComment)
}
