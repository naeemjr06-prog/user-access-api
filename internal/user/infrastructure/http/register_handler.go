package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/naeemjr06-prog/user-access-api/internal/user/application"
)

type RegisterHandler struct {
	userHandler  *application.RegisterUserHandler
	adminHandler *application.RegisterUserAdminHandler
}

func NewRegisterHandler(
	userHandler *application.RegisterUserHandler,
	adminHandler *application.RegisterUserAdminHandler,
) *RegisterHandler {
	return &RegisterHandler{
		userHandler:  userHandler,
		adminHandler: adminHandler,
	}
}

type registerRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *RegisterHandler) Register(c *gin.Context) {
	var req registerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	user, err := h.userHandler.Handle(c.Request.Context(), application.RegisterUserInput{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (h *RegisterHandler) RegisterAdmin(c *gin.Context) {
	// TODO: Later → extract role from JWT middleware
	role := c.GetHeader("X-Role")
	if role != "ADMIN" {
		c.JSON(http.StatusForbidden, gin.H{"error": "admin access required"})
		return
	}

	var req registerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	user, err := h.adminHandler.HandleAdmin(c.Request.Context(), application.RegisterUserAdminInput{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}
