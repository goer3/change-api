package main

import (
    "change-api/cmd"
    "change-api/common"
    "embed"
)

//go:embed config/*
var fs embed.FS // 固定格式，打包的时候会将 config 下面的文件都一起打包

func main() {
    common.FS = fs
    cmd.Execute()
}
