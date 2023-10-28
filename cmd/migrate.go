package cmd

import (
    "change-api/initialize"
    "change-api/pkg/log2"
    "github.com/spf13/cobra"
)

func init() {
    rootCmd.AddCommand(migrateCmd)
}

// 同步数据结构命令
var migrateCmd = &cobra.Command{
    Use:   "migrate",
    Short: "you can migrate table to MySQL database",
    Run: func(cmd *cobra.Command, args []string) {
        log2.INFO("start migrate table to MySQL database")
        initialize.Config()  // 初始化配置
        initialize.Logger()  // 初始化日志
        initialize.MySQL()   // 初始化数据库
        initialize.Migrate() // 同步数据结构
        log2.INFO("migrate table to MySQL database success")
    },
}
