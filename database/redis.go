package database

import (
	"context"
	"goshaka/configs"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var Rdb *redis.Client

// Connect app to the Redis server
func RedisConnect() error {
	redisDb, _ := strconv.Atoi(configs.GetEnv("REDIS_DB"))
	Rdb = redis.NewClient(&redis.Options{
		Addr:     configs.GetEnv("REDIS_HOST"),
		Password: configs.GetEnv("REDIS_PASSWORD"),
		DB:       redisDb,
	})
	err := Rdb.Set(ctx, "test", "test_goshaja", 0).Err()
	if err != nil {
		return err
	}

	return nil
}

// Get a value from Redis
//
//	params	key	string	key to be fetched
//	return string
func RedisGet(key string) string {
	val, err := Rdb.Get(ctx, key).Result()
	if err != nil {
		return ""
	}

	return val
}

// Set a value to Redis
//
//	params	key	string	key to be implemented
//	params	val	string	value to be implemented
//	params	sec	int	duration in seconds
//	return bool
func RedisSet(key, val string, sec int) bool {
	err := Rdb.Set(ctx, key, val, time.Duration(sec)*time.Second)
	return err == nil
}

// Get a value from Redis, otherwise save new value
//
//	params	key	string	key to be fetched/saved
//	params	val	string	value to be fetched/saved
//	params	sec	int	duration in seconds
//	return string
func RedisGetOrSet(key, val string, sec int) string {
	var v string = RedisGet(key)

	if v == "" {
		RedisSet(key, val, sec)
		v = val
	}
	return v
}
