package cmd

import (
    "change-api/common"
    "fmt"
    "github.com/spf13/cobra"
)

func init() {
    RootCmd.AddCommand(VersionCmd)
}

// VersionCmd 版本信息
var VersionCmd = &cobra.Command{
    Use:   "version",
    Short: "Show the current version of the service",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("System version:", common.Version.SystemVersion)
        fmt.Println("Golang version:", common.Version.GoVersion)
    },
}
