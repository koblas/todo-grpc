package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/koblas/grpc-todo/pkg/confmgr"
)

type SsmGetParamertsAPI interface {
	GetParameters(context.Context, *ssm.GetParametersInput, ...func(*ssm.Options)) (*ssm.GetParametersOutput, error)
}

type Provider struct {
	Path   string
	Client SsmGetParamertsAPI
}

type Option func(p *Provider)

func WithClient(client SsmGetParamertsAPI) Option {
	return func(p *Provider) {
		p.Client = client
	}
}

func NewLoaderSsm(ctx context.Context, path string, opts ...Option) *Provider {
	p := Provider{Path: path}

	for _, opt := range opts {
		opt(&p)
	}

	if p.Client == nil {
		cfg, err := config.LoadDefaultConfig(ctx)
		if err != nil {
			panic("configuration error, " + err.Error())
		}

		p.Client = ssm.NewFromConfig(cfg)
	}

	return &p
}

func (p *Provider) Loader(ctx context.Context, conf interface{}, specs []*confmgr.ConfigSpec) ([]*confmgr.ConfigSpec, error) {
	names := []string{}
	reverse := map[string]*confmgr.ConfigSpec{}

	for _, spec := range specs {
		tagValue, ok := spec.Field.Tag.Lookup("ssm")
		if !ok {
			tagValue = spec.Field.Name
		}

		qualified := p.Path + tagValue
		reverse[qualified] = spec

		names = append(names, qualified)
	}

	// If we have no work to do...
	if len(names) == 0 {
		return specs, nil
	}

	input := &ssm.GetParametersInput{
		Names:          names,
		WithDecryption: aws.Bool(true),
	}

	output, err := p.Client.GetParameters(ctx, input)
	if err != nil {
		return specs, err
	}
	if output == nil {
		return specs, nil
	}

	didSetValue := map[*confmgr.ConfigSpec]struct{}{}

	for _, outValue := range output.Parameters {
		item, ok := reverse[*outValue.Name]
		if !ok {
			continue
		}

		if err := confmgr.SetValueString(item.Value, *outValue.Value); err != nil {
			return specs, err
		}
		didSetValue[item] = struct{}{}
	}

	reducedSpec := []*confmgr.ConfigSpec{}
	for _, item := range specs {
		if _, found := didSetValue[item]; found {
			continue
		}

		reducedSpec = append(reducedSpec, item)
	}

	return reducedSpec, nil
}
