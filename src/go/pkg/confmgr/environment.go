// Package ssmconfig is a utility for loading configuration values from AWS SSM (Parameter
// Store) directly into a struct.
package confmgr

import (
	"os"
)

type envLoader struct{}

func NewLoaderEnvironment() envLoader {
	return envLoader{}
}

func (envLoader) Loader(conf interface{}, specs []*ConfigSpec) ([]*ConfigSpec, error) {
	reducedSpec := []*ConfigSpec{}

	for _, spec := range specs {
		tagValue, ok := spec.Field.Tag.Lookup("environment")
		if !ok {
			reducedSpec = append(reducedSpec, spec)
			continue
		}
		envValue, ok := os.LookupEnv(tagValue)
		if !ok {
			reducedSpec = append(reducedSpec, spec)
			continue
		}

		if err := SetValueString(spec.Value, envValue); err != nil {
			return reducedSpec, err
		}
	}

	return reducedSpec, nil
}
