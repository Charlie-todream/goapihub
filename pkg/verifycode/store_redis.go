package verifycode

import (
	"goapihub/pkg/app"
	"goapihub/pkg/config"
	"goapihub/pkg/redis"
	"time"
)

type RedisStore struct {
	RedisClient *redis.RedisClient
	KeyPrefix string
}

 // Ser 实现verifycode.Store interface 的set方法
func (s *RedisStore) Set(key string,value string) bool  {

	ExpireTime := time.Minute * time.Duration(config.GetInt64("verifycode.expire_time"))

	// 本地环境方便测试
	if app.IsLocal() {
		ExpireTime = time.Minute * time.Duration(config.GetInt64("verifycode.debug_expire_time"))
	}

	return s.RedisClient.Set(s.KeyPrefix+key,value,ExpireTime)
}

// Get实现 verifycode.Store interface 的get方法
func (s *RedisStore) Get(key string,clear bool) (value string)  {
	key = s.KeyPrefix + key
	val := s.RedisClient.Get(key)
	if clear {
		s.RedisClient.Del(key)
	}
	return val
}
func (s *RedisStore) Verify(key, answer string, clear bool) bool {
	v := s.Get(key, clear)
	return v == answer
}