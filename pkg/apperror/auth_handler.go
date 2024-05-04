package apperror

import (
	"errors"
	"net/http"

	"github.com/go-sql-driver/mysql"
)

var (
	errorUnauthorized         = errors.New("dbUnauthorized")
	errorDatabaseNotFound     = errors.New("databaseNotFound")
	errorDuplicateEntry       = errors.New("username already exists")
	errorTableNotFound        = errors.New("user table not found")
	errorForeignKeyConstraint = errors.New("cant create or update User: a foreign key constraint fails")
	errorMySQLConnection      = errors.New("cantConnectToMySQLServer")
)

func HandleAuthError(err error) (int, error) {
	var mysqlErr *mysql.MySQLError

	if errors.As(err, &mysqlErr) {
		switch mysqlErr.Number {
		case 1045:
			return http.StatusUnauthorized, errorUnauthorized
		case 1049:
			return http.StatusNotFound, errorDatabaseNotFound
		case 1062:
			return http.StatusConflict, errorDuplicateEntry
		case 1146:
			return http.StatusNotFound, errorTableNotFound
		case 1217, 1452:
			return http.StatusConflict, errorForeignKeyConstraint
		case 2002:
			return http.StatusServiceUnavailable, errorMySQLConnection
		}
	}
	return http.StatusInternalServerError, err
}
