package auth

import (
	"time"

	"golang.org/x/net/context"
)

type attemptsStub struct{}

var _ AttemptService = (*attemptsStub)(nil)

func NewAttemptsStub() *attemptsStub {
	return &attemptsStub{}
}

func (svc *attemptsStub) GetTries(ctx context.Context, group string, key string) (int64, error) {
	return 0, nil
}

func (svc *attemptsStub) Incr(ctx context.Context, group string, key string, timeout time.Duration) error {
	return nil
}

func (svc *attemptsStub) Reset(ctx context.Context, group string, key string) error {
	return nil
}
