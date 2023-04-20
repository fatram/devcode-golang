package model

type BaseResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type Binder interface {
	Bind(i interface{}) error
}

type Validator interface {
	Validate(data interface{}) error
}

type BindFunc func(binder Binder) (data interface{}, err error)

type ValideFunc func(validator Validator, data interface{}) error
