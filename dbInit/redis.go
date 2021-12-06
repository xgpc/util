package dbInit

import (
	"bytes"
	"context"
	"dog3pack/gos/env"
	"github.com/go-redis/redis/v8"
	"strconv"
)

// RedisInit 公用初始化方法
func RedisInit(option env.RedisOption) (*redis.Client, error) {
	var addr bytes.Buffer
	addr.WriteString(option.Host)
	addr.WriteString(":")
	addr.WriteString(strconv.Itoa(option.Port))

	conn := redis.NewClient(&redis.Options{
		Addr:         addr.String(),
		Password:     option.Password,
		DB:           option.Db,
		PoolSize:     option.PoolSize,
		MinIdleConns: option.MinIdleConns,
		DialTimeout:  env.RedisDefaultDialTimeout,
		ReadTimeout:  env.RedisDefaultReadTimeout,
		WriteTimeout: env.RedisDefaultWriteTimeout,
		PoolTimeout:  env.RedisDefaultPoolTimeout,
		IdleTimeout:  env.RedisDefaultIdleTimeout,
	})
	var ctx = context.Background()
	err := conn.Ping(ctx).Err()
	return conn, err
}
