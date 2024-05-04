package apperror

import (
	"errors"
	"net/http"

	"github.com/go-sql-driver/mysql"
)

var (
	ErrUnauthorized     = errors.New("AccessingDatabaseError")
	ErrDatabaseNotFound = errors.New("DatabaseNotFound")
	// ErrDuplicateEntry         = errors.New("Duplicate entry")
	// ErrUnknownColumn          = errors.New("Unknown column")
	// ErrTableNotFound          = errors.New("Table not found")
	// ErrForeignKeyConstraint   = errors.New("Foreign key constraint fails")
	ErrMySQLConnection        = errors.New("CantConnectToMySQLServer")
	ErrMySQLConnectionTimeout = errors.New("CantConnectToMySQLServerOnHost")
	ErrMySQLLostConnection    = errors.New("LostConnectionToMySQLServerDuringQuery")
	ErrInternalServerError    = errors.New("InternalServerError")
)

func HandleDBError(err error) (int, error) {
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) {
		switch mysqlErr.Number {
		case 1045:
			return http.StatusUnauthorized, ErrUnauthorized
		case 1049:
			return http.StatusNotFound, ErrDatabaseNotFound
		case 2002:
			return http.StatusServiceUnavailable, ErrMySQLConnection
		case 2003:
			return http.StatusGatewayTimeout, ErrMySQLConnectionTimeout
		case 2013:
			return http.StatusGatewayTimeout, ErrMySQLLostConnection
		default:
			return http.StatusInternalServerError, err
		}
	}
	return http.StatusInternalServerError, err
}
