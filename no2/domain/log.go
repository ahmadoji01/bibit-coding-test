package domain

import (
	"context"
	"time"
)

type Log struct {
	ID        int64     `json:"id"`
	Action    string    `json:"action"`
	CreatedAt time.Time `json:"created_at"`
}

// ArticleUsecase represent the article's usecases
type LogUsecase interface {
	Fetch(ctx context.Context, cursor string, num int64) ([]Log, string, error)
	GetByID(ctx context.Context, id int64) (Log, error)
	Store(context.Context, *Log) error
}

// ArticleRepository represent the article's repository contract
type LogRepository interface {
	Fetch(ctx context.Context, cursor string, num int64) (res []Log, nextCursor string, err error)
	GetByID(ctx context.Context, id int64) (Log, error)
	Store(ctx context.Context, l *Log) error
}
