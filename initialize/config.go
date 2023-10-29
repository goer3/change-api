package initialize

import (
    "bytes"
    "change-api/common"
    "change-api/pkg/log2"
    "change-api/pkg/utils"
    "os"

    "github.com/spf13/viper"
)

// 配置文件初始化
func Config() {
    // 读取的数据
    var bs []byte
    var err error

    // viper 初始化
    v := viper.New()
    v.SetConfigType("yaml")

    // 读取指定的配置的文件
    filename := common.Runtime.Config

    // 先找本地是否存在，只有当本地不存在的时候才会去读取打包的默认配置
    exists, _ := utils.FileExists(filename)
    if !exists {
        // 读取 embed 打包的配置
        bs, err = common.FS.ReadFile(filename)
        if err != nil {
            panic(err)
        }
    } else {
        // 本地存在则直接读取配置文件
        bs, err = os.ReadFile(filename)
        if err != nil {
            panic(err)
        }
    }

    // viper 读取解析配置文件
    err = v.ReadConfig(bytes.NewReader(bs))
    if err != nil {
        panic(err)
    }

    // 将配置存到内存中
    settings := v.AllSettings()
    for idx, setting := range settings {
        v.Set(idx, setting)
    }

    // 将配置复制给全局变量，方便后续使用
    err = v.Unmarshal(&common.Config)
    if err != nil {
        panic(err)
    }

    // 配置读取完成
    log2.INFO("config initialize success:", filename)
}
