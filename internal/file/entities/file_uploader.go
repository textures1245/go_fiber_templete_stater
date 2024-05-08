package entities

type FileUploaderReq struct {
	FileName string `json:"file_name" form:"file_nae" validate:"required" binding:"required"`
	FileData string `json:"file_data" form:"file_data" validate:"base64" binding:"required"`
	FileType string `json:"file_type" form:"file_type" validate:"required" binding:"required"`
}
