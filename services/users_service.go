package services

import (
	"github.com/Kungfucoding23/bookstore_users_api/domain/users"
	"github.com/Kungfucoding23/bookstore_users_api/utils/errors"
)

func GetUser(userID int64) (*users.User, *errors.RestErr) {
	result := &users.User{ID: userID}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

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

func UpdateUser(user users.User) (*users.User, *errors.RestErr) {
	current, err := GetUser(user.ID)
	if err != nil {
		return nil, err
	}
	current.FirstName = user.FirstName
	current.LastName = user.LastName
	current.Email = user.Email

	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}
