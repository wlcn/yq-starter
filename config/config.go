// Package config defines the config structs and some config parser interfaces and implementations
package config

import (
	"bytes"
	"io/ioutil"
	"strings"

	"github.com/spf13/viper"
)

var defaultConf = []byte(`
core:
  mode: "debug"
  version: "yq-starter-1.0"
  address: "" # ip address to bind (default: any)
  port: "8080" # ignore this port number if auto_tls is enabled (listen 443).  
log:
  gin_log: "/tmp/yq-starter/logs/gin.log" # define log path like "logs/gin.log"
  format: "string" # string or json
  access_log: "/tmp/yq-starter/logs/access.log" # stdout: output to console, or define log path like "logs/access.log"
  access_level: "debug"
  error_log: "/tmp/yq-starter/logs/error.log" # stderr: output to console, or define log path like "logs/error.log"
  error_level: "error"
database:
  engine: "sqlite3" # support mysql, sqlite3
  log_mode: true # 是否打印sql语句
  mysql: 
    url: "user:password@/dbname?charset=utf8&parseTime=True&loc=Local"
  sqlite3: 
    url: "/tmp/yq-starter/app_test.db"
`)

// ConfYaml is config structure.
type ConfYaml struct {
	Core     SectionCore     `yaml:"core"`
	Log      SectionLog      `yaml:"log"`
	Database SectionDatabase `yaml:"database"`
}

// SectionCore is sub section of config.
type SectionCore struct {
	Mode    string `yaml:"mode"`
	Version string `yaml:"version"`
	Enabled bool   `yaml:"enabled"`
	Address string `yaml:"address"`
	Port    string `yaml:"port"`
}

// SectionLog is sub section of config.
type SectionLog struct {
	GinLog      string `yaml:"gin_log"`
	Format      string `yaml:"format"`
	AccessLog   string `yaml:"access_log"`
	AccessLevel string `yaml:"access_level"`
	ErrorLog    string `yaml:"error_log"`
	ErrorLevel  string `yaml:"error_level"`
}

// SectionDatabase is sub section of config.
type SectionDatabase struct {
	Engine  string         `yaml:"engine"`
	LogMode bool           `yaml:"log_mode"`
	Mysql   SectionMysql   `yaml:"mysql"`
	Sqlite3 SectionSqlite3 `yaml:"sqlite3"`
}

// SectionMysql is sub section of config.
type SectionMysql struct {
	URL string `yaml:"url"`
}

// SectionSqlite3 is sub section of config.
type SectionSqlite3 struct {
	URL string `yaml:"url"`
}

// LoadConf load config from file and read in environment variables that match
func LoadConf(confPath string) (ConfYaml, error) {
	var conf ConfYaml
	viper.SetConfigType("yaml")
	viper.SetEnvPrefix("yq") // will be uppercased automatically
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetConfigName("config")   // name of config file (without extension)
	viper.AddConfigPath("/etc/yq/") // path to look for the config file in
	viper.AddConfigPath("$HOME/yq") // call multiple times to add many search paths
	viper.AddConfigPath(".")        // optionally look for config in the working directory
	viper.AddConfigPath("./config") // optionally look for config in the ./config directory

	if confPath != "" {
		content, err := ioutil.ReadFile(confPath)
		if err != nil {
			return conf, err
		}
		if err := viper.ReadConfig(bytes.NewBuffer(content)); err != nil {
			return conf, err
		}
		// fmt.Println("load from confPath:", confPath)
	} else {
		if err := viper.ReadInConfig(); err == nil {
			// fmt.Println("Using config file:", viper.ConfigFileUsed())
		} else {
			// fmt.Println("load default config:")
			// load default config
			if err := viper.ReadConfig(bytes.NewBuffer(defaultConf)); err != nil {
				return conf, err
			}
		}
	}

	// Core
	conf.Core.Mode = viper.GetString("core.mode")
	conf.Core.Address = viper.GetString("core.address")
	conf.Core.Port = viper.GetString("core.port")
	conf.Core.Version = viper.GetString("core.version")
	conf.Core.Enabled = viper.GetBool("core.enabled")

	// log
	conf.Log.GinLog = viper.GetString("log.gin_log")
	conf.Log.Format = viper.GetString("log.format")
	conf.Log.AccessLog = viper.GetString("log.access_log")
	conf.Log.AccessLevel = viper.GetString("log.access_level")
	conf.Log.ErrorLog = viper.GetString("log.error_log")
	conf.Log.ErrorLevel = viper.GetString("log.error_level")

	// Database
	conf.Database.Engine = viper.GetString("database.engine")
	conf.Database.LogMode = viper.GetBool("database.log_mode")
	conf.Database.Mysql.URL = viper.GetString("database.mysql.url")
	conf.Database.Sqlite3.URL = viper.GetString("database.sqlite3.url")
	return conf, nil
}
