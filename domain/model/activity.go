package model

type Activity struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type ActivityCreate struct {
	Title string `json:"title"`
	Email string `json:"email"`
}

type ActivityUpdate struct {
	Title string `json:"title"`
}

func ValidateActivityCreate(validator Validator, data interface{}) error {
	activity := data.(*ActivityCreate)
	return validator.Validate(activity)
}

func BindActivityCreate(binder Binder) (data interface{}, err error) {
	activity := new(ActivityCreate)
	if err = binder.Bind(activity); err != nil {
		return nil, err
	}
	return activity, nil
}

func ValidateActivityUpdate(validator Validator, data interface{}) error {
	activity := data.(*ActivityUpdate)
	return validator.Validate(activity)
}

func BindActivityUpdate(binder Binder) (data interface{}, err error) {
	activity := new(ActivityUpdate)
	if err = binder.Bind(activity); err != nil {
		return nil, err
	}
	return activity, nil
}
