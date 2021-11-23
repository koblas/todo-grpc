package auth

import (
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
)

type attemptsService struct {
	prefix string
	rdb    *redis.Client
}

type AttemptService interface {
	GetTries(ctx context.Context, group string, key string) (int64, error)
	Incr(ctx context.Context, group string, key string, timeout time.Duration) error
	Reset(ctx context.Context, group string, key string) error
}

func NewAttemptCounter(prefix string, rdb *redis.Client) AttemptService {
	return &attemptsService{
		prefix: prefix,
		rdb:    rdb,
	}
}

func (svc *attemptsService) buildKey(group string, key string) string {
	return svc.prefix + ":" + group + ":" + key
}

func (svc *attemptsService) GetTries(ctx context.Context, group string, key string) (int64, error) {
	attemptsKey := svc.buildKey(group, key)

	attempts, err := svc.rdb.Get(ctx, attemptsKey).Result()
	if err != nil && err != redis.Nil {
		return 0, err
	}

	var count int64
	if err != redis.Nil {
		count, err = strconv.ParseInt(attempts, 10, 32)
		if err != nil {
			return 0, err
		}
	}

	return count, nil
}

func (svc *attemptsService) Incr(ctx context.Context, group string, key string, timeout time.Duration) error {
	attemptsKey := svc.buildKey(group, key)

	if _, err := svc.rdb.Incr(ctx, attemptsKey).Result(); err != nil {
		return fmt.Errorf("redis attempt to set key %w", err)
	} else if _, err := svc.rdb.Expire(ctx, attemptsKey, time.Minute*LOGIN_LOCKOUT_MINUTES).Result(); err != nil {
		return fmt.Errorf("redis attempt to set expires %w", err)
	}

	return nil
}

func (svc *attemptsService) Reset(ctx context.Context, group string, key string) error {
	attemptsKey := svc.buildKey(group, key)

	if _, err := svc.rdb.Del(ctx, attemptsKey).Result(); err != nil {
		return fmt.Errorf("redis attempt to del key %w", err)
	}

	return nil
}
