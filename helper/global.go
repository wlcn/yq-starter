package helper

import (
	"sync"

	"github.com/jinzhu/gorm"
)

var (
	// Once is to init something
	Once sync.Once
	// DB is to access database
	DB *gorm.DB
)
