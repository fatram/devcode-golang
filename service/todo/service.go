package todo

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/fatram/devcode-golang/domain/entity"
	"github.com/fatram/devcode-golang/domain/model"
	"github.com/fatram/devcode-golang/domain/repository"
	"github.com/fatram/devcode-golang/pkg/genlog"
	"github.com/labstack/echo/v4"
)

type TodoService struct {
	logger     genlog.Logger
	repository repository.TodoRepository
}

func (s *TodoService) Create(ctx context.Context, data interface{}) (interface{}, error) {
	todo, ok := data.(*model.TodoCreate)
	if !ok {
		s.logger.Errorf("data tidak sesuai")
		return nil, echo.NewHTTPError(http.StatusBadRequest, "data tidak sesuai")
	}
	if todo.Title == "" {
		s.logger.Errorf("title cannot be null")
		return nil, echo.NewHTTPError(http.StatusBadRequest, "title cannot be null")
	}
	if todo.ActivityID < 1 {
		s.logger.Errorf("title cannot be null")
		return nil, echo.NewHTTPError(http.StatusBadRequest, "activity_group_id cannot be null")
	}

	timeNow := int(time.Now().Unix())
	todoEntity := entity.Todo{
		ActivityID: todo.ActivityID,
		Title:      todo.Title,
		Priority:   "very-high",
		IsActive:   true,
		CreatedAt:  timeNow,
		UpdatedAt:  timeNow,
	}
	if todo.IsActive != nil {
		todoEntity.IsActive = *todo.IsActive
	}
	if todo.Priority != "" {
		todoEntity.Priority = todo.Priority
	}
	var err error
	todoEntity.ID, err = s.repository.Create(ctx, todoEntity)
	if err != nil {
		s.logger.Errorf("error at TodoService.Create")
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "error at TodoService.Create")
	}
	return model.Todo{
		ID:         todoEntity.ID,
		ActivityID: todoEntity.ActivityID,
		Title:      todoEntity.Title,
		Priority:   todoEntity.Priority,
		IsActive:   todoEntity.IsActive,
		CreatedAt:  time.Unix(int64(todoEntity.CreatedAt), 0).Format(time.RFC3339),
		UpdatedAt:  time.Unix(int64(todoEntity.UpdatedAt), 0).Format(time.RFC3339),
	}, nil
}

func (s *TodoService) Get(ctx context.Context, identifier interface{}) (interface{}, error) {
	stringID, _ := identifier.(string)
	id, err := strconv.Atoi(stringID)
	if err != nil {
		s.logger.Errorf("error when get todo: %s", err.Error())
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "error at getting todo").SetInternal(err)
	}
	todo, err := s.repository.Get(ctx, id)
	if err != nil {
		s.logger.Errorf("error when get todo: %s", err.Error())
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "error at getting todo").SetInternal(err)
	}
	if todo == nil {
		return nil, echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("Todo with ID %d Not Found", id))
	}
	response := model.Todo{
		ID:         todo.ID,
		ActivityID: todo.ActivityID,
		Title:      todo.Title,
		Priority:   todo.Priority,
		IsActive:   todo.IsActive,
		CreatedAt:  time.Unix(int64(todo.CreatedAt), 0).Format(time.RFC3339),
		UpdatedAt:  time.Unix(int64(todo.UpdatedAt), 0).Format(time.RFC3339),
	}
	return response, nil
}

func (s *TodoService) GetAll(ctx context.Context, filter model.ToDoFilter) (data []interface{}, err error) {
	activities, err := s.repository.GetAll(ctx, filter.ActivityID)
	if err != nil {
		s.logger.Errorf("error when get all todos: %s", err.Error())
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "tidak dapat menghambil data todos").SetInternal(err)
	}
	data = make([]interface{}, len(activities))
	for i, todo := range activities {
		todoModel := model.Todo{
			ID:         todo.ID,
			ActivityID: todo.ActivityID,
			Title:      todo.Title,
			Priority:   todo.Priority,
			IsActive:   todo.IsActive,
			CreatedAt:  time.Unix(int64(todo.CreatedAt), 0).Format(time.RFC3339),
			UpdatedAt:  time.Unix(int64(todo.UpdatedAt), 0).Format(time.RFC3339),
		}
		data[i] = todoModel
	}
	return data, err
}

func (s *TodoService) Delete(ctx context.Context, identifier interface{}) error {
	stringID, _ := identifier.(string)
	id, err := strconv.Atoi(stringID)
	if err != nil {
		s.logger.Errorf("error when delete todo: %s", err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "error at deleting todo").SetInternal(err)
	}
	err = s.repository.Delete(ctx, id)
	if err != nil {
		s.logger.Errorf("error when delete todo: %s", err.Error())
		return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("Todo with ID %d Not Found", id)).SetInternal(err)
	}
	return nil
}

func (s *TodoService) Update(ctx context.Context, identifier interface{}, data interface{}) (interface{}, error) {
	stringID, _ := identifier.(string)
	id, err := strconv.Atoi(stringID)
	if err != nil {
		s.logger.Errorf("error when update todo: %s", err.Error())
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "error at updating todo").SetInternal(err)
	}
	todoEntity, err := s.repository.Get(ctx, id)
	if err != nil {
		s.logger.Errorf("error at TodoService.Update")
		return nil, echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("Todo with ID %d Not Found", id))
	}
	if todoEntity == nil {
		s.logger.Errorf("error at TodoService.Update")
		return nil, echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("Todo with ID %d Not Found", id))
	}
	todo, ok := data.(*model.TodoUpdate)
	if !ok {
		s.logger.Errorf("data tidak sesuai")
		return nil, echo.NewHTTPError(http.StatusBadRequest, "data tidak sesuai")
	}

	timeNow := int(time.Now().Unix())
	todoEntity.UpdatedAt = timeNow
	if todo.Title != "" {
		todoEntity.Title = todo.Title
	}
	if todo.Priority != "" {
		todoEntity.Priority = todo.Priority
	}
	if todo.IsActive != nil {
		todoEntity.IsActive = *todo.IsActive
	}
	s.logger.Print("jaka")
	err = s.repository.Update(ctx, *todoEntity)
	if err != nil {
		s.logger.Errorf("error at TodoService.Update")
		return nil, echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("Todo with ID %d Not Found", id))
	}
	return model.Todo{
		ID:         todoEntity.ID,
		ActivityID: todoEntity.ActivityID,
		Title:      todoEntity.Title,
		Priority:   todoEntity.Priority,
		IsActive:   todoEntity.IsActive,
		CreatedAt:  time.Unix(int64(todoEntity.CreatedAt), 0).Format(time.RFC3339),
		UpdatedAt:  time.Unix(int64(todoEntity.UpdatedAt), 0).Format(time.RFC3339),
	}, nil
}
