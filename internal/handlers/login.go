package handlers

import (
	"myapp/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)


func Login(c *gin.Context) {
	mutex.Lock()
	defer mutex.Unlock()
	user, password, _ := c.Request.BasicAuth()
	if storage.ContainsUser(user, password) {
		c.Redirect(http.StatusAccepted, "/")
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid username or password"})
	}
}
