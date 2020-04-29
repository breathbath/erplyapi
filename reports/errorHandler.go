package reports

import (
	"github.com/gin-gonic/gin"
)

func sendError(c *gin.Context, err error, format Format, status int) {
	switch format {
	case Json:
		c.JSON(status, gin.H{"error": err.Error()})
	case Html:
		c.HTML(status, "error.gohtml", gin.H{"error": err.Error()})
	default:
		c.JSON(status, gin.H{"error": err.Error()})
	}
}
