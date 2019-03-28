package helper

import (
	"sync"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"github.com/wlcn/yq-starter/config"
)

var (
	// Once is to init something
	Once sync.Once
	// AppConfig is App Config
	AppConfig config.ConfYaml
	// LogAccess is log server request log
	LogAccess *logrus.Logger
	// LogError is log server error log
	LogError *logrus.Logger
	// DB is to access database
	DB *gorm.DB
)
