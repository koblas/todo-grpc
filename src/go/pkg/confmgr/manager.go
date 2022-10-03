package confmgr

import (
	"reflect"
	"strconv"

	"github.com/pkg/errors"
)

type Loadable interface {
	Loader(interface{}, []*ConfigSpec) ([]*ConfigSpec, error)
}

type ConfigSpec struct {
	Field  reflect.StructField
	Value  reflect.Value
	Parent *ConfigSpec
}

// Accumulate
// Execute
// Post

type ConfigLoader struct {
	loaders []Loadable
}

func NewLoader(loaders ...Loadable) *ConfigLoader {
	return &ConfigLoader{
		loaders: append(loaders, NewLoaderEnvironment(), NewLoaderDefault(), NewLoaderValidate()),
	}
}

// Parse is a convience function to combine the NewLoader().Parse() calls
func Parse(conf interface{}, loaders ...Loadable) error {
	mgr := &ConfigLoader{
		loaders: append(loaders, NewLoaderEnvironment(), NewLoaderDefault(), NewLoaderValidate()),
	}

	return mgr.Parse(conf)
}

func (mgr *ConfigLoader) Parse(conf interface{}) error {
	spec, err := walkInput(conf)

	if err != nil {
		return err
	}

	for _, loader := range mgr.loaders {
		spec, err = loader.Loader(conf, spec)
		if err != nil {
			return err
		}
	}

	return nil
}

func walkInput(input interface{}) ([]*ConfigSpec, error) {
	v := reflect.ValueOf(input)

	if v.Kind() != reflect.Ptr || v.IsNil() {
		return []*ConfigSpec{}, errors.New("confmgr: input must be a pointer and not nil")
	}

	v = reflect.Indirect(reflect.ValueOf(input))
	if v.Kind() != reflect.Struct {
		return []*ConfigSpec{}, errors.New("confmgr: input must be a pointer to a struct " + v.Kind().String())
	}

	spec := walkStruct(v, nil)

	return spec, nil
}

func walkStruct(v reflect.Value, parent *ConfigSpec) []*ConfigSpec {
	spec := []*ConfigSpec{}
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		item := ConfigSpec{
			Field:  field,
			Value:  value,
			Parent: parent,
		}

		if value.Kind() == reflect.Struct {
			children := walkStruct(value, &item)
			spec = append(spec, children...)
		} else if value.Kind() == reflect.Ptr && reflect.Indirect(value).Kind() == reflect.Struct {
			children := walkStruct(reflect.Indirect(value), &item)
			spec = append(spec, children...)
		} else {
			spec = append(spec, &item)
		}
	}

	return spec
}

func SetValueString(v reflect.Value, s string) error {
	switch v.Kind() {
	case reflect.String:
		v.SetString(s)

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		i, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return errors.Errorf("could not decode %q into type %v", s, v.Type().String())
		}
		v.SetInt(i)

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		i, err := strconv.ParseUint(s, 10, 64)
		if err != nil {
			return errors.Errorf("could not decode %q into type %v", s, v.Type().String())
		}
		v.SetUint(i)

	case reflect.Float32:
		f, err := strconv.ParseFloat(s, 32)
		if err != nil {
			return errors.Errorf("could not decode %q into type %v: %v", s, v.Type().String(), err)
		}
		v.SetFloat(f)

	case reflect.Float64:
		f, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return errors.Errorf("could not decode %q into type %v: %v", s, v.Type().String(), err)
		}
		v.SetFloat(f)

	case reflect.Bool:
		if s != "true" && s != "false" {
			return errors.Errorf("could not decode %q into type %v", s, v.Type().String())
		}
		v.SetBool(s == "true")

	default:
		return errors.Errorf("could not decode %q into type %v", s, v.Type().String())
	}

	return nil
}
