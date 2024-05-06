package apperror

import (
	"database/sql"
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

	// custom Error
	ErrorInvalidCredentials = errors.New("Invalid credentials")
)

func HandleAuthError(err error) (int, *CErr) {
	var mysqlErr *mysql.MySQLError

	if errors.Is(err, sql.ErrNoRows) {
		return http.StatusNotFound, NewCErr(errors.New("username not found"), err)
	}

	if errors.As(err, &mysqlErr) {
		switch mysqlErr.Number {
		case 1045:
			return http.StatusUnauthorized, NewCErr(errorUnauthorized, err)
		case 1049:
			return http.StatusNotFound, NewCErr(errorDatabaseNotFound, err)
		case 1062:
			return http.StatusConflict, NewCErr(errorDuplicateEntry, err)
		case 1146:
			return http.StatusNotFound, NewCErr(errorTableNotFound, err)
		case 1217, 1452:
			return http.StatusConflict, NewCErr(errorForeignKeyConstraint, err)
		case 2002:
			return http.StatusServiceUnavailable, NewCErr(errorMySQLConnection, err)
		}
	}
	return http.StatusInternalServerError, NewCErr(errors.New("UnknownCustomErrorCanNotBeMade"), err)
}
