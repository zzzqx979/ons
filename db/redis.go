package db

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

var redisCli *redis.Client

const (
	RedisFieldTid         = "tid"
	RedisFieldEpc         = "epc"
	RedisFieldUris        = "uris"
	RedisFieldModel       = "model"
	RedisFieldFactoryCode = "factory_code"
	RedisFieldProductCode = "product_code"
)

func RedisCon() {
	redisCli = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.addr"),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
	})
}

// redis缓存脚本，保证缓存数据和过期时间之间的一致性
var AddTidRecordScript = redis.NewScript(`
redis.call("HSET", KEYS[1], KEYS[2], ARGV[1], KEYS[3], ARGV[2], KEYS[4], ARGV[3])
redis.call("EXPIRE", KEYS[1], ARGV[4])
return 1`)

var AddEpcRecordScript = redis.NewScript(`
redis.call("HSET", KEYS[1], KEYS[2], ARGV[1], KEYS[3], ARGV[2], KEYS[4], ARGV[3])
redis.call("EXPIRE", KEYS[1], ARGV[4])
return 1`)

var AddCodesRecordScript = redis.NewScript(`
redis.call("SET", KEYS[1], ARGV[1])
redis.call("EXPIRE", KEYS[1], ARGV[2])
return 1`)

func GetRedisCli() *redis.Client {
	return redisCli
}
