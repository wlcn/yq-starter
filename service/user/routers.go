package user

import (
	"fmt"
	"net/http"

	"github.com/wlcn/yq-starter/service/common"

	"github.com/gin-gonic/gin"
)

// Routers 用户路由
func Routers(r *gin.RouterGroup) {
	r.GET("/", Find)
	r.POST("/", Save)
	r.PUT("/", Update)
	r.PATCH("/", Patch)
	r.DELETE("/", Delete)
}

// Find logic
// GET /user?page=1&size=1&id=1&name=wl
func Find(c *gin.Context) {
	// 获取分页参数，必输项
	var page common.Page
	if err := c.BindQuery(&page); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if page.Order == "" {
		page.Order = " id DESC"
	}
	var user User
	if err := c.BindQuery(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// log.Printf("page is %+v, user is %+v", page, user)
	result, err := FindCondition(&user, page)
	if err != nil {
		c.JSON(http.StatusNotFound, fmt.Sprintf("Database error %+v", err))
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": result})
}

// Save logic
func Save(c *gin.Context) {
	var user User
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.setPassword(user.Password)
	if err := SaveOne(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, fmt.Sprintf("Database error %+v", err))
		return
	}
	c.JSON(http.StatusCreated, gin.H{"user": user})
}

// Update logic 全部更新
func Update(c *gin.Context) {
	var user User
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := UpdateOne(user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, fmt.Sprintf("Database error is %+v", err))
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}

// Patch logic 局部更新
func Patch(c *gin.Context) {
	var user User
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := PatchOne(user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, fmt.Sprintf("Database error is %+v", err))
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}

// Delete logic
func Delete(c *gin.Context) {
	var user User
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := DeleteOne(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, fmt.Sprintf("Database error %+v", err))
		return
	}
	c.JSON(http.StatusCreated, gin.H{"user": user})
}
