package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/wlcn/yq-starter/config"
	"github.com/wlcn/yq-starter/helper"
	"github.com/wlcn/yq-starter/service/article"
	"github.com/wlcn/yq-starter/service/auth"
	"github.com/wlcn/yq-starter/service/image"
	"github.com/wlcn/yq-starter/service/user"
)

func routerEngine() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	// Global middleware
	r.Use(helper.LogMiddleware())
	r.Use(helper.VersionMiddleware())

	auth.Routers(r.Group("/auth"))
	// group: v1
	v1 := r.Group("/api/v1")
	v1.Use(helper.JWTMiddleware())
	user.Routers(v1.Group("/user"))
	article.Routers(v1.Group("/article"))
	image.Routers(v1.Group("/image"))
	return r
}

func main() {
	opts := config.ConfYaml{}
	var (
		showVersion bool
		configFile  string
	)
	flag.StringVar(&opts.Core.Address, "A", "", "address to bind")
	flag.StringVar(&opts.Core.Address, "address", "", "address to bind")
	flag.StringVar(&opts.Core.Port, "p", "", "port number for YQ")
	flag.StringVar(&opts.Core.Port, "port", "", "port number for YQ")
	flag.BoolVar(&showVersion, "version", false, "Print version information.")
	flag.BoolVar(&showVersion, "v", false, "Print version information.")
	flag.StringVar(&configFile, "c", "", "Configuration file path.")
	flag.StringVar(&configFile, "config", "", "Configuration file path.")
	flag.Usage = usage
	flag.Parse()

	var err error
	// set default parameters.
	helper.AppConfig, err = config.LoadConf(configFile)
	if err != nil {
		helper.LogAccess.Errorf("Load yaml config file error: '%v'", err)
		return
	}
	// 设置版本号
	helper.SetVersion(helper.AppConfig.Core.Version)

	// Show version and exit
	if showVersion {
		helper.PrintVersion()
		os.Exit(0)
	}
	// overwrite server port and address
	if opts.Core.Port != "" {
		helper.AppConfig.Core.Port = opts.Core.Port
	}
	if opts.Core.Address != "" {
		helper.AppConfig.Core.Address = opts.Core.Address
	}

	if err = initDir(); err != nil {
		log.Fatalf("Can't init directory for application, error: %v", err)
	}

	if err = helper.InitLog(); err != nil {
		log.Fatalf("Can't load log module, error: %v", err)
	}

	if err = helper.InitAppDatabase(); err != nil {
		log.Fatalf("Can't load Database module, error: %v", err)
	}

	defer helper.DB.Close()

	Migrate()
	// Logging to a file.
	f, _ := os.Create(helper.AppConfig.Log.GinLog)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	// set server mode
	gin.SetMode(helper.AppConfig.Core.Mode)

	r := routerEngine()

	r.Run(strings.Join([]string{helper.AppConfig.Core.Address, helper.AppConfig.Core.Port}, ":"))
}

var usageStr = `
Usage: yq [options]
Server Options:
    -A, --address <address>          Address to bind (default: any)
    -p, --port <port>                Use port for clients (default: 8080)
    -c, --config <file>              Configuration file path
Common Options:
    -h, --help                       Show this message
    -v, --version                    Show version
`

// usage will print out the flag options for the server.
func usage() {
	fmt.Printf("%s\n", usageStr)
	os.Exit(0)
}

func initDir() (err error) {
	if helper.AppConfig.Log.GinLog != "" {
		err = os.MkdirAll(filepath.Dir(helper.AppConfig.Log.GinLog), os.ModePerm)
		if err != nil {
			return
		}
		err = os.MkdirAll(filepath.Dir(helper.AppConfig.Log.AccessLevel), os.ModePerm)
		if err != nil {
			return
		}
		err = os.MkdirAll(filepath.Dir(helper.AppConfig.Log.ErrorLevel), os.ModePerm)
		if err != nil {
			return
		}
		err = os.MkdirAll(filepath.Dir(helper.AppConfig.Database.Sqlite3.URL), os.ModePerm)
		if err != nil {
			return
		}
	}
	return
}

// Migrate schame
func Migrate() {
	helper.DB.AutoMigrate(&user.User{})
	helper.DB.AutoMigrate(&article.Article{})
	helper.DB.AutoMigrate(&image.Image{})
}
