package utils

import (
	"encoding/base64"
	"net/http"

	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/textures1245/go-template/pkg/apperror"
)

func SchemaValidator[T any](req *T) *apperror.CErr {
	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(req)

	if err != nil {
		err_validator := err.(validator.ValidationErrors)
		return apperror.NewCErr(errors.New(http.StatusText(http.StatusBadRequest)), err_validator)
	}

	return nil
}

func Base64Validate(dat []byte) error {

	_, err := base64.StdEncoding.DecodeString(string(dat))
	if err != nil {
		return errors.New("FileData must be a valid base64 string")
	}

	return nil
}
