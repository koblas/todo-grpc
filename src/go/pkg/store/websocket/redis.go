package websocket

import (
	"context"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/koblas/grpc-todo/pkg/redisutil"
)

type redisStore struct {
	rdb *redis.Client
}

const KEY_PREFIX = "store:websocket-connections:"
const LIFETIME = time.Minute * 15

func NewRedisStore(redisAddr string) *redisStore {
	return &redisStore{
		rdb: redisutil.NewClient(redisAddr),
	}
}

func (store *redisStore) Create(ctx context.Context, userId string, connectionId string) error {
	_, err := store.rdb.Set(ctx, KEY_PREFIX+connectionId, userId, LIFETIME).Result()

	return err
}

func (store *redisStore) Delete(ctx context.Context, connectionId string) error {
	_, err := store.rdb.Del(ctx, KEY_PREFIX+connectionId).Result()

	return err
}

func (store *redisStore) ForUser(ctx context.Context, userId string) ([]string, error) {
	conns := []string{}

	for pos, newPos := uint64(0), uint64(1); newPos != 0; pos = newPos {
		var err error
		var res []string

		res, newPos, err = store.rdb.Scan(ctx, pos, KEY_PREFIX+"*", 100).Result()
		if err != nil {
			return conns, err
		}

		if len(res) == 0 {
			continue
		}

		values, err := store.rdb.MGet(ctx, res...).Result()
		if err != nil {
			return conns, err
		}

		for idx, uId := range values {
			if sval, ok := uId.(string); ok && sval == userId {
				conns = append(conns, strings.Replace(res[idx], KEY_PREFIX, "", 1))
			}
		}
	}

	return conns, nil
}

func (store *redisStore) Heartbeat(ctx context.Context, connectionId string) error {
	_, err := store.rdb.Expire(ctx, KEY_PREFIX+connectionId, LIFETIME).Result()

	return err
}
