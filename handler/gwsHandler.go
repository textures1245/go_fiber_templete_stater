package handler

import (
	"payso/payment-service/model"

	log "github.com/sirupsen/logrus"
)

type GWSHandler interface {
	SampleFunction(data string) (model.SampleModel, error)
	GetGateWay(input map[string]string) (string, error)
}

type gwsHandler struct {
}

func NewGWSHandler() gwsHandler {
	return gwsHandler{}
}

func (obj gwsHandler) SampleFunction(data string) (model.SampleModel, error) {
	/** Define log component **/
	log := log.WithFields(log.Fields{
		"component": "SampleHandler",
		"funciton":  "SampleFunction",
	})
	log.Info("-= SampleHandler:SampleFunction =-")
	/** Write Some Code that call other Services **/
	return model.SampleModel{}, nil
}

func (obj gwsHandler) GetGateWay(input map[string]string) (string, error) {
	log := log.WithFields(log.Fields{
		"component": "GWSHandler",
		"funciton":  "GetGateWay",
	})
	log.Infof("input : %#v", input)
	return "KBANK", nil
}
