package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//CreateUser creates a user
func CreateUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}

//GetUser get a user
func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}
