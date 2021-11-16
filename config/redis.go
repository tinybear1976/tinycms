package config

import (
	"github.com/spf13/viper"
	"github.com/tinybear1976/redisdb2"
	"github.com/tinybear1976/tinycms/defines"
)

func InitRedis() {
	initAuth_Token()
}

func initAuth_Token() {
	redisdb2.New(
		defines.REDIS_AUTH_TOKEN,
		viper.GetString("redis.audit_auth_token.server"),
		viper.GetString("redis.audit_auth_token.pwd"),
		viper.GetInt("redis.audit_auth_token.db"))
}
