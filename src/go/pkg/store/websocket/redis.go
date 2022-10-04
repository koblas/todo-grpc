package websocket

import (
	"strings"
	"time"

	"github.com/go-redis/redis"
)

type redisStore struct {
	rdb *redis.Client
}

const KEY_PREFIX = "store:websocket-connections:"
const LIFETIME = time.Minute * 15

func NewRedisStore(redisAddr string) *redisStore {
	rdb := redis.NewClient(&redis.Options{
		Addr:        redisAddr,
		Password:    "",                     // no password set
		DB:          0,                      // use default DB
		DialTimeout: time.Millisecond * 200, // either it happens or it doesn't
	})

	return &redisStore{
		rdb: rdb,
	}
}

func (store *redisStore) Create(userId string, connectionId string) error {
	_, err := store.rdb.Set(KEY_PREFIX+connectionId, userId, LIFETIME).Result()

	return err
}

func (store *redisStore) Delete(connectionId string) error {
	_, err := store.rdb.Del(KEY_PREFIX + connectionId).Result()

	return err
}

func (store *redisStore) ForUser(userId string) ([]string, error) {
	conns := []string{}

	for pos, newPos := uint64(0), uint64(1); newPos != 0; pos = newPos {
		var err error
		var res []string

		res, newPos, err = store.rdb.Scan(pos, KEY_PREFIX+"*", 100).Result()
		if err != nil {
			return conns, err
		}

		if len(res) == 0 {
			continue
		}

		values, err := store.rdb.MGet(res...).Result()
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

func (store *redisStore) Heartbeat(connectionId string) error {
	_, err := store.rdb.Expire(KEY_PREFIX+connectionId, LIFETIME).Result()

	return err
}
