package todo

import (
	"github.com/labstack/echo/v4"
)

type TodoController struct {
	echo *echo.Echo
	*GetAllTodoControllerImpl
	*GetTodoControllerImpl
	*CreateTodoControllerImpl
	*DeleteTodoControllerImpl
	*UpdateTodoControllerImpl
}

func NewTodoController(echo *echo.Echo) *TodoController {
	return &TodoController{
		echo:                     echo,
		GetAllTodoControllerImpl: newGetAllTodoController(echo.Logger),
		GetTodoControllerImpl:    newGetTodoController(echo.Logger),
		CreateTodoControllerImpl: newCreateTodoController(echo.Logger),
		DeleteTodoControllerImpl: newDeleteTodoController(echo.Logger),
		UpdateTodoControllerImpl: newUpdateTodoController(echo.Logger),
	}
}

func (ctr *TodoController) Start() {
	r := ctr.echo.Group("/")
	r.GET("todo-items", ctr.GetAll)
	r.GET("todo-items/:id", ctr.Get)
	r.POST("todo-items", ctr.Create)
	r.DELETE("todo-items/:id", ctr.Delete)
	r.PATCH("todo-items/:id", ctr.Update)
}
