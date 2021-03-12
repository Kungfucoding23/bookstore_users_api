package services

import (
	"github.com/Kungfucoding23/bookstore_users_api/domain/users"
	"github.com/Kungfucoding23/bookstore_users_api/utils/errors"
)

//CreateUser service
func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	//attempt  to save the user in the database
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}
