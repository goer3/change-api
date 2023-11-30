package cmd

import (
	"change-api/initialize"
	"change-api/pkg/log2"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

// 数据初始化命令
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize data to MySQL database",
	Run: func(cmd *cobra.Command, args []string) {
		// 数据
		log2.INFO("start initialize data to MySQL database")
		initialize.Config() // 初始化配置
		initialize.Logger() // 初始化日志
		initialize.MySQL()  // 初始化数据库
		initialize.Data()   // 数据初始化
		log2.INFO("initialize data to MySQL database success")

		// Minio Bucket
		log2.INFO("start initialize minio bucket")
		initialize.Minio()       // 初始化 Minio 连接
		initialize.MinioBucket() // 初始化 Bucket
		log2.INFO("initialize minio bucket success")
	},
}
