package user

import (
	"net/http"

	"github.com/wlcn/yq-starter/helper"

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
		c.JSON(http.StatusBadRequest, gin.H{
			helper.Code:  http.StatusBadRequest,
			helper.Error: err.Error(),
		})
		return
	}
	if page.Order == "" {
		page.Order = helper.Order
	}
	var user User
	if err := c.BindQuery(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			helper.Code:  http.StatusBadRequest,
			helper.Error: err.Error(),
		})
		return
	}
	// log.Printf("page is %+v, user is %+v", page, user)
	result, err := FindCondition(&user, page)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			helper.Code:  http.StatusNotFound,
			helper.Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		helper.Code: http.StatusOK,
		helper.Data: result,
	})
}

// Save logic
func Save(c *gin.Context) {
	var user User
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			helper.Code:  http.StatusBadRequest,
			helper.Error: err.Error(),
		})
		return
	}
	if user.Password != "" {
		user.SetPassword(user.Password)
	}
	if err := SaveOne(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			helper.Code:  http.StatusUnprocessableEntity,
			helper.Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		helper.Code: http.StatusOK,
		helper.Data: user,
	})
}

// Update logic 全部字段更新，如果没有传值则值为空
func Update(c *gin.Context) {
	var user User
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			helper.Code:  http.StatusBadRequest,
			helper.Error: err.Error(),
		})
		return
	}
	if user.Password != "" {
		user.SetPassword(user.Password)
	}
	if err := UpdateOne(user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			helper.Code:  http.StatusUnprocessableEntity,
			helper.Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		helper.Code: http.StatusOK,
		helper.Data: user,
	})
}

// Patch logic 局部字段更新,只更新传入的有效值, 0,false,nil,空字符串不更新
func Patch(c *gin.Context) {
	var user User
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			helper.Code:  http.StatusBadRequest,
			helper.Error: err.Error(),
		})
		return
	}
	if user.Password != "" {
		user.SetPassword(user.Password)
	}
	if err := PatchOne(user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			helper.Code:  http.StatusUnprocessableEntity,
			helper.Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		helper.Code: http.StatusOK,
		helper.Data: user,
	})
}

// Delete logic
func Delete(c *gin.Context) {
	var user User
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			helper.Code:  http.StatusBadRequest,
			helper.Error: err.Error(),
		})
		return
	}
	if err := DeleteOne(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			helper.Code:  http.StatusUnprocessableEntity,
			helper.Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		helper.Code: http.StatusOK,
		helper.Data: user,
	})
}
