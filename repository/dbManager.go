package repository

import (
	"context"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"github.com/textures1245/go-template/util"

	"github.com/spf13/viper"
)

var (
	Db       *sqlx.DB
	server   string
	port     int
	user     string
	password string
	database string
)

func init() {
	util.Init()
	server = viper.GetString("DB_SERVER")
	port = viper.GetInt("DB_PORT")
	if port == 0 {
		log.Fatalf("Invalid or missing DB_PORT")
	}
	user = viper.GetString("DB_USER")
	password = viper.GetString("DB_PASS")
	database = viper.GetString("DB_INST")
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, server, port, database)
	log.Debug(connString)

	// Create connection pool
	Db, err := sqlx.Open("mysql", connString)

	// db, err := sql.Open("mysql", "userName:password@tcp(123.45.67.89:3030)/myDB")
	if err != nil {
		log.Error("**** Error creating connection pool: " + err.Error())
		panic(err.Error())
	}

	ctx := context.Background()
	err = Db.PingContext(ctx)
	if err != nil {
		fmt.Println("Catching ERR")
		log.Fatal(err.Error())
	}

	log.Debug("Connected!\n")

}

func GetDb() *sqlx.DB {
	var err error

	if Db == nil {
		// connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s", server, user, password, port, database)
		connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, server, port, database)
		log.Debug(connString)

		// Create connection pool
		Db, err = sqlx.Open("mysql", connString)
		if err != nil {
			log.Error("#### Error creating connection pool: " + err.Error())
		}

	} else {
		log.Debug("-= have current connection =-")
	}

	return Db
}
