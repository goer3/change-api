package cmd

import (
    "change-api/common"
    "fmt"
    "github.com/spf13/cobra"
)

func init() {
    RootCmd.AddCommand(RunCmd)

    // 子命令和参数
    RunCmd.AddCommand(ServeCmd)
    ServeCmd.Flags().StringVarP(&common.Runtime.Listen, "listen", "l", common.Runtime.Listen, "Specify listening address for service, for example: 0.0.0.0, 127.0.0.1")
    ServeCmd.Flags().StringVarP(&common.Runtime.Port, "port", "p", common.Runtime.Port, "Specify listening port for service")
    ServeCmd.Flags().StringVarP(&common.Runtime.Config, "config", "f", common.Runtime.Config, "Specify running config for service")
}

// RunCmd 运行命令
var RunCmd = &cobra.Command{
    Use:   "run",
    Short: "You can run the service with `run serve`",
}

// ServeCmd 启动命令
var ServeCmd = &cobra.Command{
    Use:   "serve",
    Short: "You can run the service with some flags",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("服务开始运行")
    },
}
