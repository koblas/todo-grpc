package confmgr

import (
	"context"
	"os"
	"strings"

	"github.com/koblas/grpc-todo/pkg/util"
)

type envLoader struct {
	prefix    string
	seperator string
}

func NewLoaderEnvironment(prefix string, seperator string) envLoader {
	return envLoader{prefix, seperator}
}

func (e envLoader) getName(spec *ConfigSpec) []string {
	name, ok := spec.Field.Tag.Lookup("environment")
	if !ok {
		name = strings.ToUpper(util.ToSnake(spec.Field.Name))
	}
	if spec.Parent == nil {
		return []string{name}
	}

	return append(e.getName(spec.Parent), name)
}

func (e envLoader) Loader(_ context.Context, conf interface{}, specs []*ConfigSpec) ([]*ConfigSpec, error) {
	reducedSpec := []*ConfigSpec{}

	for _, spec := range specs {
		names := e.getName(spec)
		if len(e.prefix) != 0 {
			names = append([]string{e.prefix}, names...)
		}
		name := strings.Join(names, e.seperator)

		envValue, ok := os.LookupEnv(name)
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
