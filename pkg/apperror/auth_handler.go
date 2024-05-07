package apperror

import (
	"database/sql"
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
			return http.StatusUnauthorized, NewCErr(ErrorUnauthorized, err)
		case 1049:
			return http.StatusNotFound, NewCErr(ErrorDatabaseNotFound, err)
		case 1062:
			return http.StatusConflict, NewCErr(ErrorDuplicateEntry, err)
		case 1146:
			return http.StatusNotFound, NewCErr(ErrorTableNotFound, err)
		case 1217, 1452:
			return http.StatusConflict, NewCErr(ErrorForeignKeyConstraint, err)
		case 2002:
			return http.StatusServiceUnavailable, NewCErr(ErrorMySQLConnection, err)
		}
	}
	return http.StatusInternalServerError, NewCErr(errors.New("UnknownCustomErrorCanNotBeMade"), err)
}
