package utils

import (
	"net/http"

	"errors"

	"github.com/go-playground/validator"
	"github.com/textures1245/go-template/pkg/apperror"
)

func SchemaValidator[T any](req *T) *apperror.CErr {
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		err_validator := err.(validator.ValidationErrors)
		return apperror.NewCErr(errors.New(http.StatusText(http.StatusBadRequest)), err_validator)
	}

	return nil
}
