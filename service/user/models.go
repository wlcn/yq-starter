package user

import (
	"fmt"
	"log"
	"time"

	"github.com/wlcn/yq-starter/service/common"

	"github.com/wlcn/yq-starter/helper"

	"github.com/jinzhu/gorm"

	"golang.org/x/crypto/bcrypt"
)

// User should only be concerned with database schema, more strict checking should be put in validator.
type User struct {
	gorm.Model
	Name         string `gorm:"type:varchar(100);unique_index"`
	Age          *int
	Birthday     time.Time
	Email        string `gorm:"type:varchar(100);unique_index"`
	Role         string `gorm:"size:255"`        // set field size to 255
	MemberNumber string `gorm:"unique;not null"` // set member number to unique and not null
	Address      string `gorm:"index:addr"`      // create index with name `addr` for address
	Password     string `gorm:"column:password;not null"`
}

// SetPassword 设置密码加密
func (u *User) SetPassword(password string) error {
	if len(password) == 0 {
		return fmt.Errorf("Password should not be empty, actually is [%s]", password)
	}
	bytePassword := []byte(password)
	// Make sure the second param `bcrypt generator cost` between [4, 32)
	passwordHash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	u.Password = string(passwordHash)
	return nil
}

// CheckPassword 验证密码
func (u *User) CheckPassword(password string) error {
	bytePassword := []byte(password)
	byteHashedPassword := []byte(u.Password)
	log.Printf("password is %v, passwordHash is %v", password, u.Password)
	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}

// FindOne 查询
func FindOne(condition interface{}) (User, error) {
	var model User
	err := helper.DB.Where(condition).First(&model).Error
	return model, err
}

// FindCondition 分页条件查询
func FindCondition(condition interface{}, page common.Page) (interface{}, error) {
	var result = make(map[string]interface{}, 0)
	var models []User
	var count int
	err := helper.DB.Where(condition).Limit(page.Size).Offset((page.Page - 1) * page.Size).Order(page.Order).Find(&models).Count(&count).Error
	if err != nil {
		log.Printf("Database query error is %+v", err)
		return result, err
	}
	// log.Printf("models is %+v, count is %v", models, count)
	result["total"] = count
	result["users"] = models
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
// For below Update, nothing will be updated as "", 0, false are blank values of their types
// db.Model(&user).Updates(User{Name: "", Age: 0, Actived: false})
func PatchOne(data interface{}) error {
	var user User
	err := helper.DB.Model(&user).Updates(data).Error
	return err
}

// DeleteOne 软删除
func DeleteOne(data interface{}) error {
	err := helper.DB.Delete(data).Error
	return err
}
