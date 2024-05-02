package repository_query

import _ "embed"

var (
	//go:embed user/get_users.sql
	InsertUsers string
)
