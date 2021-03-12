package services

import (
	"github.com/Kungfucoding23/bookstore_users_api/domain/users"
	"github.com/Kungfucoding23/bookstore_users_api/utils/errors"
)

//CreateUser service
func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	return &user, nil
}
