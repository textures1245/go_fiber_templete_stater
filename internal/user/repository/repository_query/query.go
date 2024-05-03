package repository_query

import _ "embed"

var (
	//go:embed user/get_users.sql
	GetUsers string

	//go:embed user/find_user_by_id.sql
	FindUserById string

	//go:embed user/find_user_by_user_credential.sql
	FindUserByUsernameAndPassword string

	//go:embed user/insert.sql
	InsertUser string

	//go:embed user/find_user_by_email.sql
	FindUserByEmail string
)
