package mysql_utils

import (
	"strings"

	"github.com/Kungfucoding23/bookstore_users_api/utils/errors"
	"github.com/go-sql-driver/mysql"
)

const (
	errorNoRows      = "no rows in result set"
	indexUniqueEmail = "email_UNIQUE"
)

func ParseError(err error) *errors.RestErr {
	sqlError, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError("no record matching given id")
		}
		return errors.NewInternalServerError("error parsing database response")
	}
	switch sqlError.Number {
	case 1062:
		return errors.NewBadRequestError("invalid data")
	}
	return errors.NewInternalServerError("error processing request")
}
