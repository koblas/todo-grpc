// Package ssmconfig is a utility for loading configuration values from AWS SSM (Parameter
// Store) directly into a struct.
package confmgr

import (
	"encoding/json"
	"io"
	"os"
	"strings"

	"github.com/pkg/errors"
)

type jsonLoader struct {
	path   string
	reader io.Reader
}

type jsonMatch map[*ConfigSpec]struct{}

func NewJsonFile(path string) *jsonLoader {
	return &jsonLoader{
		path: path,
	}
}

func NewJsonReader(reader io.Reader) *jsonLoader {
	return &jsonLoader{
		reader: reader,
	}
}

func (ldr *jsonLoader) Loader(conf interface{}, specs []*ConfigSpec) ([]*ConfigSpec, error) {
	if ldr.reader == nil && ldr.path == "" {
		return specs, nil
	}
	reader := ldr.reader
	if ldr.reader == nil {
		var err error
		reader, err = os.Open(ldr.path)
		if os.IsNotExist(err) {
			// Ignore unable to open
			return specs, nil
		}
		if err != nil {
			return specs, err
		}
	}

	decoder := json.NewDecoder(reader)
	decoder.UseNumber()

	t, err := decoder.Token()
	if err == io.EOF {
		return specs, nil
	} else if err != nil {
		return specs, err
	}
	if t != json.Delim('{') {
		return specs, errors.New("json must start with a map")
	}

	matches, err := ldr.process(decoder, nil, specs, false)
	if err != nil {
		return specs, err
	}

	reducedSpec := []*ConfigSpec{}
	for _, item := range specs {
		if _, found := matches[item]; found {
			continue
		}
		reducedSpec = append(reducedSpec, item)
	}

	return reducedSpec, nil
}

func (ldr *jsonLoader) process(decoder *json.Decoder, parent *ConfigSpec, specs []*ConfigSpec, inArray bool) (jsonMatch, error) {
	matches := jsonMatch{}
	canidates := []*ConfigSpec{}

	for _, item := range specs {
		if item.Parent == parent {
			canidates = append(canidates, item)
		}
	}

	var node *ConfigSpec
	isKey := !inArray

	for {
		t, err := decoder.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			return jsonMatch{}, err
		}

		if isKey {
			var key string
			switch v := t.(type) {
			case json.Delim:
				if v == json.Delim('}') {
					break
				}
			case json.Number:
				key = v.String()
			case string:
				key = v
			default:
				return matches, errors.Errorf("unexpected %T as key", t)
			}

			node = nil
			for _, item := range canidates {
				tag := item.Field.Tag.Get("json")
				if tag == "-" {
					continue
				}
				if (tag != "" && tag == key) || (tag == "" && strings.EqualFold(item.Field.Name, key)) {
					node = item
					break
				}
			}
			isKey = false
		} else {
			switch v := t.(type) {
			case json.Delim:
				if t == json.Delim('{') {
					m, err := ldr.process(decoder, node, specs, false)
					if err != nil {
						return matches, err
					}
					for k := range m {
						matches[k] = struct{}{}
					}
				} else if t == json.Delim('[') {
					_, err := ldr.process(decoder, node, specs, true)
					if err != nil {
						return matches, err
					}
				} else if t == json.Delim('}') || t == json.Delim(']') {
					break
				}
			case bool:
				if node != nil {
					val := "false"
					if v {
						val = "true"
					}

					SetValueString(node.Value, val)
				}
			case json.Number:
				if node != nil {
					SetValueString(node.Value, v.String())
				}
			case string:
				if node != nil {
					SetValueString(node.Value, v)
					matches[node] = struct{}{}
				}
			case nil:
				if node != nil {
					matches[node] = struct{}{}
				}
			}
			isKey = !inArray
		}
	}

	return matches, nil
}
