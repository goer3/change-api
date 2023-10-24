package cmd

import (
    "change-api/common"
    "fmt"
    "github.com/spf13/cobra"
)

func init() {
    RootCmd.AddCommand(InfoCmd)
}

// InfoCmd 开发者信息
var InfoCmd = &cobra.Command{
    Use:   "info",
    Short: "Show the information of the service",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Developer:", common.Developer.Name)
        fmt.Println("Email:", common.Developer.Email)
    },
}
