package common

import (
    "github.com/redis/go-redis/v9"
    "go.uber.org/zap"
    "gorm.io/gorm"
)

// 时间格式化
var (
    MsecTimeFormat = "2006-01-02 15:04:05.000"
    SecTimeFormat  = "2006-01-02 15:04:05"
    DateTimeFormat = "2006-01-02"
)

// 全局工具
var (
    Log   *zap.SugaredLogger // 日志工具
    DB    *gorm.DB           // 数据库连接
    Cache *redis.Client      // Redis 连接
)
