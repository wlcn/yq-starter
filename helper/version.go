package helper

import (
	"fmt"
	"runtime"

	"github.com/gin-gonic/gin"
)

var version string

// SetVersion for setup version string.
func SetVersion(ver string) {
	version = ver
}

// GetVersion for get current version.
func GetVersion() string {
	return version
}

// PrintVersion provide print server engine
func PrintVersion() {
	fmt.Printf(`yq-starter Version is %s, Compiler: %s %s, Copyright (C) 2019 Long Wang, Inc.`,
		version,
		runtime.Compiler,
		runtime.Version())
	fmt.Println()
}

// VersionMiddleware : add version on header.
func VersionMiddleware() gin.HandlerFunc {
	// Set out header value for each response
	return func(c *gin.Context) {
		c.Header("X-App-Version", version)
		c.Next()
	}
}
