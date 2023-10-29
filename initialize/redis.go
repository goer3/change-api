package initialize

import (
    "change-api/common"
    "context"
    "fmt"
    "github.com/redis/go-redis/v9"
    "time"
)

// Redis 初始化
func Redis() {
    // Redis 连接串
    dsn := fmt.Sprintf("%s:%d", common.Config.Redis.Host, common.Config.Redis.Port)

    // 配置 Redis 连接
    client := redis.NewClient(&redis.Options{
        Network:         "tcp",                                                        // 协议
        Addr:            dsn,                                                          // 连接串
        DB:              common.Config.Redis.Database,                                 // 数据库
        Password:        common.Config.Redis.Password,                                 // 密码，没有则为空
        MaxRetries:      0,                                                            // 最大重试次数，0 为不重试
        MinRetryBackoff: 8 * time.Millisecond,                                         // 重试时间间隔下限
        MaxRetryBackoff: 512 * time.Millisecond,                                       // 重试时间间隔上限
        DialTimeout:     time.Duration(common.Config.Redis.Timeout) * time.Second,     // 连接超时时间
        ReadTimeout:     3 * time.Second,                                              // 读超时时间
        WriteTimeout:    3 * time.Second,                                              // 写超时时间
        PoolSize:        common.Config.Redis.MaxOpenConns,                             // 最大连接数，一般比 CPU 核数 4 倍少一点
        PoolTimeout:     4,                                                            // 连接等待超时时间，一般是 read 超时时间 +1
        MinIdleConns:    common.Config.Redis.MinIdleConns,                             // 最小空闲连接
        MaxIdleConns:    common.Config.Redis.MaxIdleConns,                             // 最大空闲连接
        ConnMaxIdleTime: time.Duration(common.Config.Redis.MaxIdleTime) * time.Minute, // 最大空闲时间
        ConnMaxLifetime: 0,                                                            // 连接存活时长
    })

    // 测试能否连接
    _, err := client.Ping(context.Background()).Result()
    if err != nil {
        common.Log.Error("redis connect failed")
        common.Log.Error(err)
        panic("redis connect failed")
    }

    // 配置全局，方便后续使用
    common.Cache = client
    common.Log.Info("redis initialize success: ", fmt.Sprintf("%s:%d/%d",
        common.Config.Redis.Host,
        common.Config.Redis.Port,
        common.Config.Redis.Database))
}
