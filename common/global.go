package common

import (
	"github.com/minio/minio-go/v7"
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
	Log         *zap.SugaredLogger // 日志工具
	DB          *gorm.DB           // 数据库连接
	Cache       *redis.Client      // Redis 连接
	MinioClient *minio.Client      // Minio 客户端
)

// uint 类型
var (
	Male           uint = 1 // 性别：男
	Female         uint = 2 // 性别：女
	BrokenStatus   uint = 0 // 用户状态：禁用
	NormalStatus   uint = 1 // 用户状态：正常
	UnActiveStatus uint = 2 // 用户状态：未激活
	False          uint = 0 // 否
	True           uint = 1 // 是
)

// 密码复杂度层级
const (
	InvalidLevel = iota // 非法密码，密码小于 8 位或者大于 20 位，或者不在指定字符串中
	WeakLevel           // 弱密码，只包含数字或小写字母或大写字母中的一种
	MediumLevel         // 中等密码，只包含数字，大小写字母
	StrongLevel         // 强密码，8-12 位以上，包含数字，大小写字母，特殊符号
	FinalLevel          // 究极密码，12 位以上，包含数字，大小写字母，特殊符号
)

// 密码复杂度提示信息
var PasswordLevelMessage = map[int]string{
	InvalidLevel: "非法密码，密码不允许小于 8 位或者大于 20 位，或包含非法字符串",
	WeakLevel:    "密码至少包含数字，大小写字母中的一种，且大于 8 位，小于 20 位",
	MediumLevel:  "密码至少包含数字，大小写字母中两种以上，且大于 8 位，小于 20 位",
	StrongLevel:  "密码至少包含数字，大小写字母，，特殊符号中三种以上，且大于 8 位，小于 20 位",
	FinalLevel:   "密码至少包含数字，大小写字母，，特殊符号中三种以上，且大于 12 位，小于 20 位",
}
