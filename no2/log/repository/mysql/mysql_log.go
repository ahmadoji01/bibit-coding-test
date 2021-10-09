package mysql

import (
	"context"
	"database/sql"
	"go_bibit_test/domain"
	"go_bibit_test/log/repository"

	"github.com/sirupsen/logrus"
)

type mysqlLogRepository struct {
	Conn *sql.DB
}

// NewMysqlArticleRepository will create an object that represent the article.Repository interface
func NewMysqlLogRepository(Conn *sql.DB) domain.LogRepository {
	return &mysqlLogRepository{Conn}
}

func (m *mysqlLogRepository) Store(ctx context.Context, l *domain.Log) (err error) {
	query := `INSERT log SET action=?, created_at=?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, l.Action, l.CreatedAt)
	if err != nil {
		return
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		return err
	}
	l.ID = lastID
	return
}

func (m *mysqlLogRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []domain.Log, err error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			logrus.Error(errRow)
		}
	}()

	result = make([]domain.Log, 0)
	for rows.Next() {
		l := domain.Log{}
		err = rows.Scan(
			&l.ID,
			&l.Action,
			&l.CreatedAt,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		result = append(result, l)
	}

	return result, nil
}

func (m *mysqlLogRepository) Fetch(ctx context.Context, cursor string, num int64) (res []domain.Log, nextCursor string, err error) {
	query := `SELECT id,title,content, author_id, updated_at, created_at
  						FROM article WHERE created_at > ? ORDER BY created_at LIMIT ? `

	decodedCursor, err := repository.DecodeCursor(cursor)
	if err != nil && cursor != "" {
		return nil, "", domain.ErrBadParamInput
	}

	res, err = m.fetch(ctx, query, decodedCursor, num)
	if err != nil {
		return nil, "", err
	}

	if len(res) == int(num) {
		nextCursor = repository.EncodeCursor(res[len(res)-1].CreatedAt)
	}

	return
}
func (m *mysqlLogRepository) GetByID(ctx context.Context, id int64) (res domain.Log, err error) {
	query := `SELECT id,action,created_at
  						FROM log WHERE ID = ?`

	list, err := m.fetch(ctx, query, id)
	if err != nil {
		return domain.Log{}, err
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return res, domain.ErrNotFound
	}

	return
}
