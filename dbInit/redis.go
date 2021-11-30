package dbInit

import (
	"bytes"
	"context"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

const (
	RedisDefaultPoolSize     = 120
	RedisDefaultMinIdleConns = 20
	RedisDefaultDialTimeout  = 60 * time.Second
	RedisDefaultReadTimeout  = 60 * time.Second
	RedisDefaultWriteTimeout = 60 * time.Second
	RedisDefaultPoolTimeout  = 60 * time.Second
	RedisDefaultIdleTimeout  = 60 * time.Second
)

type RedisOption struct {
	Host         string
	Port         int
	Db           int
	Password     string
	PoolSize     int
	MinIdleConns int
}

// RedisInit 公用初始化方法
func RedisInit(option RedisOption) (*redis.Client, error) {
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
		DialTimeout:  RedisDefaultDialTimeout,
		ReadTimeout:  RedisDefaultReadTimeout,
		WriteTimeout: RedisDefaultWriteTimeout,
		PoolTimeout:  RedisDefaultPoolTimeout,
		IdleTimeout:  RedisDefaultIdleTimeout,
	})
	var ctx = context.Background()
	err := conn.Ping(ctx).Err()
	return conn, err
}
