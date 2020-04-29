package reports

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func sendError(c *gin.Context, err error, format Format, status int) {
	//debug.PrintStack()
	log.Error(err)
	switch format {
	case Json:
		c.JSON(status, gin.H{"error": err.Error()})
	case Html:
		c.HTML(status, "error.gohtml", gin.H{"error": err.Error()})
	default:
		c.JSON(status, gin.H{"error": err.Error()})
	}
}
