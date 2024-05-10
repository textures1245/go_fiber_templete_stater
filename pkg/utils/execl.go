package utils

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"time"

	"github.com/textures1245/go-template/pkg/apperror"
	"github.com/xuri/excelize/v2"
)

type Excel[T interface{}] struct {
	Data []*T
}

func (e Excel[T]) ExportData() (*struct {
	FileName   string
	FileBuffer bytes.Buffer
}, *apperror.CErr) {
	genTimestamp := time.Now().Format("20060102150405")
	fileName := fmt.Sprintf("file_%s.xlsx", genTimestamp)
	f := excelize.NewFile()
	index, err := f.NewSheet(fileName)
	if err != nil {
		return nil, apperror.NewCErr(errors.New("Failed to create new sheet"), err)
	}

	for i, dat := range e.Data {
		// TODO: Refactor this to skip the pointer check if the value is a inside a struct not an pointer
		v := reflect.ValueOf(dat).Elem()

		// typeOfP := v.Type()

		// Loop over the fields of the product
		for j := 0; j < v.NumField(); j++ {
			// Get the interface{} value
			val := v.Field(j).Interface()

			// Check if the value is a pointer
			rv := reflect.ValueOf(val)
			if rv.Kind() == reflect.Ptr || rv.Kind() == reflect.Slice || rv.Kind() == reflect.Map || rv.Kind() == reflect.Chan || rv.Kind() == reflect.Func || rv.Kind() == reflect.Interface {
				// If the value is a pointer and not nil, dereference it
				if !rv.IsNil() {
					rv = rv.Elem()
					val = rv.Interface()
				} else {
					// If the value is a nil pointer, set val to nil
					val = nil
				}
			}

			// Set the value of the cell in the Excel file
			cellName := fmt.Sprintf("%c%d", 'A'+j, i+1)
			f.SetCellValue(fileName, cellName, val)
		}
	}
	f.SetActiveSheet(index)

	var buf bytes.Buffer
	if errOnW := f.Write(&buf); errOnW != nil {
		return nil, apperror.NewCErr(errors.New("Failed to export Excel file"), errOnW)

	}

	return &struct {
		FileName   string
		FileBuffer bytes.Buffer
	}{
		FileName:   fileName,
		FileBuffer: buf,
	}, nil
}
