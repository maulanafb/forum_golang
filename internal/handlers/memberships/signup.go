package memberships

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maulanafb/forum_golang/internal/model/memberships"
)

func (h *Handler) SignUp(c *gin.Context) {
	ctx := c.Request.Context()
	var request memberships.SignUpRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Printf("Error binding JSON: %v\n", err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	log.Printf("SignUp request: %+v\n", request)

	err := h.membershipService.SignUp(ctx, request)
	if err != nil {
		log.Printf("Error in membershipService.SignUp: %v\n", err)
		c.JSON(500, gin.H{"error di handlers": err.Error()})
		return
	}

	log.Println("User successfully created")
	c.Status(http.StatusCreated)
}
