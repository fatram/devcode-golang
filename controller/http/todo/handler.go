package todo

import "github.com/labstack/echo/v4"

// Listtodo godoc
// @Summary			List todo
// @Description		Menampilkan daftar todo
// @Tags         	todo
// @Accept       	json
// @Produce      	json
// @Success      	200  {object}   model.BaseResponse{data=[]model.Todo}
// @Failure      	400,401,500  {object}  	pkg.Error
// @Router       	/todo-groups [get]
func (ctr TodoController) GetAll(c echo.Context) error {
	return ctr.GetAllTodoControllerImpl.GetAll(c)
}

// Create todo godoc
// @Summary			Register todo
// @Description		Membuat todo
// @Tags         	todo
// @Accept       	json
// @Accept       	x-www-form-urlencoded
// @Produce      	json
// @Param			body			body	model.TodoCreate	true	"body"
// @Success      	201  {object}   model.BaseResponse{data=model.Todo}
// @Failure      	400,500  {object}  	pkg.Error
// @Router       	/todo-groups [post]
func (ctr TodoController) Create(c echo.Context) error {
	return ctr.CreateTodoControllerImpl.Create(c)
}

// Gettodo godoc
// @Summary			Get one todo
// @Description		Menampilkan satu todo
// @Tags         	todo
// @Accept       	json
// @Produce      	json
// @param 	   		id		path 	string 	true 	"id todo"
// @Success      	200  {object}   model.BaseResponse{data=model.Todo}
// @Failure      	400,401,500  {object}  	pkg.Error
// @Router       	/todo-groups/{id} [get]
func (ctr TodoController) Get(c echo.Context) error {
	return ctr.GetTodoControllerImpl.Get(c)
}

// Deletetodo godoc
// @Summary			Delete one todo
// @Description		Menampilkan satu todo
// @Tags         	todo
// @Accept       	json
// @Produce      	json
// @param 	   		id		path 	string 	true 	"id todo"
// @Success      	200  {object}   model.BaseResponse
// @Failure      	400,401,500  {object}  	pkg.Error
// @Router       	/todo-groups/{id} [delete]
func (ctr TodoController) Delete(c echo.Context) error {
	return ctr.DeleteTodoControllerImpl.Delete(c)
}

// Updatetodo godoc
// @Summary			Update one todo
// @Description		Menampilkan satu todo
// @Tags         	todo
// @Accept       	json
// @Produce      	json
// @param 	   		id		path 	string 	true 	"id todo"
// @Param			body			body	model.TodoUpdate	true	"body"
// @Success      	200  {object}   model.BaseResponse{data=model.Todo}
// @Failure      	400,401,500  {object}  	pkg.Error
// @Router       	/todo-groups/{id} [patch]
func (ctr TodoController) Update(c echo.Context) error {
	return ctr.UpdateTodoControllerImpl.Update(c)
}
