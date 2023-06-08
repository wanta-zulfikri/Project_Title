package services_test

import (
	"errors"
	"testing"

	books "Project_Title/app/features/books"
	"Project_Title/app/features/books/mocks"

	"github.com/stretchr/testify/assert"
)

func TestService_CreateBook_Success(t *testing.T) {
	// Create a new instance of the mock Service
	service := mocks.NewService(t)

	// Set the expected return values
	expectedBook := books.Core{
		ID:           1,
		Title:        "Book Title",
		PublishedYear: "2023",
		ISBN:         "1234567890",
	}
	service.On("CreateBook", expectedBook, uint(1)).Return(nil)

	// Call the CreateBook function
	err := service.CreateBook(expectedBook, uint(1))

	// Assertions
	assert.NoError(t, err)

	// Verify that the expected function calls were made
	service.AssertExpectations(t)
}

func TestService_CreateBook_Failure(t *testing.T) {
	// Create a new instance of the mock Service
	service := mocks.NewService(t)

	// Set the expected return values
	expectedError := errors.New("Failed to create book")
	service.On("CreateBook", books.Core{}, uint(1)).Return(expectedError)

	// Call the CreateBook function
	err := service.CreateBook(books.Core{}, uint(1))

	// Assertions
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)

	// Verify that the expected function calls were made
	service.AssertExpectations(t)
}

func TestService_GetBook_Success(t *testing.T) {
	// Create a new instance of the mock Service
	service := mocks.NewService(t)

	// Set the expected return values
	expectedBook := books.Core{
		ID:           1,
		Title:        "Book Title",
		PublishedYear: "2023",
		ISBN:         "1234567890",
	}
	service.On("GetBook", uint(1)).Return(expectedBook, nil)

	// Call the GetBook function
	book, err := service.GetBook(uint(1))

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, expectedBook, book)

	// Verify that the expected function calls were made
	service.AssertExpectations(t)
}

func TestService_GetBook_Failure(t *testing.T) {
	// Create a new instance of the mock Service
	service := mocks.NewService(t)

	// Set the expected return values
	expectedError := errors.New("Failed to get book")
	service.On("GetBook", uint(1)).Return(books.Core{}, expectedError)

	// Call the GetBook function
	book, err := service.GetBook(uint(1))

	// Assertions
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	assert.Equal(t, books.Core{}, book)

	// Verify that the expected function calls were made
	service.AssertExpectations(t)
}

// Test the remaining functions in a similar way...

