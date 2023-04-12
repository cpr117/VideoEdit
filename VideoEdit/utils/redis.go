// @User CPR
package utils

import (
	"VideoEdit/config"
	"context"
	"github.com/redis/go-redis/v9"
	_ "github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"log"
	"time"
)

const REDIS_UTIL_ERR_PREFIX = "utils/redis.go ->"

var (
	ctx = context.Background()
	rdb *redis.Client
)

type _redis struct{}

var Redis = new(_redis)

func InitRedis() *redis.Client {
	redisCfg := config.Cfg.Redis
	rdb = redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password,
		DB:       redisCfg.DB,
		//PoolSize: redisCfg.PoolSize,
		//MaxIdleConns: redisCfg.MaxIdleConns,
	})
	// 测试连接状况
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Panic("Redis 连接失败: ", err)
	}
	log.Println("Redis 连接成功 ")
	return rdb
}

// 通用执行指令方法
func (*_redis) Execute(cmd string, args ...any) (any, error) {
	return rdb.Do(ctx, args...).Result()
}

// Keys
// redis 获取根据匹配项获取键名列表
func (*_redis) Keys(pattern string) []string {
	return rdb.Keys(ctx, pattern).Val()
}

// Del
// redis 删除值
func (*_redis) Del(key string) {
	err := rdb.Del(ctx, key).Err()
	if err != nil {
		Logger.Error(REDIS_UTIL_ERR_PREFIX+"Del: ", zap.Error(err))
		panic(err)
	}
}

// Set
// redis 设置值
func (*_redis) Set(key string, value interface{}, expiration time.Duration) {
	err := rdb.Set(ctx, key, value, expiration).Err()
	if err != nil {
		Logger.Error(REDIS_UTIL_ERR_PREFIX+"Set: ", zap.Error(err))
		panic(err)
	}
}

func (*_redis) Incr(key string) int64 {
	ret := rdb.Incr(ctx, key)
	if err := ret.Err(); err != nil {
		Logger.Error(REDIS_UTIL_ERR_PREFIX+"Incr: ", zap.Error(err))
		panic(err)
	}
	return ret.Val()
}

// GetVal
// redis string类型
func (*_redis) GetVal(key string) string {
	ret := rdb.Get(ctx, key)
	if err := ret.Err(); err != nil {
		Logger.Error(REDIS_UTIL_ERR_PREFIX+"GetVal: ", zap.Error(err))
	}
	return ret.Val()
}

// GetInt
// redis获取数字
func (*_redis) GetInt(key string) int {
	val, err := rdb.Get(ctx, key).Int()
	if err != nil {
		Logger.Error(REDIS_UTIL_ERR_PREFIX+"GetInt: ", zap.Error(err))
		panic(err)
	}
	return val
}

// 从redis中取值，不存在会有 redis:nil 的错误
func (*_redis) GetResult(key string) (string, error) {
	return rdb.Get(ctx, key).Result()
}

// 获取[哈希表(Hash)]中 指定字段的值
func (*_redis) HGet(key, filed string) int {
	val, _ := rdb.HGet(ctx, key, filed).Int()
	return val
}

// 获取[哈希表(Hash)]中 指定字段的值
func (*_redis) HGetCode(key, filed string) string {
	val, _ := rdb.HGet(ctx, key, filed).Result()
	return val
}

func (*_redis) HSet(key, field string, value interface{}, expiration time.Duration) {
	err := rdb.HSet(ctx, key, field, value).Err()
	if err != nil {
		Logger.Error(REDIS_UTIL_ERR_PREFIX+"HSet: ", zap.Error(err))
		panic(err)
	}
	if expiration > 0 {
		err = rdb.Expire(ctx, key, expiration).Err()
		if err != nil {
			Logger.Error(REDIS_UTIL_ERR_PREFIX+"HSet: ", zap.Error(err))
			panic(err)
		}
	}
}

// 获取[哈希表(Hash)]中 所有的字段和值
func (*_redis) HGetAll(key string) map[string]string {
	return rdb.HGetAll(ctx, key).Val()
}
