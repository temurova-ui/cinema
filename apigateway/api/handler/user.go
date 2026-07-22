package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	userpb "github.com/temurova-ui/cinema/userservice/userpb/v1"
)

func (h *handler) GetUser(c *gin.Context) {
	idStr := c.Param("user_id")
	if idStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id is required"})
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id is required"})
		return
	}

	response, err := h.serviceManager.UserService().GetByID(c.Request.Context(), &userpb.GetUserRequest{
		Id: id,
	})

	if err != nil {
		log.Println("handler GetUser h.serviceManager.UserService().GetByID", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	c.JSON(http.StatusOK, response)
}