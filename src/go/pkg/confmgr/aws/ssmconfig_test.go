package aws_test

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/koblas/grpc-todo/pkg/confmgr"
	ssmconfig "github.com/koblas/grpc-todo/pkg/confmgr/aws"
	"github.com/stretchr/testify/assert"
)

type mockSSMClient struct {
	calledWithInput *ssm.GetParametersInput
	output          *ssm.GetParametersOutput
	err             error
}

func (c *mockSSMClient) GetParameters(ctx context.Context, input *ssm.GetParametersInput, opts ...func(*ssm.Options)) (*ssm.GetParametersOutput, error) {
	c.calledWithInput = input
	if c.err != nil {
		return nil, c.err
	}
	return c.output, nil
}

func TestBaseCase(t *testing.T) {
	var s struct {
		S1      string  `ssm:"/strings/s1"`
		S2      string  `ssm:"/strings/s2" default:"string2"`
		I1      int     `ssm:"/int/i1"`
		I2      int     `ssm:"/int/i2" default:"42"`
		B1      bool    `ssm:"/bool/b1"`
		B2      bool    `ssm:"/bool/b2" default:"false"`
		F321    float32 `ssm:"/float32/f321"`
		F322    float32 `ssm:"/float32/f322" default:"42.42"`
		F641    float64 `ssm:"/float64/f641"`
		F642    float64 `ssm:"/float64/f642" default:"42.42"`
		Invalid string
	}

	mc := &mockSSMClient{
		output: &ssm.GetParametersOutput{
			Parameters: []types.Parameter{
				{
					Name:  aws.String("/base/strings/s1"),
					Value: aws.String("string1"),
				},
				{
					Name:  aws.String("/base/int/i1"),
					Value: aws.String("42"),
				},
				{
					Name:  aws.String("/base/bool/b1"),
					Value: aws.String("true"),
				},
				{
					Name:  aws.String("/base/float32/f321"),
					Value: aws.String("42.42"),
				},
				{
					Name:  aws.String("/base/float64/f641"),
					Value: aws.String("42.42"),
				},
			},
		},
	}

	err := confmgr.NewLoader(
		ssmconfig.NewLoaderSsm(context.TODO(), "/base", ssmconfig.WithClient(mc)),
	).Parse(context.TODO(), &s)

	if err != nil {
		t.Errorf("LoadConfig() unexpected error: %q", err.Error())
	}

	names := make([]string, len(mc.calledWithInput.Names))
	copy(names, mc.calledWithInput.Names)

	expectedNames := []string{
		"/base/strings/s1",
		"/base/strings/s2",
		"/base/int/i1",
		"/base/int/i2",
		"/base/bool/b1",
		"/base/bool/b2",
		"/base/float32/f321",
		"/base/float32/f322",
		"/base/float64/f641",
		"/base/float64/f642",
	}

	if !reflect.DeepEqual(names, expectedNames) {
		t.Errorf("LoadConfig() unexpected input names: have %v, want %v", names, expectedNames)
	}

	assert.Equal(t, "string1", s.S1)
	assert.Equal(t, "string2", s.S2)
	assert.Equal(t, 42, s.I1)
	assert.Equal(t, 42, s.I2)
	assert.Equal(t, true, s.B1)
	assert.Equal(t, false, s.B2)
	assert.EqualValues(t, 42.42, s.F321)
	assert.EqualValues(t, 42.42, s.F322)
	assert.Equal(t, 42.42, s.F641)
	assert.Equal(t, 42.42, s.F642)

	if s.Invalid != "" {
		t.Errorf("LoadConfig() Missing unexpected value: want %q, have %q", "", s.Invalid)
	}
}

func TestProvider_LoadSsmConfig(t *testing.T) {
	for _, tt := range []struct {
		name       string
		configPath string
		c          interface{}
		want       interface{}
		client     *mockSSMClient
		shouldErr  bool
	}{
		{
			name:       "invalid int",
			configPath: "/base/",
			c: &struct {
				I1 int `ssm:"/int/i1" default:"notAnInt"`
			}{},
			client:    &mockSSMClient{},
			shouldErr: true,
		},
		{
			name:       "invalid float32",
			configPath: "/base/",
			c: &struct {
				F32 float32 `ssm:"/float32/f32" default:"notAFloat"`
			}{},
			client:    &mockSSMClient{},
			shouldErr: true,
		},
		{
			name:       "invalid float64",
			configPath: "/base/",
			c: &struct {
				F32 float64 `ssm:"/float64/f64" default:"notAFloat"`
			}{},
			client:    &mockSSMClient{},
			shouldErr: true,
		},
		{
			name:       "invalid bool",
			configPath: "/base/",
			c: &struct {
				B1 bool `ssm:"/bool/b1" default:"notABool"`
			}{},
			client:    &mockSSMClient{},
			shouldErr: true,
		},
		{
			name:       "missing required parameter",
			configPath: "/base",
			c: &struct {
				S1 string `ssm:"/strings/s1" validate:"required"`
			}{},
			client: &mockSSMClient{
				output: &ssm.GetParametersOutput{
					InvalidParameters: []string{"/base/strings/s1"},
				},
			},
			shouldErr: true,
		},
		{
			name:       "unsupported field type",
			configPath: "/base",
			c: &struct {
				M1 map[string]string `ssm:"/map/m1"`
			}{},
			client: &mockSSMClient{
				output: &ssm.GetParametersOutput{
					Parameters: []types.Parameter{{
						Name:  aws.String("/base/map/m1"),
						Value: aws.String("notSupported"),
					}},
				},
			},
			shouldErr: true,
		},
		{
			name:       "blank value from ssm",
			configPath: "/base",
			c: &struct {
				S1 string `ssm:"/strings/s1"`
			}{},
			want: &struct {
				S1 string `ssm:"/strings/s1"`
			}{},
			client: &mockSSMClient{
				output: &ssm.GetParametersOutput{
					Parameters: []types.Parameter{{
						Name:  aws.String("/base/strings/s1"),
						Value: aws.String(""),
					}},
				},
			},
			shouldErr: false,
		},
		{
			name:       "input config not a pointer",
			configPath: "/base/",
			c: struct {
				S1 string `ssm:"/strings/s1"`
			}{},
			client:    &mockSSMClient{},
			shouldErr: true,
		},
		{
			name:       "input config not a struct",
			configPath: "/base/",
			c: &[]struct {
				S1 string `ssm:"/strings/s1"`
			}{},
			client:    &mockSSMClient{},
			shouldErr: true,
		},
		{
			name:       "ssm client error",
			configPath: "/base/",
			c: &struct {
				S1 string `ssm:"/strings/s1" required:"true"`
			}{},
			client: &mockSSMClient{
				err: errors.New("ssm client error"),
			},
			shouldErr: true,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			err := confmgr.NewLoader(
				ssmconfig.NewLoaderSsm(context.TODO(), tt.configPath, ssmconfig.WithClient(tt.client)),
			).Parse(context.TODO(), tt.c)

			if tt.shouldErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)

				assert.Equal(t, tt.want, tt.c)
			}
		})
	}
}
