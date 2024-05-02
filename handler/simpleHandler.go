package handler

import (
	"runtime"
	"strings"

	"github.com/textures1245/go-template/model"

	log "github.com/sirupsen/logrus"
)

type SimpleHandler interface {
	SampleFunction(data string) (model.SampleModel, error)
}

type simpleHandler struct {
}

func NewSimpleHandler() simpleHandler {
	return simpleHandler{}
}

func (obj simpleHandler) SampleFunction(data string) (model.SampleModel, error) {
	/** Define log component **/
	_, file, _, _ := runtime.Caller(0)
	pc, _, _, _ := runtime.Caller(0)
	functionName := strings.Split(runtime.FuncForPC(pc).Name(), ".")[len(strings.Split(runtime.FuncForPC(pc).Name(), "."))-1]

	log := log.WithFields(log.Fields{
		"component": strings.Split(file, "/")[len(strings.Split(file, "/"))-1],
		"funciton":  functionName,
	})

	log.Infof("data : %v", data)
	/** Write Some Code that call other Services **/
	return model.SampleModel{}, nil
}
