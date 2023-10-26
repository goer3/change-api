package cmd

import (
    "change-api/common"
    "change-api/pkg/log2"
    "os"

    "github.com/spf13/cobra"
)

func init() {
    rootCmd.AddCommand(exportCmd)
    exportCmd.AddCommand(exportConfigCmd)
}

// 导出命令
var exportCmd = &cobra.Command{
    Use:   "export",
    Short: "you can export default config with `export config`",
}

// 导出默认配置命令
var exportConfigCmd = &cobra.Command{
    Use:   "config",
    Short: "you can export default config to file config.yaml",
    Run: func(cmd *cobra.Command, args []string) {
        log2.INFO("start export default config to config.yaml")
        bs, _ := common.FS.ReadFile(common.Runtime.Config)
        err := os.WriteFile("config.yaml", bs, 0644) // 将数据写入文件
        if err != nil {
            log2.ERROR("export default config to config.yaml failed")
            log2.ERROR(err)
            return
        }
        log2.INFO("export default config to config.yaml success")
    },
}
