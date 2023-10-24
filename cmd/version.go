package cmd

import (
    "change-api/common"
    "fmt"
    "github.com/spf13/cobra"
)

func init() {
    rootCmd.AddCommand(versionCmd)
}

// 版本信息
var versionCmd = &cobra.Command{
    Use:   "version",
    Short: "Show the current version of the service",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("System version:", common.Version.SystemVersion)
        fmt.Println("Golang version:", common.Version.GoVersion)
    },
}
