package usecase_test

import (
	"context"
	"database/sql"
	"net/http"

	"testing"

	"github.com/gofiber/fiber/v2/log"
	"github.com/stretchr/testify/require"
	"github.com/textures1245/go-template/internal/auth/dtos"
	"github.com/textures1245/go-template/internal/auth/entities"
	mock_auth "github.com/textures1245/go-template/internal/auth/mock"
	userEntities "github.com/textures1245/go-template/internal/user/entities"
	mock_user "github.com/textures1245/go-template/internal/user/mock"
	"github.com/textures1245/go-template/pkg/apperror"
	"go.uber.org/mock/gomock"
	"golang.org/x/crypto/bcrypt"
)

// func hashPassword(userPw string, reqPw string) error {
// 	if err := bcrypt.CompareHashAndPassword([]byte(userPw), []byte(reqPw)); err != nil {
// 		return err
// 	}
// 	return nil
// }

func generatePasswordHash(password string) ([]byte, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return hashedPassword, nil
}

// TODO: implement auth usecase unit test
// TODO: Test_usecase_Register

func Test_usecase_Login(t *testing.T) {

	// db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	// require.NoError(t, err)
	// defer db.Close()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var (
		// conn = sqlx.NewDb(db, "sqlmock")
		// userRepo       = userRepo.NewUserRepository(conn)
		username       = "username_test"
		password       = "password_test"
		userId   int64 = 1
	)

	ctx := context.Background()
	userMockRepo := mock_user.NewMockUserRepository(ctrl)
	aMockUse := mock_auth.NewMockAuthUsecase(ctrl)

	hashPassword, _ := generatePasswordHash(password)

	userCred := &entities.UsersCredentials{
		Username: username,
		Password: string(hashPassword),
	}

	userPassport := &entities.UsersPassport{
		Id:       userId,
		Username: username,
		Password: string(hashPassword),
	}

	t.Run("login_positive_case", func(t *testing.T) {
		// mock.
		// 	ExpectExec(repository_query.InsertUser).
		// 	WithArgs(userCred.Name, userCred.Username, userCred.Password, userCred.Email, userCred.PhoneNumber, userCred.IdCard).
		// 	WillReturnResult(sqlmock.NewResult(userId, 1))

		// userID, error := userRepo.CreateUser(context.Background(), userCred)
		// assert.NoError(t, error)
		// assert.NotEmpty(t, userID)

		userMockRepo.EXPECT().FindUserAsPassport(ctx, userCred.Username).Return(userPassport, nil)
		usrPassport, pError := userMockRepo.FindUserAsPassport(ctx, userCred.Username)
		require.NoError(t, pError)
		require.NotNil(t, usrPassport)

		aMockUse.EXPECT().Login(ctx, userCred).Return(&dtos.UserTokenRes{
			TokenType: "Authorization",
		}, http.StatusOK, nil)
		userToken, status, err := aMockUse.Login(ctx, userCred)
		if err != nil {
			require.NoError(t, err.RawError)
			require.NoError(t, err.CError)
		}
		require.Equal(t, status, http.StatusOK)
		require.NotNil(t, userToken)
	})

	t.Run("login_negative_password_invalid", func(t *testing.T) {
		// mock.
		// 	ExpectExec(repository_query.InsertUser).
		// 	WithArgs(userCred.Name, userCred.Username, userCred.Password, userCred.Email, userCred.PhoneNumber, userCred.IdCard).
		// 	WillReturnResult(sqlmock.NewResult(userId, 1))

		// userID, error := userRepo.CreateUser(context.Background(), userCred)
		// assert.NoError(t, error)
		// assert.NotEmpty(t, userID)

		userMockRepo.EXPECT().FindUserAsPassport(ctx, userCred.Username).Return(userPassport, nil)
		usrPassport, pError := userMockRepo.FindUserAsPassport(ctx, userCred.Username)
		require.NotNil(t, usrPassport)
		require.NoError(t, pError)

		invalidPw := "invalid_password"

		aMockUse.EXPECT().Login(ctx, &entities.UsersCredentials{
			Username: username,
			Password: invalidPw,
		}).Return(&dtos.UserTokenRes{}, http.StatusBadRequest, &apperror.CErr{
			RawError: nil,
			CError:   apperror.ErrorInvalidCredentials,
		})
		userToken, status, err := aMockUse.Login(ctx, &entities.UsersCredentials{
			Username: username,
			Password: invalidPw,
		})
		if err != nil {
			require.EqualError(t, err.CError, apperror.ErrorInvalidCredentials.Error())
		}
		require.Equal(t, status, http.StatusBadRequest)
		require.Empty(t, userToken)
	})

	t.Run("login_negative_failed_query_username", func(t *testing.T) {
		// mock.
		// 	ExpectExec(repository_query.InsertUser).
		// 	WithArgs(userCred.Name, userCred.Username, userCred.Password, userCred.Email, userCred.PhoneNumber, userCred.IdCard).
		// 	WillReturnResult(sqlmock.NewResult(userId, 1))

		// userID, error := userRepo.CreateUser(context.Background(), userCred)
		// assert.NoError(t, error)
		// assert.NotEmpty(t, userID)

		userMockRepo.EXPECT().FindUserAsPassport(ctx, userCred.Username).Return(&entities.UsersPassport{}, sql.ErrNoRows)
		usrPassport, pError := userMockRepo.FindUserAsPassport(ctx, userCred.Username)
		require.Equal(t, &entities.UsersPassport{}, usrPassport)
		require.EqualError(t, pError, sql.ErrNoRows.Error())

		aMockUse.EXPECT().Login(ctx, userCred).Return(&dtos.UserTokenRes{}, 404, &apperror.CErr{
			RawError: sql.ErrNoRows,
		})
		userToken, status, err := aMockUse.Login(ctx, userCred)
		log.Debug(err)
		require.EqualError(t, err.RawError, sql.ErrNoRows.Error())
		require.Equal(t, status, http.StatusNotFound)
		require.Equal(t, userToken, &dtos.UserTokenRes{})
	})

}

func Test_usecase_Register(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var (
		username       = "username_test"
		password       = "password_test"
		userId   int64 = 1
	)

	ctx := context.Background()
	userMockRepo := mock_user.NewMockUserRepository(ctrl)
	aMockUse := mock_auth.NewMockAuthUsecase(ctrl)

	userCreated := &userEntities.UserCreatedReq{
		Name:        "test",
		Username:    username,
		Password:    password,
		Email:       "test@gmail.com",
		PhoneNumber: "08123456789",
		IdCard:      "1234567890",
	}

	t.Run("register_positive_case", func(t *testing.T) {
		userMockRepo.EXPECT().CreateUser(ctx, userCreated).Return(&userId, nil)
		usrId, err := userMockRepo.CreateUser(ctx, userCreated)
		require.NoError(t, err)
		require.NotNil(t, usrId)

		aMockUse.EXPECT().Register(ctx, userCreated).Return(&dtos.UsersRegisteredRes{
			AccessToken: "Authorization",
		}, http.StatusOK, nil)
		userToken, status, cErr := aMockUse.Register(ctx, userCreated)

		var rawError error
		if cErr != nil {
			rawError = cErr.RawError
		}
		require.NoError(t, rawError)
		require.Equal(t, status, http.StatusOK)
		require.NotNil(t, userToken)
	})

	t.Run("register_negative_failed_unique_username_created", func(t *testing.T) {

		aMockUse.EXPECT().Register(ctx, userCreated).Return(&dtos.UsersRegisteredRes{}, 404, &apperror.CErr{
			RawError: sql.ErrNoRows,
		})
		userToken, status, cErr2 := aMockUse.Register(ctx, userCreated)
		var rawError2 error
		if cErr2 != nil {
			rawError2 = cErr2.RawError
		}
		require.EqualError(t, rawError2, sql.ErrNoRows.Error())
		require.Equal(t, status, http.StatusNotFound)
		require.Equal(t, userToken, &dtos.UsersRegisteredRes{})

		userMockRepo.EXPECT().CreateUser(ctx, userCreated).Return(nil, sql.ErrNoRows)
		usrId, err := userMockRepo.CreateUser(ctx, userCreated)
		require.Nil(t, usrId)
		require.Error(t, err)

	})
}
