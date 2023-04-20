package activity

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

type ActivityService struct {
	logger     genlog.Logger
	repository repository.ActivityRepository
}

func (s *ActivityService) Create(ctx context.Context, data interface{}) (interface{}, error) {
	activity, ok := data.(*model.ActivityCreate)
	if !ok {
		s.logger.Errorf("data tidak sesuai")
		return nil, echo.NewHTTPError(http.StatusBadRequest, "data tidak sesuai")
	}
	if activity.Title == "" {
		s.logger.Errorf("title cannot be null")
		return nil, echo.NewHTTPError(http.StatusBadRequest, "title cannot be null")
	}

	timeNow := int(time.Now().Unix())
	activityEntity := entity.Activity{
		Title:     activity.Title,
		Email:     activity.Email,
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
	}
	var err error
	activityEntity.ID, err = s.repository.Create(ctx, activityEntity)
	if err != nil {
		s.logger.Errorf("error at ActivityService.Create")
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "error at ActivityService.Create")
	}
	return model.Activity{
		ID:        activityEntity.ID,
		Title:     activityEntity.Title,
		Email:     activityEntity.Email,
		CreatedAt: time.Unix(int64(activityEntity.CreatedAt), 0).Format(time.RFC3339),
		UpdatedAt: time.Unix(int64(activityEntity.UpdatedAt), 0).Format(time.RFC3339),
	}, nil
}

func (s *ActivityService) Get(ctx context.Context, identifier interface{}) (interface{}, error) {
	id, _ := identifier.(string)
	s.logger.Print(id)
	intID, err := strconv.Atoi(id)
	if err != nil {
		s.logger.Errorf("error when get activity: %s", err.Error())
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "error at getting activity").SetInternal(err)
	}
	activity, err := s.repository.Get(ctx, intID)
	if err != nil {
		s.logger.Errorf("error when get activity: %s", err.Error())
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "error at getting activity").SetInternal(err)
	}
	if activity == nil {
		return nil, echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("Activity with ID %s Not Found", id))
	}
	response := model.Activity{
		ID:        activity.ID,
		Title:     activity.Title,
		Email:     activity.Email,
		CreatedAt: time.Unix(int64(activity.CreatedAt), 0).Format(time.RFC3339),
		UpdatedAt: time.Unix(int64(activity.UpdatedAt), 0).Format(time.RFC3339),
	}
	return response, nil
}

func (s *ActivityService) GetAll(ctx context.Context) (data []interface{}, err error) {
	activities, err := s.repository.GetAll(ctx)
	if err != nil {
		s.logger.Errorf("error when get all activities: %s", err.Error())
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "tidak dapat menghambil data activities").SetInternal(err)
	}
	data = make([]interface{}, len(activities))
	for i, activity := range activities {
		activityModel := model.Activity{
			ID:        activity.ID,
			Title:     activity.Title,
			Email:     activity.Email,
			CreatedAt: time.Unix(int64(activity.CreatedAt), 0).Format(time.RFC3339),
			UpdatedAt: time.Unix(int64(activity.UpdatedAt), 0).Format(time.RFC3339),
		}
		data[i] = activityModel
	}
	return data, err
}

func (s *ActivityService) Delete(ctx context.Context, identifier interface{}) error {
	id, _ := identifier.(string)
	intID, err := strconv.Atoi(id)
	if err != nil {
		s.logger.Errorf("error when delete activity: %s", err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "error at deleting activity").SetInternal(err)
	}
	err = s.repository.Delete(ctx, intID)
	if err != nil {
		s.logger.Errorf("error when delete activity: %s", err.Error())
		return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("Activity with ID %s Not Found", id)).SetInternal(err)
	}
	return nil
}

func (s *ActivityService) Update(ctx context.Context, identifier interface{}, data interface{}) (interface{}, error) {
	id, _ := identifier.(string)
	intID, err := strconv.Atoi(id)
	if err != nil {
		s.logger.Errorf("error when update activity: %s", err.Error())
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "error at updating activity").SetInternal(err)
	}
	activity, ok := data.(*model.ActivityUpdate)
	if !ok {
		s.logger.Errorf("data tidak sesuai")
		return nil, echo.NewHTTPError(http.StatusBadRequest, "data tidak sesuai")
	}
	if activity.Title == "" {
		s.logger.Errorf("title cannot be null")
		return nil, echo.NewHTTPError(http.StatusBadRequest, "title cannot be null")
	}

	timeNow := int(time.Now().Unix())
	activityEntity := entity.Activity{
		ID:        intID,
		Title:     activity.Title,
		UpdatedAt: timeNow,
	}
	err = s.repository.Update(ctx, activityEntity)
	if err != nil {
		s.logger.Errorf("error at ActivityService.Update")
		return nil, echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("Activity with ID %s Not Found", id))
	}
	updated, err := s.repository.Get(ctx, intID)
	if err != nil {
		s.logger.Errorf("error at ActivityService.Update")
		return nil, echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("Activity with ID %s Not Found", id))
	}
	return model.Activity{
		ID:        updated.ID,
		Title:     updated.Title,
		Email:     updated.Email,
		CreatedAt: time.Unix(int64(updated.CreatedAt), 0).Format(time.RFC3339),
		UpdatedAt: time.Unix(int64(updated.UpdatedAt), 0).Format(time.RFC3339),
	}, nil
}
