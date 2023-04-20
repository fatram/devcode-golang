package mysql

import (
	"context"
	"database/sql"
	"sync"

	"github.com/fatram/devcode-golang/domain/entity"
	"github.com/fatram/devcode-golang/domain/repository"
	"github.com/fatram/devcode-golang/internal/connector"
	"github.com/fatram/devcode-golang/pkg/genlog"
	"github.com/pkg/errors"
)

type activityRepositoryImpl struct {
	logger genlog.Logger
	db     *sql.DB
}

var (
	activityRepository     *activityRepositoryImpl
	activityRepositoryOnce sync.Once
)

func LoadActivityRepository(logger genlog.Logger) repository.ActivityRepository {
	activityRepositoryOnce.Do(func() {
		activityRepository = &activityRepositoryImpl{
			logger: logger,
			db:     connector.LoadMysqlDatabase(),
		}
	})
	return activityRepository
}

func (r *activityRepositoryImpl) Get(ctx context.Context, id int) (data *entity.Activity, err error) {
	query := `
		SELECT
			activity_id, title, email, created_at, updated_at
		FROM
			activities
		WHERE activity_id = ?
	`
	row := r.db.QueryRowContext(ctx, query, id)
	data = &entity.Activity{}
	err = row.Scan(&data.ID, &data.Title, &data.Email, &data.CreatedAt, &data.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, errors.Wrap(err, "error in activityRepositoryImpl, Get")
	}
	return data, nil
}

func (r *activityRepositoryImpl) GetAll(ctx context.Context) (data []entity.Activity, err error) {
	query := `
		SELECT
			activity_id, title, email, created_at, updated_at
		FROM
			activities
	`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, errors.Wrap(err, "error in activityRepositoryImpl, GetAll")
	}
	data = []entity.Activity{}
	for rows.Next() {
		item := entity.Activity{}
		err = rows.Scan(&item.ID, &item.Title, &item.Email, &item.CreatedAt, &item.UpdatedAt)
		if err != nil {
			return nil, errors.Wrap(err, "error in activityRepositoryImpl, GetAll")
		}
		data = append(data, item)
	}
	return data, nil
}

func (r activityRepositoryImpl) Create(ctx context.Context, data entity.Activity) (id int, err error) {
	query := `
		INSERT INTO activities (title, email, created_at, updated_at)
		VALUES (?, ?, ?, ?)
	`
	result, err := r.db.ExecContext(ctx, query, data.Title, data.Email, data.CreatedAt, data.UpdatedAt)
	if err != nil {
		return 0, errors.Wrap(err, "error in activityRepositoryImpl, Create")
	}
	insertID, err := result.LastInsertId()
	if err != nil {
		return 0, errors.Wrap(err, "error in activityRepositoryImpl, Create")
	}
	return int(insertID), nil
}

func (r activityRepositoryImpl) Delete(ctx context.Context, id int) (err error) {
	query := `
		DELETE FROM activities WHERE activity_id = ?
	`
	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return errors.Wrap(err, "error in activityRepositoryImpl, Delete")
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "error in activityRepositoryImpl, Delete")
	}
	if rowsAffected < 1 {
		return errors.New("error in activityRepositoryImpl, Delete")
	}
	return nil
}

func (r activityRepositoryImpl) Update(ctx context.Context, data entity.Activity) (err error) {
	query := `
		UPDATE activities
		SET
			title = ?,
			updated_at = ?
		WHERE activity_id = ?
	`
	result, err := r.db.ExecContext(ctx, query, data.Title, data.UpdatedAt, data.ID)
	if err != nil {
		return errors.Wrap(err, "error in activityRepositoryImpl, Update")
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "error in activityRepositoryImpl, Update")
	}
	if rowsAffected < 1 {
		return errors.New("error in activityRepositoryImpl, Update")
	}
	return nil
}
