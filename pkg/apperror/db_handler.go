package apperror

import (
	"errors"
	"net/http"

	"github.com/go-sql-driver/mysql"
)

var (
	errUnauthorized     = errors.New("AccessingDatabaseerror")
	errDatabaseNotFound = errors.New("DatabaseNotFound")
	// errDuplicateEntry         = errors.New("Duplicate entry")
	// errUnknownColumn          = errors.New("Unknown column")
	// errTableNotFound          = errors.New("Table not found")
	// errForeignKeyConstraint   = errors.New("Foreign key constraint fails")
	errMySQLConnection        = errors.New("CantConnectToMySQLServer")
	errMySQLConnectionTimeout = errors.New("CantConnectToMySQLServerOnHost")
	errMySQLLostConnection    = errors.New("LostConnectionToMySQLServerDuringQuery")
)

func HandleDBerror(err error) (int, error) {
	var mysqlerr *mysql.MySQLError
	if errors.As(err, &mysqlerr) {
		switch mysqlerr.Number {
		case 1045:
			return http.StatusUnauthorized, errUnauthorized
		case 1049:
			return http.StatusNotFound, errDatabaseNotFound
		case 2002:
			return http.StatusServiceUnavailable, errMySQLConnection
		case 2003:
			return http.StatusGatewayTimeout, errMySQLConnectionTimeout
		case 2013:
			return http.StatusGatewayTimeout, errMySQLLostConnection
		default:
			return http.StatusInternalServerError, err
		}
	}
	return http.StatusInternalServerError, err
}
