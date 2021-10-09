package movie_grpc_handler

import (
	"log"

	"context"

	"go_bibit_test/domain"
	"go_bibit_test/movie/delivery/grpc/movie_grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func NewMovieServerGRPC(gserver *grpc.Server, movieUcase domain.MovieUsecase) {
	movieServer := server{
		usecase: movieUcase,
	}

	movie_grpc.RegisterMovieHandlerServer(gserver, &movieServer)
	reflection.Register(gserver)
}

type server struct {
	movie_grpc.UnimplementedMovieHandlerServer
	usecase domain.MovieUsecase
}

func (s *server) transformArticleRPC(mov *domain.Movie) *movie_grpc.Movie {

	if mov == nil {
		return nil
	}

	res := &movie_grpc.Movie{
		Title:     mov.Title,
		Year:      mov.Year,
		ImdbID:    mov.ImdbID,
		MovieType: mov.MovieType,
		Poster:    mov.Poster,

		Rated: mov.Rated,
		//Ratings:    mov.Ratings,
		Runtime:    mov.Runtime,
		Genre:      mov.Genre,
		Director:   mov.Director,
		Writer:     mov.Writer,
		Actors:     mov.Actors,
		Plot:       mov.Plot,
		Language:   mov.Language,
		Country:    mov.Country,
		Awards:     mov.Awards,
		Metascore:  mov.Metascore,
		ImdbRating: mov.ImdbRating,
		ImdbVotes:  mov.ImdbVotes,
		DVD:        mov.DVD,
		BoxOffice:  mov.BoxOffice,
		Production: mov.Production,
		Website:    mov.Website,
		Response:   mov.Response,
		Error:      mov.Error,
	}
	return res
}

func (s *server) transformArticleData(mov *movie_grpc.Movie) *domain.Movie {
	res := &domain.Movie{
		Title:     mov.Title,
		Year:      mov.Year,
		ImdbID:    mov.ImdbID,
		MovieType: mov.MovieType,
		Poster:    mov.Poster,

		Rated: mov.Rated,
		//Ratings:    mov.Ratings,
		Runtime:    mov.Runtime,
		Genre:      mov.Genre,
		Director:   mov.Director,
		Writer:     mov.Writer,
		Actors:     mov.Actors,
		Plot:       mov.Plot,
		Language:   mov.Language,
		Country:    mov.Country,
		Awards:     mov.Awards,
		Metascore:  mov.Metascore,
		ImdbRating: mov.ImdbRating,
		ImdbVotes:  mov.ImdbVotes,
		DVD:        mov.DVD,
		BoxOffice:  mov.BoxOffice,
		Production: mov.Production,
		Website:    mov.Website,
		Response:   mov.Response,
		Error:      mov.Error,
	}
	return res
}

func (s *server) GetMovie(ctx context.Context, in *movie_grpc.SingleRequest) (*movie_grpc.Movie, error) {
	id := ""
	if in != nil {
		id = in.Id
	}
	mov, err := s.usecase.GetByID(id)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	res := s.transformArticleRPC(&mov)
	return res, nil
}

func (s *server) FetchMovie(in *movie_grpc.FetchRequest, stream movie_grpc.MovieHandler_FetchMovieServer) error {
	searchword := ""
	pagination := int64(0)
	if in != nil {
		searchword = in.Searchword
		pagination = in.Pagination
	}
	movies, err := s.usecase.Fetch(searchword, pagination)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	for _, m := range movies {
		mov := s.transformArticleRPC(&m)

		if err := stream.Send(mov); err != nil {
			log.Println(err.Error())
			return err
		}
	}
	return nil
}
