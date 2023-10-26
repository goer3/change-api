package log2

import (
    "change-api/common"
    "fmt"
    "time"
)

// 日志打印
func PrintLog(now string, level string, v ...interface{}) {
    fmt.Print(fmt.Sprintf("%s\t%s\t", now, level), fmt.Sprintln(v...))
}

// 系统日志
func SYSTEM(v ...interface{}) {
    now := time.Now().Format(common.MsecTimeFormat)
    level := "SYSTEM"
    PrintLog(now, level, v...)
}

// DEBUG 日志
func DEBUG(v ...interface{}) {
    now := time.Now().Format(common.MsecTimeFormat)
    level := "DEBUG"
    PrintLog(now, level, v...)
}

// INFO 日志
func INFO(v ...interface{}) {
    now := time.Now().Format(common.MsecTimeFormat)
    level := "INFO"
    PrintLog(now, level, v...)
}

// WARN 日志
func WARN(v ...interface{}) {
    now := time.Now().Format(common.MsecTimeFormat)
    level := "WARN"
    PrintLog(now, level, v...)
}

// ERROR 日志
func ERROR(v ...interface{}) {
    now := time.Now().Format(common.MsecTimeFormat)
    level := "ERROR"
    PrintLog(now, level, v...)
}
