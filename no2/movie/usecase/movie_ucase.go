package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"go_bibit_test/domain"
	"net/http"
	"strconv"
	"time"
)

type movieUsecase struct {
	logRepo        domain.LogRepository
	contextTimeout time.Duration
}

type omdbSearchResponse struct {
	Response     string         `json:"Response"`
	TotalResults string         `json:"totalResults"`
	Error        string         `json:"Error"`
	Search       []domain.Movie `json:"Search"`
}

// NewArticleUsecase will create new an articleUsecase object representation of domain.ArticleUsecase interface
func NewMovieUsecase(l domain.LogRepository, timeout time.Duration) domain.MovieUsecase {
	return &movieUsecase{
		logRepo:        l,
		contextTimeout: timeout,
	}
}

func (m *movieUsecase) Fetch(query string, pagination int64) (res []domain.Movie, err error) {
	var omdbResp omdbSearchResponse
	resp, err := http.Get("http://www.omdbapi.com/?apikey=faf7e5bb&s=" + query + "&page=" + strconv.Itoa(int(pagination)))

	if resp.StatusCode != http.StatusOK {
		return make([]domain.Movie, 0), err
	}

	err = json.NewDecoder(resp.Body).Decode(&omdbResp)
	if err != nil {
		return make([]domain.Movie, 0), err
	}

	if omdbResp.Error != "" {
		return make([]domain.Movie, 0), fmt.Errorf(omdbResp.Error)
	}

	var log domain.Log
	log.Action = "GET Movies Search Word: " + query + " Page: " + strconv.Itoa(int(pagination))
	log.CreatedAt = time.Now()
	err = m.logRepo.Store(context.TODO(), &log)
	if err != nil {
		return make([]domain.Movie, 0), err
	}

	return omdbResp.Search, nil
}

func (m *movieUsecase) GetByID(id string) (res domain.Movie, err error) {
	var result domain.Movie
	resp, err := http.Get("http://www.omdbapi.com/?apikey=faf7e5bb&i=" + id)

	if resp.StatusCode != http.StatusOK {
		return domain.Movie{}, err
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return domain.Movie{}, err
	}

	if result.Error != "" {
		return domain.Movie{}, fmt.Errorf(result.Error)
	}

	var log domain.Log
	log.Action = "GET Movie IMDB ID: " + id
	log.CreatedAt = time.Now()
	err = m.logRepo.Store(context.TODO(), &log)
	if err != nil {
		return domain.Movie{}, err
	}

	return result, nil
}
