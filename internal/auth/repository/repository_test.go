package repository_test

import (
	"context"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	authEntities "github.com/textures1245/go-template/internal/auth/entities"
	"github.com/textures1245/go-template/internal/auth/repository"
	userEntities "github.com/textures1245/go-template/internal/user/entities"
	userRepo "github.com/textures1245/go-template/internal/user/repository"
	"github.com/textures1245/go-template/internal/user/repository/repository_query"
	"go.uber.org/mock/gomock"
)

// TODO: implement auth repository unit test

func Test_repo_SignUsersAccessToken(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	defer db.Close()

	var (
		conn     = sqlx.NewDb(db, "sqlmock")
		authRepo = repository.NewAuthRepository(conn)
		userRepo = userRepo.NewUserRepository(conn)
	)

	username := "username_test"
	userId := int64(1)

	t.Run("SignUsersAccessToken_positive_case", func(t *testing.T) {
		testDat := &userEntities.UserCreatedReq{
			Name:        "test",
			Username:    username,
			Password:    "pw_test",
			Email:       "test@gmail.com",
			PhoneNumber: "08123456789",
			IdCard:      "1234567890",
		}

		mock.
			ExpectExec(repository_query.InsertUser).
			WithArgs(testDat.Name, testDat.Username, testDat.Password, testDat.Email, testDat.PhoneNumber, testDat.IdCard).
			WillReturnResult(sqlmock.NewResult(userId, 1))

		userID, err := userRepo.CreateUser(context.Background(), testDat)
		assert.NoError(t, err)
		assert.NotEmpty(t, userID)

		// mock.ExpectQuery(repository_query.FindUserByUsername).
		// 	WithArgs(username).
		// 	WillReturnResult(sqlmock.NewResult(userId, 1))

		testDat2 := &authEntities.UserSignToken{
			Id:       userId,
			Username: username,
		}

		userToken, err := authRepo.SignUsersAccessToken(testDat2)
		assert.NoError(t, err)
		assert.NotEmpty(t, userToken)
	})

}
