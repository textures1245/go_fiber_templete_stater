package repository

import (
	"context"
	"database/sql"
	"payso/payment-service/model"

	"github.com/blockloop/scan"
	log "github.com/sirupsen/logrus"
)

func GetSomeData(id string) (model.SampleModel, error) {
	/** Define log component **/
	log := log.WithFields(log.Fields{
		"component": "SampleRepository",
		"funciton":  "GetSomeData",
	})
	log.Debugf("input ==>%s ", id)

	db := GetDb()
	ctx := context.Background()

	// Check if database is alive.
	err := db.PingContext(ctx)

	if err != nil {
		log.Errorf("#%v", err)
		return model.SampleModel{}, err
	}

	// Execute query
	rows, err := db.QueryContext(ctx, model.SQL_simple_get_date, sql.Named("ID", id))

	if err != nil {
		log.Errorf(" %#v", err)
		return model.SampleModel{}, err
	}

	defer rows.Close()

	var someData model.SampleModel

	err = scan.Row(&someData, rows)

	defer rows.Close()

	if err != nil {
		log.Errorf(" %#v", err)
		return model.SampleModel{}, err
	}
	log.Infof("data %#v", someData)

	return someData, nil
}

func AddSomeData(someData model.SampleModel) (string, error) {
	/** Define log component **/
	log := log.WithFields(log.Fields{
		"component": "SampleRepository",
		"funciton":  "AddSomeData",
	})
	log.Debugf("input : %v", someData)
	db := GetDb()
	ctx := context.Background()

	// Check if database is alive.
	err := db.PingContext(ctx)

	if err != nil {
		log.Errorf("#%v", err)
		return "NONE", err
	}

	stmt, err := db.Prepare(model.SQL_simple_add)

	if err != nil {
		log.Errorf(" %#v", err)
		return "ERROR", err
	}

	_, err = stmt.Exec(someData.Column1,
		someData.Column2)

	if err != nil {
		log.Errorf(" %#v", err)
		return "ERROR", err
	}

	return "COMPLETE", nil
}
