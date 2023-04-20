package activity

import (
	"github.com/labstack/echo/v4"
)

type ActivityController struct {
	echo *echo.Echo
	*GetAllActivityControllerImpl
	*GetActivityControllerImpl
	*CreateActivityControllerImpl
	*DeleteActivityControllerImpl
	*UpdateActivityControllerImpl
}

func NewActivityController(echo *echo.Echo) *ActivityController {
	return &ActivityController{
		echo:                         echo,
		GetAllActivityControllerImpl: newGetAllActivityController(echo.Logger),
		GetActivityControllerImpl:    newGetActivityController(echo.Logger),
		CreateActivityControllerImpl: newCreateActivityController(echo.Logger),
		DeleteActivityControllerImpl: newDeleteActivityController(echo.Logger),
		UpdateActivityControllerImpl: newUpdateActivityController(echo.Logger),
	}
}

func (ctr *ActivityController) Start() {
	r := ctr.echo.Group("/")
	r.GET("activity-groups", ctr.GetAll)
	r.GET("activity-groups/:id", ctr.Get)
	r.POST("activity-groups", ctr.Create)
	r.DELETE("activity-groups/:id", ctr.Delete)
	r.PATCH("activity-groups/:id", ctr.Update)
}
