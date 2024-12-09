// backend/controllers/user_controller.go
package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
    service *UserService
}

func NewUserController(service *UserService) *UserController {
    return &UserController{service}
}

type RegisterInput struct {
    Name     string `json:"name" binding:"required"`
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required,min=6"`
}

type LoginInput struct {
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required"`
}

func (c *UserController) Register(ctx *gin.Context) {
    var input RegisterInput
    if err := ctx.ShouldBindJSON(&input); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user, err := c.service.Register(input.Name, input.Email, input.Password)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusCreated, gin.H{"message": "registration successful", "user": user})
}

func (c *UserController) Login(ctx *gin.Context) {
    var input LoginInput
    if err := ctx.ShouldBindJSON(&input); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    token, err := c.service.Login(input.Email, input.Password)
    if err != nil {
        ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func (c *UserController) GetUser(ctx *gin.Context) {
    idParam := ctx.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
        return
    }

    user, err := c.service.GetUser(uint(id))
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"user": user})
}
