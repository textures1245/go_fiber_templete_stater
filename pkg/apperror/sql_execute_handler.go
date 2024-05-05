package apperror

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/go-sql-driver/mysql"
)

func CustomSqlExecuteHandler(modelName string, err error) (int, *CErr) {
	if errors.Is(err, sql.ErrNoRows) {
		return http.StatusNotFound, NewCErr(errors.New(modelName+" not found"), err)
	}

	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) {
		switch mysqlErr.Number {
		case 1045:
			return http.StatusUnauthorized, NewCErr(errors.New("InvalidDbCredentialsWhileExecute"), err)
		case 1049:
			return http.StatusNotFound, NewCErr(errors.New("DatabaseNotSelected"), err)
		case 1062:
			return http.StatusConflict, NewCErr(errors.New("Duplicate entry in "+modelName), err)
		case 1054:
			return http.StatusBadRequest, NewCErr(errors.New("Unknown column for "+modelName), err)
		case 1146:
			return http.StatusNotFound, NewCErr(errors.New("Table "+modelName+" not found"), err)
		case 1217, 1452:
			return http.StatusConflict, NewCErr(errors.New("Foreign key constraint fails in "+modelName), err)
		case 2013:
			return http.StatusGatewayTimeout, NewCErr(errors.New("LostConnectionToMySQLserverDuringQuery"), err)
		default:
			return http.StatusInternalServerError, NewCErr(errors.New("UnknownCustomCanNotBeMade"), err)
		}
	}
	return http.StatusInternalServerError, NewCErr(errors.New("UnknownCustomCanNotBeMade"), err)
}
