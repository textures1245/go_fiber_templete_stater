package service

import (
	"payso/go_template/handler"
	"payso/go_template/model"
	"payso/go_template/repository"

	log "github.com/sirupsen/logrus"
)

type SampleService interface {
	SampleServiceFunction(someData model.SampleModel) (string, error)
}

type sampleService struct {
	someHandler handler.SampleHandler
}

func NewSampleService(someHandler handler.SampleHandler) sampleService {
	return sampleService{someHandler}
}

func (obj sampleService) SampleServiceFunction(someData model.SampleModel) (string, error) {
	/** Define log component **/
	log := log.WithFields(log.Fields{
		"component": "SampleService",
		"funciton":  "SampleServiceFunction",
	})

	log.Debugf("input : %v", someData)

	log.Debug("-= Sample Handler Call =-")
	handlerResult, err := obj.someHandler.SampleFunction("TEST")
	if err != nil {
		log.Errorf("Handler Error : %#v", err)
		return "ERROR", err
	}
	log.Debugf("handlerResult : %v", handlerResult)

	log.Debug("-= Sample Repository Call =-")
	repositoryResult, err := repository.AddSomeData(someData)

	if err != nil {
		log.Errorf("DB ERROR : %#v", err)
		return repositoryResult, err
	}

	return "OK", nil
}
