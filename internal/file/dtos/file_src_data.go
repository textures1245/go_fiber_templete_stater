package dtos

type FileSourceDataRes struct {
	FileName      string `json:"file_name"`
	FileBase64URL string `json:"file_src_base64_url"`
	FilePathData  string `json:"file_path_data"`
	FileType      string `json:"file_type"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}
