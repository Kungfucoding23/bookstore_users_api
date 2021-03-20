package mysql_utils

import (
	"errors"
	"fmt"
	"strings"

	"github.com/Kungfucoding23/bookstore_utils-go/rest_errors"
	"github.com/go-sql-driver/mysql"
)

const (
	ErrorNoRows      = "no rows in result set"
	indexUniqueEmail = "email_UNIQUE"
)

func ParseError(err error) *rest_errors.RestErr {
	sqlError, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), ErrorNoRows) {
			return rest_errors.NewNotFoundError("no record matching given id")
		}
		return rest_errors.NewInternalServerError("error parsing database response", err)
	}
	switch sqlError.Number {
	case 1062:
		return rest_errors.NewBadRequestError(fmt.Sprintf("invalid data: " + err.Error()))
	}
	return rest_errors.NewInternalServerError("error processing request", errors.New("database error"))
}
