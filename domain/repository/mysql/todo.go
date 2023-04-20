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

type todoRepositoryImpl struct {
	logger genlog.Logger
	db     *sql.DB
}

var (
	todoRepository     *todoRepositoryImpl
	todoRepositoryOnce sync.Once
)

func LoadTodoRepository(logger genlog.Logger) repository.TodoRepository {
	todoRepositoryOnce.Do(func() {
		todoRepository = &todoRepositoryImpl{
			logger: logger,
			db:     connector.LoadMysqlDatabase(),
		}
	})
	return todoRepository
}

func (r *todoRepositoryImpl) Get(ctx context.Context, id int) (data *entity.Todo, err error) {
	query := `
		SELECT
			todo_id, activity_group_id, title, is_active, priority, created_at, updated_at
		FROM
			todos
		WHERE todo_id = ?
	`
	row := r.db.QueryRowContext(ctx, query, id)
	data = &entity.Todo{}
	err = row.Scan(&data.ID, &data.ActivityID, &data.Title, &data.IsActive, &data.Priority, &data.CreatedAt, &data.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, errors.Wrap(err, "error in todoRepositoryImpl, Get")
	}
	return data, nil
}

func (r *todoRepositoryImpl) GetAll(ctx context.Context, activityID string) (data []entity.Todo, err error) {
	query := `
	SELECT
		todo_id, activity_group_id, title, is_active, priority, created_at, updated_at
	FROM
		todos
	`
	args := []interface{}{}
	if activityID != "" {
		query += " WHERE activity_group_id = ? "
		args = append(args, activityID)
	}
	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, errors.Wrap(err, "error in todoRepositoryImpl, GetAll")
	}
	data = []entity.Todo{}
	for rows.Next() {
		item := entity.Todo{}
		err = rows.Scan(&item.ID, &item.ActivityID, &item.Title, &item.IsActive, &item.Priority, &item.CreatedAt, &item.UpdatedAt)
		if err != nil {
			return nil, errors.Wrap(err, "error in todoRepositoryImpl, GetAll")
		}
		data = append(data, item)
	}
	return data, nil
}

func (r todoRepositoryImpl) Create(ctx context.Context, data entity.Todo) (id int, err error) {
	query := `
		INSERT INTO todos (activity_group_id, title, is_active, priority, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?)
	`
	result, err := r.db.ExecContext(ctx, query, data.ActivityID, data.Title, data.IsActive, data.Priority, data.CreatedAt, data.UpdatedAt)
	if err != nil {
		return 0, errors.Wrap(err, "error in todoRepositoryImpl, Create")
	}
	insertID, err := result.LastInsertId()
	if err != nil {
		return 0, errors.Wrap(err, "error in todoRepositoryImpl, Create")
	}
	return int(insertID), nil
}

func (r todoRepositoryImpl) Delete(ctx context.Context, id int) (err error) {
	query := `
		DELETE FROM todos WHERE todo_id = ?
	`
	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return errors.Wrap(err, "error in todoRepositoryImpl, Delete")
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "error in todoRepositoryImpl, Delete")
	}
	if rowsAffected < 1 {
		return errors.New("error in todoRepositoryImpl, Delete")
	}
	return nil
}

func (r todoRepositoryImpl) Update(ctx context.Context, data entity.Todo) (err error) {
	query := `
		UPDATE todos
		SET
			title = ?,
			priority = ?,
			is_active = ?,
			updated_at = ?
		WHERE todo_id = ?
	`
	result, err := r.db.ExecContext(ctx, query, data.Title, data.Priority, data.IsActive, data.UpdatedAt, data.ID)
	if err != nil {
		return errors.Wrap(err, "error in todoRepositoryImpl, Update")
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "error in todoRepositoryImpl, Update")
	}
	if rowsAffected < 1 {
		return errors.New("error in todoRepositoryImpl, Update")
	}
	return nil
}
