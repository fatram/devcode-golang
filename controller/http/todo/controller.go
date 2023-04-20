package todo

import (
	"net/http"

	"github.com/fatram/devcode-golang/domain/model"
	"github.com/fatram/devcode-golang/pkg/genlog"
	service "github.com/fatram/devcode-golang/service/todo"
	"github.com/labstack/echo/v4"
)

type GetTodoControllerImpl struct {
	Service service.TodoService
}

func (ctr *GetTodoControllerImpl) Get(c echo.Context) (err error) {
	ctx := c.Request().Context()
	id := c.Param("id")
	data, err := ctr.Service.Get(ctx, id)
	if err != nil {
		c.Logger().Errorf("Error on GetTodoControllerImpl.Get: %s", err.Error())
		return err
	}
	response := model.BaseResponse{
		Status:  "Success",
		Message: "Success",
		Data:    data,
	}
	return c.JSON(http.StatusOK, response)
}

func newGetTodoController(logger genlog.Logger) *GetTodoControllerImpl {
	return &GetTodoControllerImpl{
		Service: *service.LoadTodoService(logger),
	}
}

type GetAllTodoControllerImpl struct {
	Service service.TodoService
	Filter  model.ToDoFilter
}

func (ctr *GetAllTodoControllerImpl) GetAll(c echo.Context) (err error) {
	filter := ctr.Filter

	if err := c.Bind(filter); err != nil {
		c.Logger().Errorf("Error on GetAllTodoControllerImpl.GetAll: %s", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body").SetInternal(err)
	}

	data, err := ctr.Service.GetAll(c.Request().Context(), filter)
	if err != nil {
		c.Logger().Errorf("Error on GetAllTodoControllerImpl.GetAllTodo: %s", err.Error())
		return err
	}

	response := model.BaseResponse{
		Status:  "Success",
		Message: "Success",
		Data:    data,
	}
	return c.JSON(http.StatusOK, response)
}

func newGetAllTodoController(logger genlog.Logger) *GetAllTodoControllerImpl {
	return &GetAllTodoControllerImpl{
		Service: *service.LoadTodoService(logger),
		Filter:  *new(model.ToDoFilter),
	}
}

type CreateTodoControllerImpl struct {
	Service  service.TodoService
	Bind     model.BindFunc
	Validate model.ValideFunc
}

func (ctr *CreateTodoControllerImpl) Create(c echo.Context) (err error) {
	data, err := ctr.Bind(c)
	if err != nil {
		c.Logger().Errorf("Error on CreateTodoControllerImpl.Create: %s", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body").SetInternal(err)
	}
	if err := ctr.Validate(c, data); err != nil {
		c.Logger().Errorf("Validation error on CreateTodoControllerImpl.Create: %s", err.Error())
		return err
	}

	created, err := ctr.Service.Create(c.Request().Context(), data)
	if err != nil {
		c.Logger().Errorf("Error on CreateControllerImpl.Create: %s", err.Error())
		return err
	}
	response := model.BaseResponse{
		Status:  "Success",
		Message: "Success",
		Data:    created,
	}
	return c.JSON(http.StatusCreated, response)
}

func newCreateTodoController(logger genlog.Logger) *CreateTodoControllerImpl {
	return &CreateTodoControllerImpl{
		Service:  *service.LoadTodoService(logger),
		Bind:     model.BindTodoCreate,
		Validate: model.ValidateTodoCreate,
	}
}

type DeleteTodoControllerImpl struct {
	Service service.TodoService
}

func (ctr *DeleteTodoControllerImpl) Delete(c echo.Context) (err error) {
	ctx := c.Request().Context()
	id := c.Param("id")
	err = ctr.Service.Delete(ctx, id)
	if err != nil {
		c.Logger().Errorf("Error on DeleteTodoControllerImpl.Delete: %s", err.Error())
		return err
	}
	response := model.BaseResponse{
		Status:  "Success",
		Message: "Success",
	}
	return c.JSON(http.StatusOK, response)
}

func newDeleteTodoController(logger genlog.Logger) *DeleteTodoControllerImpl {
	return &DeleteTodoControllerImpl{
		Service: *service.LoadTodoService(logger),
	}
}

type UpdateTodoControllerImpl struct {
	Service  service.TodoService
	Bind     model.BindFunc
	Validate model.ValideFunc
}

func (ctr *UpdateTodoControllerImpl) Update(c echo.Context) (err error) {
	data, err := ctr.Bind(c)
	if err != nil {
		c.Logger().Errorf("Error on UpdateTodoControllerImpl.Update: %s", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body").SetInternal(err)
	}
	if err := ctr.Validate(c, data); err != nil {
		c.Logger().Errorf("Validation error on UpdateTodoControllerImpl.Update: %s", err.Error())
		return err
	}

	updated, err := ctr.Service.Update(c.Request().Context(), c.Param("id"), data)
	if err != nil {
		c.Logger().Errorf("Error on UpdateTodoControllerImpl.Update: %s", err.Error())
		return err
	}
	response := model.BaseResponse{
		Status:  "Success",
		Message: "Success",
		Data:    updated,
	}
	return c.JSON(http.StatusOK, response)
}

func newUpdateTodoController(logger genlog.Logger) *UpdateTodoControllerImpl {
	return &UpdateTodoControllerImpl{
		Service:  *service.LoadTodoService(logger),
		Bind:     model.BindTodoUpdate,
		Validate: model.ValidateTodoUpdate,
	}
}
