package repository_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/textures1245/go-template/internal/user/entities"
	"github.com/textures1245/go-template/internal/user/repository"
	"github.com/textures1245/go-template/internal/user/repository/repository_query"
	"go.uber.org/mock/gomock"
)

// TODO: implement user repository unit test

func Test_repo_CreateUser(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var (
		conn     = sqlx.NewDb(db, "sqlmock")
		userRepo = repository.NewUserRepository(conn)
	)

	userId := int64(1)
	testDat := &entities.UserCreatedReq{
		Name:        "test",
		Username:    "username_test",
		Password:    "pw_test",
		Email:       "test@gmail.com",
		PhoneNumber: "08123456789",
		IdCard:      "1234567890",
	}

	t.Run("CreateUser_positive_case", func(t *testing.T) {

		mock.
			ExpectExec(repository_query.InsertUser).
			WithArgs(testDat.Name, testDat.Username, testDat.Password, testDat.Email, testDat.PhoneNumber, testDat.IdCard).
			WillReturnResult(sqlmock.NewResult(userId, 1))

		userID, err := userRepo.CreateUser(context.Background(), testDat)
		assert.NoError(t, err)
		assert.NotEmpty(t, userID)
	})

	t.Run("CreateUser_negative_username_unique_constant_failed", func(t *testing.T) {

		mock.
			ExpectExec(repository_query.InsertUser).
			WithArgs(testDat.Name, testDat.Username, testDat.Password, testDat.Email, testDat.PhoneNumber, testDat.IdCard).
			WillReturnResult(sqlmock.NewResult(userId, 1))

		userID, err1 := userRepo.CreateUser(context.Background(), testDat)
		assert.NoError(t, err1)
		assert.NotNil(t, userID)

		//- creating user with same username
		mock.
			ExpectExec(repository_query.InsertUser).
			WithArgs(testDat.Name, testDat.Username, testDat.Password, testDat.Email, testDat.PhoneNumber, testDat.IdCard).
			WillReturnError(sql.ErrTxDone)

		userID, err2 := userRepo.CreateUser(context.Background(), testDat)
		assert.EqualError(t, err2, sql.ErrTxDone.Error())
		assert.Nil(t, userID)
	})

}

func Test_repo_UpdateUserById(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var (
		conn     = sqlx.NewDb(db, "sqlmock")
		userRepo = repository.NewUserRepository(conn)
	)

	userId := int64(1)
	testDat := &entities.UserCreatedReq{
		Name:        "test",
		Username:    "username_test",
		Password:    "pw_test",
		Email:       "test@gmail.com",
		PhoneNumber: "08123456789",
		IdCard:      "1234567890",
	}

	t.Run("UpdateUserById_positive_case", func(t *testing.T) {
		mock.
			ExpectExec(repository_query.InsertUser).
			WithArgs(testDat.Name, testDat.Username, testDat.Password, testDat.Email, testDat.PhoneNumber, testDat.IdCard).
			WillReturnResult(sqlmock.NewResult(userId, 1))

		userID, err1 := userRepo.CreateUser(context.Background(), testDat)
		assert.NoError(t, err1)
		assert.NotNil(t, userID)

		updUser := &entities.UserUpdateReq{
			Name:        "test",
			Email:       "update_test@gmail.com",
			PhoneNumber: "12312edasda",
			IdCard:      "12345s67890",
		}

		mock.
			ExpectExec(repository_query.UpdateUserById).
			WithArgs(updUser.Name, updUser.Email, updUser.IdCard, updUser.PhoneNumber, userId).
			WillReturnResult(sqlmock.NewResult(userId, 1))

		err := userRepo.UpdateUserById(context.Background(), userId, updUser)
		assert.NoError(t, err)
	})

}

func Test_repo_GetUserById(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var (
		conn     = sqlx.NewDb(db, "sqlmock")
		userRepo = repository.NewUserRepository(conn)
	)

	userId := int64(1)
	testDat := &entities.UserCreatedReq{
		Name:        "test",
		Username:    "username_test",
		Password:    "pw_test",
		Email:       "test@gmail.com",
		PhoneNumber: "08123456789",
		IdCard:      "1234567890",
	}

	t.Run("GetUserById_positive_case", func(t *testing.T) {
		mock.
			ExpectExec(repository_query.InsertUser).
			WithArgs(testDat.Name, testDat.Username, testDat.Password, testDat.Email, testDat.PhoneNumber, testDat.IdCard).
			WillReturnResult(sqlmock.NewResult(userId, 1))

		userID, err := userRepo.CreateUser(context.Background(), testDat)
		assert.NoError(t, err)
		assert.NotEmpty(t, userID)

		columnsQ := []string{"id", "name", "username", "password", "email", "phone_number", "id_card", "updated_at", "created_at"}
		currentTime := time.Now()

		mock.
			ExpectQuery(repository_query.FindUserById).
			WithArgs(userId).
			WillReturnRows(sqlmock.NewRows(columnsQ).AddRow(userId, testDat.Name, testDat.Username, testDat.Password, testDat.Email, testDat.PhoneNumber, testDat.IdCard, currentTime, currentTime))

		usr, err := userRepo.GetUserById(context.Background(), userId)
		assert.NoError(t, err)
		assert.NotEmpty(t, usr)
	})

	t.Run("GetUserById_negative_user_not_found", func(t *testing.T) {
		mock.
			ExpectExec(repository_query.InsertUser).
			WithArgs(testDat.Name, testDat.Username, testDat.Password, testDat.Email, testDat.PhoneNumber, testDat.IdCard).
			WillReturnResult(sqlmock.NewResult(userId, 1))

		userID, err := userRepo.CreateUser(context.Background(), testDat)
		assert.NoError(t, err)
		assert.NotEmpty(t, userID)

		invalidUserId := int64(-1)

		mock.
			ExpectQuery(repository_query.FindUserById).
			WithArgs(invalidUserId).
			WillReturnError(sql.ErrNoRows)

		usr, err := userRepo.GetUserById(context.Background(), invalidUserId)
		assert.EqualError(t, err, sql.ErrNoRows.Error())
		assert.Empty(t, usr)
	})

}
