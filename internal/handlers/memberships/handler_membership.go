package memberships

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/maulanafb/forum_golang/internal/model/memberships"
)

type membershipService interface {
	SignUp(ctx context.Context, req memberships.SignUpRequest) error
}

type Handler struct {
	*gin.Engine

	membershipService membershipService
}

func NewHandler(api *gin.Engine, membershipSvc membershipService) *Handler {
	return &Handler{
		Engine:            api,
		membershipService: membershipSvc,
	}
}

func (h *Handler) RegisterRoutes() {
	route := h.Group("/memberships")
	route.GET("/ping", h.ping)
	route.POST("/sign-up", h.SignUp)
}
