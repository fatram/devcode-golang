package activity

import "github.com/labstack/echo/v4"

// Listactivity godoc
// @Summary			List activity
// @Description		Menampilkan daftar activity
// @Tags         	activity
// @Accept       	json
// @Produce      	json
// @Success      	200  {object}   model.BaseResponse{data=[]model.Activity}
// @Failure      	400,401,500  {object}  	pkg.Error
// @Router       	/activity-groups [get]
func (ctr ActivityController) GetAll(c echo.Context) error {
	return ctr.GetAllActivityControllerImpl.GetAll(c)
}

// Create activity godoc
// @Summary			Register activity
// @Description		Membuat activity
// @Tags         	activity
// @Accept       	json
// @Accept       	x-www-form-urlencoded
// @Produce      	json
// @Param			body			body	model.ActivityCreate	true	"body"
// @Success      	201  {object}   model.BaseResponse{data=model.Activity}
// @Failure      	400,500  {object}  	pkg.Error
// @Router       	/activity-groups [post]
func (ctr ActivityController) Create(c echo.Context) error {
	return ctr.CreateActivityControllerImpl.Create(c)
}

// Getactivity godoc
// @Summary			Get one activity
// @Description		Menampilkan satu activity
// @Tags         	activity
// @Accept       	json
// @Produce      	json
// @param 	   		id		path 	string 	true 	"id activity"
// @Success      	200  {object}   model.BaseResponse{data=model.Activity}
// @Failure      	400,401,500  {object}  	pkg.Error
// @Router       	/activity-groups/{id} [get]
func (ctr ActivityController) Get(c echo.Context) error {
	return ctr.GetActivityControllerImpl.Get(c)
}

// Deleteactivity godoc
// @Summary			Delete one activity
// @Description		Menampilkan satu activity
// @Tags         	activity
// @Accept       	json
// @Produce      	json
// @param 	   		id		path 	string 	true 	"id activity"
// @Success      	200  {object}   model.BaseResponse
// @Failure      	400,401,500  {object}  	pkg.Error
// @Router       	/activity-groups/{id} [delete]
func (ctr ActivityController) Delete(c echo.Context) error {
	return ctr.DeleteActivityControllerImpl.Delete(c)
}

// Updateactivity godoc
// @Summary			Update one activity
// @Description		Menampilkan satu activity
// @Tags         	activity
// @Accept       	json
// @Produce      	json
// @param 	   		id		path 	string 	true 	"id activity"
// @Param			body			body	model.ActivityUpdate	true	"body"
// @Success      	200  {object}   model.BaseResponse{data=model.Activity}
// @Failure      	400,401,500  {object}  	pkg.Error
// @Router       	/activity-groups/{id} [patch]
func (ctr ActivityController) Update(c echo.Context) error {
	return ctr.UpdateActivityControllerImpl.Update(c)
}
