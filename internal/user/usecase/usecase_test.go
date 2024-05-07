package usecase

import (
	"context"
	"database/sql"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/textures1245/go-template/internal/user/dtos"
	"github.com/textures1245/go-template/internal/user/entities"
	mock_user "github.com/textures1245/go-template/internal/user/mock"
	"github.com/textures1245/go-template/pkg/apperror"

	"go.uber.org/mock/gomock"
)

// TODO: implement user usecase unit test
func Test_usecase_OnFetchUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var (
		ctx          = context.Background()
		userMockRepo = mock_user.NewMockUserRepository(ctrl)
		uc           = mock_user.NewMockUserUsecase(ctrl)
		currentTime  = time.Now()
	)

	usersReturn := []*entities.User{
		{
			Id:          1,
			Name:        "test",
			Username:    "test",
			Password:    "test",
			Email:       "test@gmail.com",
			IdCard:      "1234567890",
			PhoneNumber: "08123456789",
			UpdatedAt:   currentTime.String(),
			CreatedAt:   currentTime.String(),
		},
		{
			Id:          2,
			Name:        "test",
			Username:    "test",
			Password:    "test",
			Email:       "test@gmail.com",
			IdCard:      "1234567890",
			PhoneNumber: "08123456789",
			UpdatedAt:   currentTime.String(),
			CreatedAt:   currentTime.String(),
		},
	}

	t.Run("getUsers_positive_case", func(t *testing.T) {

		userMockRepo.EXPECT().GetUsers(ctx).Return(usersReturn, nil)

		users, err := userMockRepo.GetUsers(ctx)
		require.NotNil(t, users)
		require.NoError(t, err)

		var userDetailReturn []*dtos.UserDetailRespond
		for _, user := range users {
			userDetailReturn = append(userDetailReturn, entities.NewUserDetail(user))
		}

		uc.EXPECT().OnFetchUsers(ctx).Return(userDetailReturn, http.StatusOK, nil)
		res, status, err2 := uc.OnFetchUsers(ctx)
		var rawError error
		if err2 != nil {
			rawError = err2.RawError
		}
		require.NoError(t, rawError)
		require.Equal(t, http.StatusOK, status)
		require.NotEmpty(t, res)

	})

	t.Run("getUsers_negative_case", func(t *testing.T) {

		userMockRepo.EXPECT().GetUsers(ctx).Return(nil, sql.ErrNoRows)

		users, err := userMockRepo.GetUsers(ctx)
		require.Nil(t, users)
		require.Error(t, err)

		expectCode, cErr := apperror.CustomSqlExecuteHandler("User", err)

		uc.EXPECT().OnFetchUsers(ctx).Return(nil, expectCode, cErr)
		res, status, err2 := uc.OnFetchUsers(ctx)
		var rawError error
		if err2 != nil {
			rawError = err2.RawError
		}
		require.Error(t, rawError)
		require.EqualError(t, rawError, sql.ErrNoRows.Error())
		require.Equal(t, http.StatusNotFound, status)
		require.Nil(t, res)

	})

}
