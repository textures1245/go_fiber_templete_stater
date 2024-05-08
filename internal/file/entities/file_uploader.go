package entities

type FileUploaderReq struct {
	FileName string `json:"file_name" form:"file_nae" validate:"required" binding:"required"`
	FileData []byte `json:"file_data" form:"file_data" validate:"required" binding:"required"`
	FileType string `json:"file_type" form:"file_type" validate:"required" binding:"required"`
}
