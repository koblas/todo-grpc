package confmgr

import (
	"context"
	"reflect"
	"strconv"

	"github.com/pkg/errors"
)

type Loadable interface {
	Loader(context.Context, interface{}, []*ConfigSpec) ([]*ConfigSpec, error)
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
		loaders: append(loaders, NewLoaderValidate()),
	}
}

// Parse is a convience function to combine the NewLoader().Parse() calls
func Parse(conf any, loaders ...Loadable) error {
	return ParseWithContext(context.TODO(), conf, loaders...)
}

func ParseWithContext(ctx context.Context, conf any, loaders ...Loadable) error {
	mgr := NewLoader(loaders...)

	return mgr.Parse(ctx, conf)
}

func (mgr *ConfigLoader) Parse(ctx context.Context, conf any) error {
	spec, err := walkInput(conf)

	if err != nil {
		return err
	}

	for _, loader := range mgr.loaders {
		spec, err = loader.Loader(ctx, conf, spec)
		if err != nil {
			return err
		}
	}

	return nil
}

func walkInput(input any) ([]*ConfigSpec, error) {
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

		// fmt.Println("HERE 0 ", field.Name)
		// fmt.Println("HERE k  ", value.Kind())
		// fmt.Println("HERE t  ", value.Type())
		// fmt.Println("HERE ks ", value.Type().Kind().String())

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

	case reflect.Ptr:
		pval := reflect.New(v.Type().Elem())
		v.Set(pval)
		SetValueString(pval.Elem(), s)

	default:
		return errors.Errorf("could not decode %q into type %v", s, v.Type().String())
	}

	return nil
}
