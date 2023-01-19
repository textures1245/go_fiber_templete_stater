package handler

import (
	"payso/go_template/model"

	log "github.com/sirupsen/logrus"
)

type SampleHandler interface {
	SampleFunction(data string) (model.SampleModel, error)
}

type sampleHandler struct {
}

func NewSampleHandler() sampleHandler {
	return sampleHandler{}
}

func (obj sampleHandler) SampleFunction(data string) (model.SampleModel, error) {
	/** Define log component **/
	log := log.WithFields(log.Fields{
		"component": "SampleHandler",
		"funciton":  "SampleFunction",
	})
	log.Info("-= SampleHandler:SampleFunction =-")
	/** Write Some Code that call other Services **/
	return model.SampleModel{}, nil
}
