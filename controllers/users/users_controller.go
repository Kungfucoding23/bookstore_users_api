package users

import (
	"net/http"
	"strconv"

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
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

//GetUser get a user
func GetUser(c *gin.Context) {
	userID, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("invalid user id")
		c.JSON(err.Status, err)
		return
	}
	user, getErr := services.GetUser(userID)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}
