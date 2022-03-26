package middleware

import (
	"errors"
	"github.com/go-playground/validator/v10"
	postmodel "strider-backend-test.com/api/routes/post/model"
)

type Validator struct {
	validate *validator.Validate
}

func NewValidator() *Validator {
	return &Validator{validate: validator.New()}
}

func (v *Validator) Validate(i interface{}) error {

	switch t := i.(type) {
	case *postmodel.ListRequest:
		if len(t.User) > 0 && len(t.Follower) > 0 {
			return errors.New("cannot have user and follower parameters at the same request")
		}
	}

	return v.validate.Struct(i)
}
