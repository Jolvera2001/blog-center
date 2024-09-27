package handlers

import "github.com/gin-gonic/gin"

func GroupUserHandlers(r *gin.Engine) {
	v1 := r.Group("api/v1") 
	{
		v1.GET("user/:userid", GetUserProfile)
		v1.POST("user/create", CreateUserProfile)
		v1.PUT("user/update", UpdateUserProfile)
		v1.DELETE("user/delete", DeleteUserProfile)
	}
}

func CreateUserProfile(c *gin.Context) {
	
}

func GetUserProfile(c *gin.Context) {

}

func UpdateUserProfile(c *gin.Context) {

}

func DeleteUserProfile(c *gin.Context) {

}

