package handlers

import "github.com/gin-gonic/gin"

func Health(c *gin.Context) {
	c.JSON(200, gin.H{"status": "ok"})
}

func Root(c *gin.Context) {
	c.JSON(200, gin.H{"message": "todo api is running"})
}
