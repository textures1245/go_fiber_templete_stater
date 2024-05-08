package dtos

type FileSourceDataRes struct {
	FileName   string `json:"file_name"`
	FileSrcURL string `json:"file_src_url"`
	FileType   string `json:"file_type"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}
