package repository

import (
	"context"

	"github.com/fatram/devcode-golang/domain/entity"
)

type ActivityRepository interface {
	Get(ctx context.Context, id int) (data *entity.Activity, err error)
	GetAll(ctx context.Context) (data []entity.Activity, err error)
	Create(ctx context.Context, data entity.Activity) (id int, err error)
	Delete(ctx context.Context, id int) (err error)
	Update(ctx context.Context, data entity.Activity) (err error)
}

type TodoRepository interface {
	Get(ctx context.Context, id int) (data *entity.Todo, err error)
	GetAll(ctx context.Context, activityID *int) (data []entity.Todo, err error)
	Create(ctx context.Context, data entity.Todo) (id int, err error)
	Delete(ctx context.Context, id int) (err error)
	Update(ctx context.Context, data entity.Todo) (err error)
}
