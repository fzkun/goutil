package magicapi

import (
	"github.com/patrickmn/go-cache"
	"go.uber.org/zap"
	"time"
)

type Context struct {
	Config  Config
	cache   *cache.Cache //缓存模块
	Logger  *zap.Logger
	ApiHook ApiHook
}

func NewContext(config Config, logger *zap.Logger, apiHook ApiHook) *Context {
	return &Context{
		Config:  config,
		Logger:  logger,
		cache:   cache.New(5*time.Minute, 10*time.Minute),
		ApiHook: apiHook,
	}
}
