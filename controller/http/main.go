package http

import (
	"fmt"
	"net/http"

	"github.com/fatram/devcode-golang/controller/http/activity"
	"github.com/fatram/devcode-golang/controller/http/todo"
	_ "github.com/fatram/devcode-golang/docs"
	"github.com/fatram/devcode-golang/domain/model"
	"github.com/fatram/devcode-golang/internal/pkg"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4" // we use echo version 4 here
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type HttpController interface {
	Start(host string, port int)
}

type httpCtr struct {
	echo *echo.Echo
}

func NewHttpController() HttpController {
	e := echo.New()
	e.HTTPErrorHandler = customHTTPErrorHandler
	e.Validator = &pkg.CustomValidator{Validator: validator.New()}

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Logger.SetLevel(log.DEBUG)
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Logger.SetHeader(pkg.DefaultLogHeader)

	return &httpCtr{e}
}

func (ctr *httpCtr) Start(host string, port int) {
	activity.NewActivityController(ctr.echo).Start()
	todo.NewTodoController(ctr.echo).Start()
	ctr.echo.Logger.Fatal(ctr.echo.Start(fmt.Sprintf("%s:%d", host, port)))
}

func customHTTPErrorHandler(err error, c echo.Context) {
	response := model.BaseResponse{}
	var msg interface{}
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		msg = he.Message
		response.Status = http.StatusText(code)
	}
	if message, ok := msg.(string); ok {
		response.Message = message
	}
	c.Logger().Error(err)
	if err := c.JSON(code, response); err != nil {
		c.Logger().Error(err)
	}
}
