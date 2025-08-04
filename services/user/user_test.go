package service_user_test

import (
	"context"
	"errors"
	"golang-api/mocks"
	"golang-api/repositories/models"
	"golang-api/services/entities/request"
	service_user "golang-api/services/user"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestUserService_GetById(t *testing.T) {
	t.Run("TestUserService_GetById_Expect_Success", func(t *testing.T) {
		mockRepo := new(mocks.UserRepository)
		svc := service_user.NewUserService(mockRepo)

		expected := &models.User{Id: 1, Name: "test1", Email: "test1@mail.com"}

		mockRepo.On("GetById", mock.Anything, uint(1)).Return(expected, nil)

		user, err := svc.GetById(context.Background(), 1)
		log.Println(user)
		assert.NoError(t, err)
		assert.Equal(t, expected, user)
		mockRepo.AssertExpectations(t)
	})

	t.Run("TestUserService_GetById_Expect_Fail", func(t *testing.T) {
		mockRepo := new(mocks.UserRepository)
		svc := service_user.NewUserService(mockRepo)

		// Mock => return error
		mockRepo.
			On("GetById", mock.Anything, uint(1)).
			Return(nil, errors.New("user not found"))

		user, err := svc.GetById(context.Background(), 1)

		assert.Nil(t, user)
		assert.Error(t, err)
		assert.EqualError(t, err, "user not found")
		mockRepo.AssertExpectations(t)
	})

}

func TestUserService_List(t *testing.T) {
	t.Run("TestUserService_List_Expect_Success", func(t *testing.T) {
		mockRepo := new(mocks.UserRepository)
		svc := service_user.NewUserService(mockRepo)

		expected := &[]models.User{
			{Id: 1, Name: "test1", Email: "test1@mail.com"},
		}

		mockRepo.On("List", mock.Anything).Return(expected, nil)

		users, err := svc.List(context.Background())

		assert.NoError(t, err)
		assert.Equal(t, expected, users)
		mockRepo.AssertExpectations(t)
	})

	t.Run("TestUserService_List_Expect_Fail", func(t *testing.T) {
		mockRepo := new(mocks.UserRepository)
		svc := service_user.NewUserService(mockRepo)

		mockRepo.On("List", mock.Anything).Return(nil, errors.New("failed to get users"))

		users, err := svc.List(context.Background())

		assert.Nil(t, users)
		assert.Error(t, err)
		assert.EqualError(t, err, "failed to get users")
	})
}

func TestUserService_Create(t *testing.T) {
	input := request.UserReq{
		Name:  "test1",
		Email: "test1@mail.com",
	}

	t.Run("TestUserService_Create_Expect_Success", func(t *testing.T) {
		mockRepo := new(mocks.UserRepository)
		svc := service_user.NewUserService(mockRepo)

		expected := &models.User{Id: 1, Name: "test1", Email: "test1@mail.com"}

		mockRepo.On("Create", mock.Anything, mock.MatchedBy(func(u *models.User) bool {
			return u.Name == input.Name && u.Email == input.Email
		})).Return(expected, nil)

		result, err := svc.Create(context.Background(), input)

		assert.NoError(t, err)
		assert.Equal(t, expected, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("TestUserService_Create_Expect_Fail", func(t *testing.T) {
		mockRepo := new(mocks.UserRepository)
		svc := service_user.NewUserService(mockRepo)

		mockRepo.
			On("Create", mock.Anything, mock.MatchedBy(func(u *models.User) bool {
				return u.Name == input.Name && u.Email == input.Email
			})).
			Return(nil, errors.New("failed to create user"))

		result, err := svc.Create(context.Background(), input)

		assert.Nil(t, result)
		assert.Error(t, err)
		assert.EqualError(t, err, "failed to create user")
		mockRepo.AssertExpectations(t)
	})
}

func TestUserService_Update(t *testing.T) {
	input := request.UserUpdateReq{
		Name:  "test1",
		Email: "test1@mail.com",
	}
	t.Run("TestUserService_Update_Expect_Success", func(t *testing.T) {
		mockRepo := new(mocks.UserRepository)
		svc := service_user.NewUserService(mockRepo)

		mockRepo.On("UpdateById", mock.Anything, uint(1), mock.Anything).Return(nil)

		err := svc.Update(context.Background(), 1, input)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("TestUserService_Update_Expect_Fail", func(t *testing.T) {
		mockRepo := new(mocks.UserRepository)
		svc := service_user.NewUserService(mockRepo)

		mockRepo.
			On("UpdateById", mock.Anything, uint(1), mock.Anything).
			Return(errors.New("update fail"))

		err := svc.Update(context.Background(), 1, input)

		assert.Error(t, err)
		assert.EqualError(t, err, "update fail")
		mockRepo.AssertExpectations(t)
	})

}

func TestUserService_Delete(t *testing.T) {

	t.Run("TestUserService_Delete_Expect_Success", func(t *testing.T) {
		mockRepo := new(mocks.UserRepository)
		svc := service_user.NewUserService(mockRepo)

		mockRepo.On("DeleteById", mock.Anything, uint(2)).Return(nil)

		err := svc.Delete(context.Background(), 2)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)

	})

	t.Run("TestUserService_Delete_Expect_Fail", func(t *testing.T) {
		mockRepo := new(mocks.UserRepository)
		svc := service_user.NewUserService(mockRepo)

		mockRepo.
			On("DeleteById", mock.Anything, uint(2)).
			Return(errors.New("delete failed"))

		err := svc.Delete(context.Background(), 2)

		assert.Error(t, err)
		assert.EqualError(t, err, "delete failed")
		mockRepo.AssertExpectations(t)
	})

	t.Run("TestUserService_Delete_Expect_NotFound", func(t *testing.T) {
		mockRepo := new(mocks.UserRepository)
		svc := service_user.NewUserService(mockRepo)

		mockRepo.
			On("DeleteById", mock.Anything, uint(100)).
			Return(gorm.ErrRecordNotFound)

		err := svc.Delete(context.Background(), 100)

		assert.Error(t, err)
		assert.Equal(t, gorm.ErrRecordNotFound, err)
		mockRepo.AssertExpectations(t)
	})
}
