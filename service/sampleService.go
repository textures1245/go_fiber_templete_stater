package service

import (
	"runtime"
	"strings"

	"github.com/textures1245/go-template/handler"
	"github.com/textures1245/go-template/model"
	"github.com/textures1245/go-template/repository"

	log "github.com/sirupsen/logrus"
)

type SampleService interface {
	SampleServiceFunction(someData model.SampleModel) (string, error)
}

type sampleService struct {
	someHandler handler.SimpleHandler
}

func NewSampleService(someHandler handler.SimpleHandler) sampleService {
	return sampleService{someHandler}
}

func (obj sampleService) SampleServiceFunction(someData model.SampleModel) (string, error) {
	/** Define log component **/
	_, file, _, _ := runtime.Caller(0)
	pc, _, _, _ := runtime.Caller(0)
	functionName := strings.Split(runtime.FuncForPC(pc).Name(), ".")[len(strings.Split(runtime.FuncForPC(pc).Name(), "."))-1]

	log := log.WithFields(log.Fields{
		"component": strings.Split(file, "/")[len(strings.Split(file, "/"))-1],
		"funciton":  functionName,
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
	repositoryResult, err := repository.GetSomeData("0")

	if err != nil {
		log.Errorf("DB ERROR : %#v", err)
		return "ERROR", err
	}
	log.Debugf("RESULT :%v", repositoryResult)

	return "OK", nil
}
