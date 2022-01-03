// Package ssmconfig is a utility for loading configuration values from AWS SSM (Parameter
// Store) directly into a struct.
package awsutil

import (
	"context"
	"os"
	"path"
	"reflect"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"

	"github.com/pkg/errors"
)

// Process processes the config with a new default provider.
//
// See Provider.Process() for full documentation.
func LoadSsmConfig(configPath string, c interface{}) error {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error, " + err.Error())
	}

	client := ssm.NewFromConfig(cfg)

	p := Provider{Client: client}

	return p.LoadSsmConfig(configPath, c)
}

// Provider is a ssm configuration provider.
type Provider struct {
	Client *ssm.Client
}

// Process loads config values from smm (parameter store) into c. Encrypted parameters
// will automatically be decrypted. c must be a pointer to a struct.
//
// The `ssm` tag is used to lookup the parameter in Parameter Store. It is joined to the
// provided base path. If the `ssm` tag is missing the struct field will be ignored.
//
// The `default` tag is used to set the default value of a parameter. The default value
// will only be set if Parameter Store returns the parameter as invalid.
//
// The `required` tag is used to mark a parameter as required. If Parameter Store returns
// a required parameter as invalid an error will be returned.
//
// The behavior of using the `default` and `required` tags on the same struct field is
// currently undefined.
func (p *Provider) LoadSsmConfig(configPath string, c interface{}) error {
	v := reflect.ValueOf(c)
	if v.Kind() != reflect.Ptr || v.IsNil() {
		return errors.New("ssmconfig: c must be a pointer to a struct")
	}

	v = reflect.Indirect(reflect.ValueOf(c))
	if v.Kind() != reflect.Struct {
		return errors.New("ssmconfig: c must be a pointer to a struct")
	}

	spec := buildStructSpec(configPath, v.Type())

	params, invalidPrams, err := p.getParameters(spec)
	if err != nil {
		return errors.Wrap(err, "ssmconfig: could not get parameters")
	}

	for i, field := range spec {
		if field.name == "" && field.defaultValue == "" {
			continue
		}

		if _, ok := invalidPrams[field.name]; ok && field.required {
			return errors.Errorf("ssmconfig: %s is required", field.name)
		}

		value, ok := params[field.name]
		if !ok {
			value = field.defaultValue
		}

		if value == "" {
			continue
		}

		err = setValue(v.Field(i), value)
		if err != nil {
			return errors.Wrapf(err, "ssmconfig: error setting field %s", v.Type().Field(i).Name)
		}
	}

	return nil
}

func (p *Provider) getParameters(spec structSpec) (params map[string]string, invalidParams map[string]struct{}, err error) {
	params = map[string]string{}
	// find all of the params that need to be requested
	var names []string
	for i := range spec {
		if spec[i].name == "" {
			continue
		}
		if spec[i].environmentVar != "" {
			if value := os.Getenv(spec[i].environmentVar); value != "" {
				params[spec[i].name] = value
				continue
			}
		}
		names = append(names, spec[i].name)
	}

	if len(names) == 0 {
		return params, invalidParams, nil
	}

	input := &ssm.GetParametersInput{
		Names:          names,
		WithDecryption: true,
	}

	output, err := p.Client.GetParameters(context.TODO(), input)
	if err != nil {
		return nil, nil, err
	}
	if output == nil {
		return nil, nil, nil
	}

	// convert the response to a map for easier use later
	for i := range output.Parameters {
		params[*output.Parameters[i].Name] = *output.Parameters[i].Value
	}

	invalidParams = map[string]struct{}{}
	for i := range output.InvalidParameters {
		invalidParams[output.InvalidParameters[i]] = struct{}{}
	}
	return params, invalidParams, nil
}

func setValue(v reflect.Value, s string) error {
	switch v.Kind() {
	case reflect.String:
		v.SetString(s)

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		i, err := strconv.Atoi(s)
		if err != nil {
			return errors.Errorf("could not decode %q into type %v", s, v.Type().String())
		}
		v.SetInt(int64(i))

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

type structSpec []fieldSpec

type fieldSpec struct {
	name           string
	environmentVar string
	defaultValue   string
	required       bool
}

func buildStructSpec(configPath string, t reflect.Type) (spec structSpec) {
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		ssmTag := field.Tag.Get("ssm")

		if ssmTag != "" {
			ssmTag = path.Join(configPath, ssmTag)
		}

		spec = append(spec, fieldSpec{
			name:           ssmTag,
			environmentVar: field.Tag.Get("environment"),
			defaultValue:   field.Tag.Get("default"),
			required:       field.Tag.Get("required") == "true",
		})
	}
	return spec
}
