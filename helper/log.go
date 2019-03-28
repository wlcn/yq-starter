package helper

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mattn/go-isatty"
	"github.com/sirupsen/logrus"
)

var (
	magenta = string([]byte{27, 91, 57, 55, 59, 52, 53, 109})
	reset   = string([]byte{27, 91, 48, 109})
)

// LogReq is http request log
type LogReq struct {
	URI         string `json:"uri"`
	Method      string `json:"method"`
	IP          string `json:"ip"`
	ContentType string `json:"content_type"`
	Agent       string `json:"agent"`
}

var isTerm bool

func init() {
	isTerm = isatty.IsTerminal(os.Stdout.Fd())
}

// InitLog use for initial log module
func InitLog() error {

	var err error

	// init logger
	LogAccess = logrus.New()
	LogError = logrus.New()

	LogAccess.Formatter = &logrus.TextFormatter{
		TimestampFormat: "2019/11/11 - 11:11:11",
		FullTimestamp:   true,
	}

	LogError.Formatter = &logrus.TextFormatter{
		TimestampFormat: "2019/11/11 - 11:11:11",
		FullTimestamp:   true,
	}

	// set logger
	if err = SetLogLevel(LogAccess, AppConfig.Log.AccessLevel); err != nil {
		return errors.New("Set access log level error: " + err.Error())
	}
	if err = SetLogLevel(LogError, AppConfig.Log.ErrorLevel); err != nil {
		return errors.New("Set error log level error: " + err.Error())
	}
	if err = SetLogOut(LogAccess, AppConfig.Log.AccessLog); err != nil {
		return errors.New("Set access log path error: " + err.Error())
	}
	if err = SetLogOut(LogError, AppConfig.Log.ErrorLog); err != nil {
		return errors.New("Set error log path error: " + err.Error())
	}
	return nil
}

// SetLogOut provide log stdout and stderr output
func SetLogOut(log *logrus.Logger, outString string) error {
	switch outString {
	case "stdout":
		log.Out = os.Stdout
	case "stderr":
		log.Out = os.Stderr
	default:
		f, err := os.OpenFile(outString, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
		if err != nil {
			return err
		}
		log.Out = f
	}
	return nil
}

// SetLogLevel is define log level what you want
// log level: panic, fatal, error, warn, info and debug
func SetLogLevel(log *logrus.Logger, levelString string) error {
	level, err := logrus.ParseLevel(levelString)
	if err != nil {
		return err
	}
	log.Level = level
	return nil
}

// LogRequest record http request
func LogRequest(uri string, method string, ip string, contentType string, agent string) {
	var output string
	log := &LogReq{
		URI:         uri,
		Method:      method,
		IP:          ip,
		ContentType: contentType,
		Agent:       agent,
	}

	if AppConfig.Log.Format == "json" {
		logJSON, _ := json.Marshal(log)
		output = string(logJSON)
	} else {
		var headerColor, resetColor string
		if isTerm {
			headerColor = magenta
			resetColor = reset
		}
		// format is string
		output = fmt.Sprintf("|%s header %s| %s %s %s %s %s",
			headerColor, resetColor,
			log.Method,
			log.URI,
			log.IP,
			log.ContentType,
			log.Agent,
		)
	}
	LogAccess.Info(output)
}

// LogMiddleware provide gin router handler.
func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		LogRequest(c.Request.URL.Path, c.Request.Method, c.ClientIP(), c.ContentType(), c.GetHeader("User-Agent"))
		c.Next()
	}
}
