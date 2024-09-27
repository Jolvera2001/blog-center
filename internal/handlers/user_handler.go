package handlers

import (
	"blog-center/internal/dtos"
	"blog-center/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService service.IUserService
}


func (h *UserHandler) CreateUserProfile(c *gin.Context) {
	var dto dtos.UserDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "invalid request"})
		return
	}

	userID, err := h.UserService.RegisterUser(dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "error with internal process"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": userID})
}

func (h *UserHandler) GetUserProfile(c *gin.Context) {
	userId := c.Param("userid")
	
	res, err := h.UserService.GetUserProfile(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "error with internal process"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"profile": res})
}

func (h *UserHandler) UpdateUserProfile(c *gin.Context) {
	type updateDto struct {
		Uuid string `json:"id"`
		UserDto dtos.UserDto
	}
	var dto updateDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "invalid request"})
		return
	}

	err := h.UserService.UpdateUserProfile(dto.Uuid, dto.UserDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "error with internal process"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Status": "Success"})
}

func (h *UserHandler) DeleteUserProfile(c *gin.Context) {
	var id string
	if err := c.ShouldBindJSON(&id); err != nil {

	}

	err := h.UserService.DeleteUserAccount(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "error with internal process"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Status": "Deleted"})
}

func GroupUserHandlers(r *gin.Engine, h *UserHandler) {
	v1 := r.Group("api/v1") 
	{
		v1.GET("user/:userid", h.GetUserProfile)
		v1.POST("user/create", h.CreateUserProfile)
		v1.PUT("user/update", h.UpdateUserProfile)
		v1.DELETE("user/delete", h.DeleteUserProfile)
	}
}