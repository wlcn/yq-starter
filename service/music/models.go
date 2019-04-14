package music

import (
	"time"

	"github.com/wlcn/yq-starter/service/common"

	"github.com/wlcn/yq-starter/helper"

	"github.com/jinzhu/gorm"
)

// Music should only be concerned with database schema, more strict checking should be put in validator.
type Music struct {
	gorm.Model
	SourceID    string
	URL         string
	Title       string `gorm:"type:varchar(512)"`
	Content     string `gorm:"type:text"`
	PublishTime time.Time
	Author      string
	Source      string
	Tag         string
	LrcLink     string
	PicLink     string
	FileLink    string
	FilePath    string
}

// FindOne 查询
func FindOne(condition interface{}) (Music, error) {
	var model Music
	err := helper.DB.Where(condition).First(&model).Error
	return model, err
}

// FindCondition 分页条件查询
func FindCondition(condition interface{}, page common.Page) (interface{}, error) {
	var result = make(map[string]interface{}, 0)
	var models []Music
	var count int
	err := helper.DB.Where(condition).Find(&models).Count(&count).Error
	if err != nil {
		return result, err
	}
	err = helper.DB.Where(condition).Limit(page.Size).Offset((page.Page - 1) * page.Size).Order(page.Order).Find(&models).Error
	if err != nil {
		return result, err
	}
	// log.Printf("models is %+v, count is %v", models, count)
	result["total"] = count
	result["musics"] = models
	return result, err
}

// SaveOne 保存
func SaveOne(data interface{}) error {
	err := helper.DB.Save(data).Error
	return err
}

// UpdateOne 全量更新对象
func UpdateOne(data interface{}) error {
	err := helper.DB.Save(data).Error
	return err
}

// PatchOne 局部更新对象，只有传值并且传的是真值才会更新
// WARNING when update with struct, GORM will only update those fields that with non blank value
func PatchOne(data interface{}) error {
	var music Music
	err := helper.DB.Model(&music).Updates(data).Error
	return err
}

// DeleteOne 软删除
func DeleteOne(data interface{}) error {
	err := helper.DB.Delete(data).Error
	return err
}
