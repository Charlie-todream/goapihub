package captcha

import (
	"errors"
	"goapihub/pkg/app"
	"goapihub/pkg/config"
	"goapihub/pkg/redis"
	"time"
)

type RedisStore struct {
	RedisClient *redis.RedisClient
	KeyPrefix string
}

// Set 实现base64Captcha.store intrerface 的set方法
func (s *RedisStore) Set(key string,value string) error  {

	ExpireTme := time.Minute * time.Duration(config.GetInt64("captcha.expire_time"))

	// 方便本地开发调试
	if app.IsLocal() {
		ExpireTme = time.Minute * time.Duration(config.GetInt64("captcha.debug_expire_time"))
	}

	if ok := s.RedisClient.Set(s.KeyPrefix + key, value,ExpireTme);!ok {
		return errors.New("无法存储图片验证码答案")
	}
	return nil
}

func (s *RedisStore) Get(key string,clear bool) string  {
	key = s.KeyPrefix + key
	val := s.RedisClient.Get(key)

	if clear {
		s.RedisClient.Del(key)
	}
	return  val
}

// 验证
func (s *RedisStore) Verify(key,answer string,clear bool) bool  {
	v := s.Get(key,clear)
	return v == answer
}