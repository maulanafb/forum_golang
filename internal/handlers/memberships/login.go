package memberships

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maulanafb/forum_golang/internal/model/memberships"
)

func (h *Handler) Login(c *gin.Context) {
	ctx := c.Request.Context()
	var request memberships.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Printf("Error binding JSON: %v\n", err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	log.Printf("SignUp request: %+v\n", request)

	accessToken, err := h.membershipService.Login(ctx, request)
	if err != nil {
		log.Printf("Error in membershipService.SignUp: %v\n", err)
		c.JSON(500, gin.H{"error di handlers": err.Error()})
		return
	}

	log.Println("User successfully created")
	response := memberships.LoginResponse{
		AccessToken: accessToken,
	}
	c.JSON(http.StatusCreated, response)

}
