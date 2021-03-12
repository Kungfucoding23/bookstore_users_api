package users

import (
	"net/http"

	"github.com/Kungfucoding23/bookstore_users_api/domain/users"
	"github.com/Kungfucoding23/bookstore_users_api/services"
	"github.com/Kungfucoding23/bookstore_users_api/utils/errors"
	"github.com/gin-gonic/gin"
)

//CreateUser creates a user
func CreateUser(c *gin.Context) {
	var user users.User
	/* bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		//handle error
		return
	}
	if err := json.Unmarshal(bytes, &user); err != nil {
		//handle json error
		return
	}*/

	//this lane replaces the ioutil and unmarshal json func
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		//handle user creation error
		return
	}
	c.JSON(http.StatusCreated, result)
}

//GetUser get a user
func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}
