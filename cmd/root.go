package cmd

import (
    "change-api/common"
    "fmt"
    "github.com/spf13/cobra"
    "os"
)

// 没有调用子命令时的基础命令
var rootCmd = &cobra.Command{
    Use:   "change-api",
    Short: "change-api is an operations management system backend developed in Golang",
    // 如果有相关的 action 要执行，请取消下面这行代码的注释
    // Run: func(cmd *cobra.Command, args []string) { },
}

// Execute 所有子命令添加到 root 命令，输入 cmd 的入口
func Execute() {
    // 打印 Logo
    fmt.Println(common.Logo)
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
