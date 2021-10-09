package http

import (
	"net/http"
	"strconv"

	"go_bibit_test/domain"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

// ResponseError represent the reseponse error struct
type ResponseError struct {
	Message string `json:"message"`
}

// ArticleHandler  represent the httphandler for article
type MovieHandler struct {
	MUsecase domain.MovieUsecase
}

// NewArticleHandler will initialize the articles/ resources endpoint
func NewMovieHandler(e *echo.Echo, us domain.MovieUsecase) {
	handler := &MovieHandler{
		MUsecase: us,
	}
	e.GET("/movies", handler.FetchMovie)
	e.GET("/movies/:id", handler.GetByID)
}

// FetchArticle will fetch the article based on given params
func (m *MovieHandler) FetchMovie(c echo.Context) error {
	query := c.QueryParam("searchword")
	pagination, err := strconv.Atoi(c.QueryParam("pagination"))

	if err != nil {
		pagination = 1
	}

	movies, err := m.MUsecase.Fetch(query, int64(pagination))
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, movies)
}

// GetByID will get article by given id
func (m *MovieHandler) GetByID(c echo.Context) error {
	idP := c.Param("id")

	mov, err := m.MUsecase.GetByID(idP)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, mov)
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)
	switch err {
	case domain.ErrInternalServerError:
		return http.StatusInternalServerError
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
