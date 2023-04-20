package activity

import (
	"net/http"

	"github.com/fatram/devcode-golang/domain/model"
	"github.com/fatram/devcode-golang/pkg/genlog"
	service "github.com/fatram/devcode-golang/service/activity"
	"github.com/labstack/echo/v4"
)

type GetActivityControllerImpl struct {
	Service service.ActivityService
}

func (ctr *GetActivityControllerImpl) Get(c echo.Context) (err error) {
	ctx := c.Request().Context()
	id := c.Param("id")
	data, err := ctr.Service.Get(ctx, id)
	if err != nil {
		c.Logger().Errorf("Error on GetActivityControllerImpl.Get: %s", err.Error())
		return err
	}
	response := model.BaseResponse{
		Status:  "Success",
		Message: "Success",
		Data:    data,
	}
	return c.JSON(http.StatusOK, response)
}

func newGetActivityController(logger genlog.Logger) *GetActivityControllerImpl {
	return &GetActivityControllerImpl{
		Service: *service.LoadActivityService(logger),
	}
}

type GetAllActivityControllerImpl struct {
	Service service.ActivityService
}

func (ctr *GetAllActivityControllerImpl) GetAll(c echo.Context) (err error) {
	data, err := ctr.Service.GetAll(c.Request().Context())
	if err != nil {
		c.Logger().Errorf("Error on GetAllActivityControllerImpl.GetAllActivity: %s", err.Error())
		return err
	}

	response := model.BaseResponse{
		Status:  "Success",
		Message: "Success",
		Data:    data,
	}
	return c.JSON(http.StatusOK, response)
}

func newGetAllActivityController(logger genlog.Logger) *GetAllActivityControllerImpl {
	return &GetAllActivityControllerImpl{
		Service: *service.LoadActivityService(logger),
	}
}

type CreateActivityControllerImpl struct {
	Service  service.ActivityService
	Bind     model.BindFunc
	Validate model.ValideFunc
}

func (ctr *CreateActivityControllerImpl) Create(c echo.Context) (err error) {
	data, err := ctr.Bind(c)
	if err != nil {
		c.Logger().Errorf("Error on CreateActivityControllerImpl.Create: %s", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body").SetInternal(err)
	}
	if err := ctr.Validate(c, data); err != nil {
		c.Logger().Errorf("Validation error on CreateActivityControllerImpl.Create: %s", err.Error())
		return err
	}

	created, err := ctr.Service.Create(c.Request().Context(), data)
	if err != nil {
		c.Logger().Errorf("Error on CreateActivityControllerImpl.Create: %s", err.Error())
		return err
	}
	response := model.BaseResponse{
		Status:  "Success",
		Message: "Success",
		Data:    created,
	}
	return c.JSON(http.StatusCreated, response)
}

func newCreateActivityController(logger genlog.Logger) *CreateActivityControllerImpl {
	return &CreateActivityControllerImpl{
		Service:  *service.LoadActivityService(logger),
		Bind:     model.BindActivityCreate,
		Validate: model.ValidateActivityCreate,
	}
}

type DeleteActivityControllerImpl struct {
	Service service.ActivityService
}

func (ctr *DeleteActivityControllerImpl) Delete(c echo.Context) (err error) {
	ctx := c.Request().Context()
	id := c.Param("id")
	err = ctr.Service.Delete(ctx, id)
	if err != nil {
		c.Logger().Errorf("Error on DeleteActivityControllerImpl.Delete: %s", err.Error())
		return err
	}
	empty := struct {
		ID int `json:"id,omitempty"`
	}{}
	response := model.BaseResponse{
		Status:  "Success",
		Message: "Success",
		Data:    empty,
	}
	return c.JSON(http.StatusOK, response)
}

func newDeleteActivityController(logger genlog.Logger) *DeleteActivityControllerImpl {
	return &DeleteActivityControllerImpl{
		Service: *service.LoadActivityService(logger),
	}
}

type UpdateActivityControllerImpl struct {
	Service  service.ActivityService
	Bind     model.BindFunc
	Validate model.ValideFunc
}

func (ctr *UpdateActivityControllerImpl) Update(c echo.Context) (err error) {
	data, err := ctr.Bind(c)
	if err != nil {
		c.Logger().Errorf("Error on UpdateActivityControllerImpl.Update: %s", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body").SetInternal(err)
	}
	if err := ctr.Validate(c, data); err != nil {
		c.Logger().Errorf("Validation error on UpdateActivityControllerImpl.Update: %s", err.Error())
		return err
	}

	updated, err := ctr.Service.Update(c.Request().Context(), c.Param("id"), data)
	if err != nil {
		c.Logger().Errorf("Error on UpdateActivityControllerImpl.Update: %s", err.Error())
		return err
	}
	response := model.BaseResponse{
		Status:  "Success",
		Message: "Success",
		Data:    updated,
	}
	return c.JSON(http.StatusOK, response)
}

func newUpdateActivityController(logger genlog.Logger) *UpdateActivityControllerImpl {
	return &UpdateActivityControllerImpl{
		Service:  *service.LoadActivityService(logger),
		Bind:     model.BindActivityUpdate,
		Validate: model.ValidateActivityUpdate,
	}
}
