package user

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/rendau/doklad_test/user/mocks"
	"github.com/rendau/doklad_test/user/model"
)

func TestGetUserName(t *testing.T) {
	// Создаем мок
	mockRepo := new(mocks.UserRepoI)

	// Настраиваем поведение мока
	mockRepo.On("GetUserByID", 1).Return(&model.User{ID: 1, Name: "John Doe"}, nil)

	// Создаем экземпляр UserHandler с моком
	handler := UserService{
		Repo: mockRepo,
	}

	// Вызываем тестируемую функцию
	result, err := handler.GetUserName(1)

	// Проверяем результаты
	require.NoError(t, err)
	require.Equal(t, "John Doe", result)

	// Проверяем, что все ожидания были выполнены
	mockRepo.AssertExpectations(t)
}

//func TestGetUserName_error(t *testing.T) {
//	mockRepo := new(mocks.UserRepoI)
//
//	someError := fmt.Errorf("some error")
//
//	mockRepo.On("GetUserByID", 1).Return(nil, someError)
//
//	handler := UserService{
//		Repo: mockRepo,
//	}
//
//	result, err := handler.GetUserName(1)
//
//	require.Error(t, err)
//	require.Equal(t, someError, err)
//	require.Empty(t, result)
//
//	mockRepo.AssertExpectations(t)
//}
//
//func TestGetUserName_notFound(t *testing.T) {
//	mockRepo := new(mocks.UserRepoI)
//
//	mockRepo.On("GetUserByID", 1).Return(nil, nil)
//
//	handler := UserService{
//		Repo: mockRepo,
//	}
//
//	result, err := handler.GetUserName(1)
//
//	require.Nil(t, err)
//	require.Equal(t, NoName, result)
//
//	mockRepo.AssertExpectations(t)
//}
