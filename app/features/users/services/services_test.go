package services_test

import (
	"errors"
	"testing"

	users "Project_Title/app/features/users"
	"Project_Title/app/features/users/mocks"

	"github.com/stretchr/testify/assert"
)

func TestService_Login_Success(t *testing.T) {
	// Create a new instance of the mock Service
	service := mocks.NewService(t)

	// Set the expected return values
	expectedUser := users.Core{
		ID:       1,
		Email:    "example@example.com",
		Password: "password",
	}
	service.On("Login", "example@example.com", "password").Return(expectedUser, nil)

	// Call the Login function
	user, err := service.Login("example@example.com", "password")

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, user)

	// Verify that the expected function calls were made
	service.AssertExpectations(t)
}

func TestService_Login_Failure(t *testing.T) {
	// Create a new instance of the mock Service
	service := mocks.NewService(t)

	// Set the expected return values
	expectedError := errors.New("Invalid credentials")
	service.On("Login", "example@example.com", "password").Return(users.Core{}, expectedError)

	// Call the Login function
	user, err := service.Login("example@example.com", "password")

	// Assertions
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	assert.Equal(t, users.Core{}, user)

	// Verify that the expected function calls were made
	service.AssertExpectations(t)
}
