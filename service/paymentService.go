package service

import (
	"payso/payment-service/handler"
	"payso/payment-service/model"
	"payso/payment-service/repository"

	log "github.com/sirupsen/logrus"
)

type PaymentService interface {
	SampleServiceFunction(someData model.SampleModel) (string, error)
}

type paymentService struct {
	someHandler handler.GWSHandler
}

func NewPaymentService(someHandler handler.GWSHandler) paymentService {
	return paymentService{someHandler}
}

func (obj paymentService) SampleServiceFunction(someData model.SampleModel) (string, error) {
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
