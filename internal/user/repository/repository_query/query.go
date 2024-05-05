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

	//go:embed user/find_user_by_username.sql
	FindUserByUsername string

	//go:embed user/update_user_by_id.sql
	UpdateUserById string

	//go:embed user/delete_user_by_id.sql
	DeleteUserById string
)
