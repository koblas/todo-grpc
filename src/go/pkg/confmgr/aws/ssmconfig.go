package aws

import (
	"context"
	"regexp"
	"strings"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/koblas/grpc-todo/pkg/confmgr"
	"github.com/koblas/grpc-todo/pkg/util"
)

type SsmGetParamertsAPI interface {
	GetParameters(context.Context, *ssm.GetParametersInput, ...func(*ssm.Options)) (*ssm.GetParametersOutput, error)
}

type Provider struct {
	Path      string
	Client    SsmGetParamertsAPI
	toSnake   bool
	cleanRe   *regexp.Regexp
	seperator string
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

	p.seperator = "/"
	p.cleanRe = regexp.MustCompile(p.seperator + p.seperator + "+")

	return &p
}

func (p *Provider) getName(spec *confmgr.ConfigSpec) []string {
	name, ok := spec.Field.Tag.Lookup("ssm")
	if !ok {
		name = spec.Field.Name
		if p.toSnake {
			name = util.ToSnake(name)
		}
	}
	if spec.Parent == nil {
		return []string{name}
	}

	return append(p.getName(spec.Parent), name)
}

func (p *Provider) Loader(ctx context.Context, conf interface{}, specs []*confmgr.ConfigSpec) ([]*confmgr.ConfigSpec, error) {
	names := []string{}
	reverse := map[string]*confmgr.ConfigSpec{}

	for _, spec := range specs {
		nlist := p.getName(spec)
		if len(p.Path) != 0 {
			nlist = append([]string{p.Path}, nlist...)
		}
		name := p.cleanRe.ReplaceAllString(strings.Join(nlist, p.seperator), p.seperator)

		reverse[name] = spec

		names = append(names, name)
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
