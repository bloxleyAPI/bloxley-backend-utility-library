package utils

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func ErrorResponse(err error) gin.H {
	if err == nil {
		return gin.H{"error": nil}
	}
	// Parse the error message and create a map of errors
	errorMap := make(map[string]string)
	for _, line := range strings.Split(err.Error(), "\n") {
		parts := strings.SplitN(line, " Error:", 2)
		if len(parts) == 2 {
			field := strings.TrimSpace(parts[0])
			message := strings.TrimSpace(parts[1])
			if field != "" && message != "" {
				errorMap[field] = message
			}
		}
	}

	if len(errorMap) == 0 {
		return gin.H{"error": err.Error()}
	}

	return gin.H{"errors": errorMap}
}
