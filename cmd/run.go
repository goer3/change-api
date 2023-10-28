package cmd

import (
    "change-api/common"
    "change-api/initialize"
    "change-api/pkg/utils"
    "context"
    "fmt"
    "net/http"
    "os"
    "os/signal"
    "time"

    "github.com/spf13/cobra"
)

func init() {
    rootCmd.AddCommand(runCmd)

    // 子命令和参数
    runCmd.AddCommand(runServeCmd)
    runServeCmd.Flags().StringVarP(&common.Runtime.Listen, "listen", "l", common.Runtime.Listen, "specify listening address for service, for example: 0.0.0.0, 127.0.0.1")
    runServeCmd.Flags().StringVarP(&common.Runtime.Port, "port", "p", common.Runtime.Port, "specify listening port for service")
    runServeCmd.Flags().StringVarP(&common.Runtime.Config, "config", "f", common.Runtime.Config, "specify running config for service")
}

// 运行命令
var runCmd = &cobra.Command{
    Use:   "run",
    Short: "you can run the service with `run serve`",
}

// 启动命令
var runServeCmd = &cobra.Command{
    Use:   "serve",
    Short: "you can run the service with some flags",
    Run: func(cmd *cobra.Command, args []string) {
        // 配置初始化
        initialize.Config()

        // 日志初始化
        initialize.Logger()

        // 数据库连接初始化
        initialize.MySQL()

        // 路由初始化
        r := initialize.Router()

        // 判断参数是否合法
        if !utils.IsIPAddress(common.Runtime.Listen) {
            common.Log.Error("listening IP address is invalid")
            return
        }

        // 检测端口是否合法
        if !utils.IsPort(common.Runtime.Port) {
            common.Log.Error("listening port is invalid")
            return
        }

        // 监听地址
        listenAddress := fmt.Sprintf("%s:%s", common.Runtime.Listen, common.Runtime.Port)
        common.Log.Info("listening address is: ", listenAddress)

        // 配置服务
        server := http.Server{
            Addr:    listenAddress,
            Handler: r,
        }

        // 启动服务
        go func() {
            err := server.ListenAndServe()
            if err != nil && err != http.ErrServerClosed {
                common.Log.Error(err)
                panic(err)
            }
        }()

        // 接收优雅关闭信号
        quit := make(chan os.Signal, 1)
        signal.Notify(quit, os.Interrupt)
        <-quit

        // 等待5秒然后停止服务
        ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
        defer cancel()
        err := server.Shutdown(ctx)
        if err != nil {
            common.Log.Error(err)
            panic(err)
        }
        common.Log.Info("service shutdown success")
    },
}
