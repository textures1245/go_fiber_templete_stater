package apperror

import (
	"errors"
	"net/http"

	"github.com/go-sql-driver/mysql"
)

var (
	ErrorUnauthorized         = errors.New("dbUnauthorized")
	ErrorDatabaseNotFound     = errors.New("databaseNotFound")
	ErrorDuplicateEntry       = errors.New("username already exists")
	ErrorTableNotFound        = errors.New("user table not found")
	ErrorForeignKeyConstraint = errors.New("cant create or update User: a foreign key constraint fails")
	ErrorMySQLConnection      = errors.New("cantConnectToMySQLServer")
)

func HandleAuthError(err error) (int, error) {
	var mysqlErr *mysql.MySQLError

	if errors.As(err, &mysqlErr) {
		switch mysqlErr.Number {
		case 1045:
			return http.StatusUnauthorized, ErrorUnauthorized
		case 1049:
			return http.StatusNotFound, ErrorDatabaseNotFound
		case 1062:
			return http.StatusConflict, ErrorDuplicateEntry
		case 1146:
			return http.StatusNotFound, ErrorTableNotFound
		case 1217, 1452:
			return http.StatusConflict, ErrorForeignKeyConstraint
		case 2002:
			return http.StatusServiceUnavailable, ErrorMySQLConnection
		}
	}
	return http.StatusInternalServerError, err
}
