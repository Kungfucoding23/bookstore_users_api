package app

import (
	"github.com/Kungfucoding23/bookstore_users_api/logger"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

//StartApp starts the application
func StartApp() {
	mapUrls()
	logger.Info("about to start the application")
	router.Run()
}
