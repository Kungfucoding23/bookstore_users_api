package mysql_utils

import (
	"fmt"
	"strings"

	"github.com/Kungfucoding23/bookstore_users_api/utils/errors"
	"github.com/go-sql-driver/mysql"
)

const (
	ErrorNoRows      = "no rows in result set"
	indexUniqueEmail = "email_UNIQUE"
)

func ParseError(err error) *errors.RestErr {
	sqlError, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), ErrorNoRows) {
			return errors.NewNotFoundError("no record matching given id")
		}
		return errors.NewInternalServerError("error parsing database response")
	}
	switch sqlError.Number {
	case 1062:
		return errors.NewBadRequestError(fmt.Sprintf("invalid data: " + err.Error()))
	}
	return errors.NewInternalServerError("error processing request")
}
