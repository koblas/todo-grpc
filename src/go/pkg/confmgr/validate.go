package confmgr

import (
	"context"

	"github.com/go-playground/validator/v10"
)

type validateLoader struct{}

func NewLoaderValidate() validateLoader {
	return validateLoader{}
}

func (validateLoader) Loader(_ context.Context, conf interface{}, spec []*ConfigSpec) ([]*ConfigSpec, error) {
	validate := validator.New()

	return spec, validate.Struct(conf)
}
