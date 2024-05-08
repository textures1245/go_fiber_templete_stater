package repository_query

import _ "embed"

var (
	//go:embed file/create_file.sql
	CreateFile string

	//go:embed file/get_files.sql
	GetFiles string
)
