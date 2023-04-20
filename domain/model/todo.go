package model

type Todo struct {
	ID         int    `json:"id"`
	ActivityID int    `json:"activity_group_id"`
	Title      string `json:"title"`
	IsActive   bool   `json:"is_active"`
	Priority   string `json:"priority"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

type TodoCreate struct {
	ActivityID int    `json:"activity_group_id"`
	Title      string `json:"title"`
	Priority   string `json:"priority"`
}

type TodoUpdate struct {
	IsActive bool   `json:"is_active"`
	Title    string `json:"title"`
	Priority string `json:"priority"`
}

type ToDoFilter struct {
	ActivityID *int `query:"activity_group_id"`
}

func (filter *ToDoFilter) SetDefaultForEmpty() {

}

func ValidateTodoCreate(validator Validator, data interface{}) error {
	todo := data.(*TodoCreate)
	return validator.Validate(todo)
}

func BindTodoCreate(binder Binder) (data interface{}, err error) {
	todo := new(TodoCreate)
	if err = binder.Bind(todo); err != nil {
		return nil, err
	}
	return todo, nil
}

func ValidateTodoUpdate(validator Validator, data interface{}) error {
	todo := data.(*TodoUpdate)
	return validator.Validate(todo)
}

func BindTodoUpdate(binder Binder) (data interface{}, err error) {
	todo := new(TodoUpdate)
	if err = binder.Bind(todo); err != nil {
		return nil, err
	}
	return todo, nil
}
