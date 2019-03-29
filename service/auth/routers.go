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
	r.POST("/register", Register)
	r.POST("/logout", Logout)
}

// Login logic
func Login(c *gin.Context) {
	var condition user.User
	if err := c.Bind(&condition); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			helper.Code:  http.StatusBadRequest,
			helper.Error: err.Error(),
		})
		return
	}
	// 查询用户
	var user user.User
	err := helper.DB.Where("name = ?", condition.Name).First(&user).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			helper.Code:  http.StatusBadRequest,
			helper.Error: err.Error(),
		})
		return
	}
	// 验证密码
	err = user.CheckPassword(condition.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			helper.Code:  http.StatusBadRequest,
			helper.Error: err.Error(),
		})
		return
	}
	token, err := helper.GenerateToken(user.Name, user.Password)
	c.JSON(http.StatusOK, gin.H{
		helper.Code:  http.StatusOK,
		helper.Token: token,
	})
}

// Register logic
func Register(c *gin.Context) {
	var u user.User
	if err := c.Bind(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			helper.Code:  http.StatusBadRequest,
			helper.Error: err.Error(),
		})
		return
	}
	if u.Password != "" {
		u.SetPassword(u.Password)
	}
	if err := user.SaveOne(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			helper.Code:  http.StatusUnprocessableEntity,
			helper.Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		helper.Code: http.StatusOK,
		helper.Data: u,
	})
}

// Logout logic
func Logout(c *gin.Context) {
	var user user.User
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			helper.Code:  http.StatusBadRequest,
			helper.Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		helper.Code: http.StatusOK,
	})
}
