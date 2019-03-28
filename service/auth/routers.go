package auth

import (
	"net/http"

	"github.com/wlcn/yq-starter/helper"

	"github.com/wlcn/yq-starter/service/user"

	"github.com/gin-gonic/gin"
)

// Routers 路由
func Routers(r *gin.RouterGroup) {
	r.POST("/login", Login)
	r.POST("/logout", Logout)
}

// Login logic
func Login(c *gin.Context) {
	var condition user.User
	if err := c.Bind(&condition); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var user user.User
	err := helper.DB.Where(condition).First(&user).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := helper.GenerateToken(user.Name, user.Password)
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

// Logout logic
func Logout(c *gin.Context) {
	var user user.User
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}
