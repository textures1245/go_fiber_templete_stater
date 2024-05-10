package dtos

import "bytes"

type ProductToExcelRes struct {
	FileName   string `json:"file_name"`
	FileBuffer bytes.Buffer
}
