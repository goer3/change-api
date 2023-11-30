package common

import (
	"embed"

	"go.uber.org/zap/zapcore"
)

// 配置打包
var FS embed.FS

// 配置引用
var Config Configuration

// 配置结构体
type Configuration struct {
	System  SystemConfiguration  `mapstructure:"system" json:"system"`
	Log     LogConfiguration     `mapstructure:"log" json:"log"`
	MySQL   MySQLConfiguration   `mapstructure:"mysql" json:"mysql"`
	Redis   RedisConfiguration   `mapstructure:"redis" json:"redis"`
	JWT     JWTConfiguration     `mapstructure:"jwt" json:"jwt"`
	Login   LoginConfiguration   `mapstructure:"login" json:"login"`
	Minio   MinioConfiguration   `mapstructure:"minio" json:"minio"`
	OTPAuth OTPAuthConfiguration `mapstructure:"otp-auth" json:"otp_auth"`
}

// 系统配置
type SystemConfiguration struct {
	ApiPrefix  string `mapstructure:"api-prefix" json:"api_prefix"`
	ApiVersion string `mapstructure:"api-version" json:"api_version"`
}

// 日志配置
type LogConfiguration struct {
	Enabled    bool          `mapstructure:"enabled" json:"enabled"`
	Level      zapcore.Level `mapstructure:"level" json:"level"`
	Path       string        `mapstructure:"path" json:"path"`
	MaxSize    int           `mapstructure:"max-size" json:"max_size"`
	MaxAge     int           `mapstructure:"max-age" json:"max_age"`
	MaxBackups int           `mapstructure:"max-backups" json:"max_backups"`
	Compress   bool          `mapstructure:"compress" json:"compress"`
}

// 数据库配置
type MySQLConfiguration struct {
	Host         string `mapstructure:"host" json:"host"`
	Port         int    `mapstructure:"port" json:"port"`
	Database     string `mapstructure:"database" json:"database"`
	Username     string `mapstructure:"username" json:"username"`
	Password     string `mapstructure:"password" json:"password"`
	Charset      string `mapstructure:"charset" json:"charset"`
	Collation    string `mapstructure:"collation" json:"collation"`
	Timeout      int    `mapstructure:"timeout" json:"timeout"`
	ExtraParam   string `mapstructure:"extra-param" json:"extra_param"`
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"max_idle_conns"`
	MaxIdleTime  int    `mapstructure:"max-idle-time" json:"max_idle_time"`
}

// Redis 配置
type RedisConfiguration struct {
	Host         string `mapstructure:"host" json:"host"`
	Port         int    `mapstructure:"port" json:"port"`
	Database     int    `mapstructure:"database" json:"database"`
	Password     string `mapstructure:"password" json:"password"`
	Timeout      int    `mapstructure:"timeout" json:"timeout"`
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"max_open_conns"`
	MinIdleConns int    `mapstructure:"min-idle-conns" json:"min_idle_conns"`
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"max_idle_conns"`
	MaxIdleTime  int    `mapstructure:"max-idle-time" json:"max_idle_time"`
}

// JWT 配置
type JWTConfiguration struct {
	Realm   string `mapstructure:"realm" json:"realm"`
	Key     string `mapstructure:"key" json:"key"`
	Timeout int    `mapstructure:"timeout" json:"timeout"`
}

// 登录配置
type LoginConfiguration struct {
	PasswordLevel  int  `mapstructure:"password-level" json:"password_level"`
	WrongTimes     int  `mapstructure:"wrong-times" json:"wrong_times"`
	LockTime       int  `mapstructure:"lock-time" json:"lock_time"`
	MultiDevices   bool `mapstructure:"multi-devices" json:"multi_devices"`
	ResetTokenTime int  `mapstructure:"reset-token-time" json:"reset_token_time"`
}

// Minio 配置
type MinioConfiguration struct {
	URL          string `mapstructure:"url" json:"url"`
	AccessKey    string `mapstructure:"access-key" json:"access_key"`
	AccessSecret string `mapstructure:"access-secret" json:"access_secret"`
	SSL          bool   `mapstructure:"ssl" json:"ssl"`
}

// OTP 双因素认证配置
type OTPAuthConfiguration struct {
	Enabled bool   `mapstructure:"enabled" json:"enabled"`
	Issuer  string `mapstructure:"issuer" json:"issuer"`
	Digits  int    `mapstructure:"digits" json:"digits"`
	Period  int    `mapstructure:"period" json:"period"`
}
