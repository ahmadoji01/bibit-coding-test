package usecase_test

import (
	"go_bibit_test/domain"
	mocks "go_bibit_test/mocks/domain"
	"go_bibit_test/movie/usecase"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFetch(t *testing.T) {
	mockLogRepo := new(mocks.LogRepository)
	mockListOfMovies := make([]domain.Movie, 0)
	mockMovieUsecase := new(mocks.MovieUsecase)

	t.Run("success", func(t *testing.T) {
		mockLogRepo.On("Store", mock.Anything, mock.Anything).Return(nil).Once()
		mockMovieUsecase.On("Fetch", mock.AnythingOfType("string"),
			mock.AnythingOfType("int64")).Return(mockListOfMovies, nil).Once()

		mucase := usecase.NewMovieUsecase(mockLogRepo, time.Second*2)
		response, err := mucase.Fetch("Batman", 1)

		assert.NoError(t, err)
		assert.NotEmpty(t, response)
	})

	t.Run("error-failed", func(t *testing.T) {
		mockLogRepo.On("Store", mock.Anything, mock.Anything).Return(nil).Once()
		mockMovieUsecase.On("Fetch", mock.AnythingOfType("string"),
			mock.AnythingOfType("int64")).Return(mockListOfMovies, nil).Once()

		mucase := usecase.NewMovieUsecase(mockLogRepo, time.Second*2)
		response, err := mucase.Fetch("", 1)

		assert.Error(t, err)
		assert.Empty(t, response)
	})
}

func TestGetByID(t *testing.T) {
	mockLogRepo := new(mocks.LogRepository)
	mockListOfMovies := make([]domain.Movie, 0)
	mockMovieUsecase := new(mocks.MovieUsecase)

	t.Run("success", func(t *testing.T) {
		mockLogRepo.On("Store", mock.Anything, mock.Anything).Return(nil).Once()
		mockMovieUsecase.On("Fetch", mock.AnythingOfType("string"),
			mock.AnythingOfType("int64")).Return(mockListOfMovies, nil).Once()

		mucase := usecase.NewMovieUsecase(mockLogRepo, time.Second*2)
		response, err := mucase.GetByID("tt2313197")

		assert.NoError(t, err)
		assert.NotEmpty(t, response)
	})

	t.Run("error-failed", func(t *testing.T) {
		mockLogRepo.On("Store", mock.Anything, mock.Anything).Return(nil).Once()
		mockMovieUsecase.On("Fetch", mock.AnythingOfType("string"),
			mock.AnythingOfType("int64")).Return(mockListOfMovies, nil).Once()

		mucase := usecase.NewMovieUsecase(mockLogRepo, time.Second*2)
		response, err := mucase.GetByID("12345")

		assert.Error(t, err)
		assert.Empty(t, response)
	})
}
