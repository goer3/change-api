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

// uint 类型
var (
    Male           uint = 1 // 性别：男
    Female         uint = 2 // 性别：女
    BrokenStatus   uint = 0 // 用户状态：禁用
    NormalStatus   uint = 1 // 用户状态：正常
    UnActiveStatus uint = 2 // 用户状态：未激活
    LockedStatus   uint = 3 // 用户状态：锁定
    False          uint = 0 // 否
    True           uint = 1 // 是
)
