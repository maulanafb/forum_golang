package posts

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/maulanafb/forum_golang/internal/model/posts"
	"net/http"
	"strconv"
)

func (h *Handler) CreateComment(c *gin.Context) {
	ctx := c.Request.Context()
	var request posts.CreateCommentRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	postIDStr := c.Param("postID")
	postID, err := strconv.ParseInt(postIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errors.New("invalid postID"),
		})
		return
	}

	userID := c.GetInt64("userID")

	err = h.postSvc.CreateComment(ctx, postID, userID, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.Status(http.StatusCreated)
}
