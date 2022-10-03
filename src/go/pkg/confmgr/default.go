// Package ssmconfig is a utility for loading configuration values from AWS SSM (Parameter
// Store) directly into a struct.
package confmgr

type defLoader struct{}

func NewLoaderDefault() defLoader {
	return defLoader{}
}

func (defLoader) Loader(conf interface{}, specs []*ConfigSpec) ([]*ConfigSpec, error) {
	reducedSpec := []*ConfigSpec{}

	for _, spec := range specs {
		tagValue, ok := spec.Field.Tag.Lookup("default")
		if !ok {
			reducedSpec = append(reducedSpec, spec)
			continue
		}

		if err := SetValueString(spec.Value, tagValue); err != nil {
			return reducedSpec, err
		}
	}

	return reducedSpec, nil
}
