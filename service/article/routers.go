package article

import (
	"net/http"

	"github.com/wlcn/yq-starter/helper"

	"github.com/wlcn/yq-starter/service/common"

	"github.com/gin-gonic/gin"
)

// Routers 路由
func Routers(r *gin.RouterGroup) {
	r.GET("/", Find)
	r.POST("/", Save)
	r.PUT("/", Update)
	r.PATCH("/", Patch)
	r.DELETE("/", Delete)
}

// Find logic
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
	var article Article
	if err := c.BindQuery(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			helper.Code:  http.StatusBadRequest,
			helper.Error: err.Error(),
		})
		return
	}
	article.DeletedAt = nil
	// log.Printf("page is %+v, article is %+v", page, article)
	result, err := FindCondition(article, page)
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
	var article Article
	if err := c.Bind(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			helper.Code:  http.StatusBadRequest,
			helper.Error: err.Error(),
		})
		return
	}
	if err := SaveOne(&article); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			helper.Code:  http.StatusUnprocessableEntity,
			helper.Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		helper.Code: http.StatusOK,
		helper.Data: article,
	})
}

// Update logic 全部字段更新，如果没有传值则值为空
func Update(c *gin.Context) {
	var article Article
	if err := c.Bind(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			helper.Code:  http.StatusBadRequest,
			helper.Error: err.Error(),
		})
		return
	}
	if err := UpdateOne(article); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			helper.Code:  http.StatusUnprocessableEntity,
			helper.Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		helper.Code: http.StatusOK,
		helper.Data: article,
	})
}

// Patch logic 局部字段更新,只更新传入的有效值, 0,false,nil,空字符串不更新
func Patch(c *gin.Context) {
	var article Article
	if err := c.Bind(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			helper.Code:  http.StatusBadRequest,
			helper.Error: err.Error(),
		})
		return
	}
	if err := PatchOne(article); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			helper.Code:  http.StatusUnprocessableEntity,
			helper.Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		helper.Code: http.StatusOK,
		helper.Data: article,
	})
}

// Delete logic
func Delete(c *gin.Context) {
	var article Article
	if err := c.Bind(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			helper.Code:  http.StatusBadRequest,
			helper.Error: err.Error(),
		})
		return
	}
	if err := DeleteOne(&article); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			helper.Code:  http.StatusUnprocessableEntity,
			helper.Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		helper.Code: http.StatusOK,
		helper.Data: article,
	})
}
