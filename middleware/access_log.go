package middleware

import (
    "fmt"
    "time"

    "github.com/gin-gonic/gin"
)

// 访问日志中间件
func AccessLog(ctx *gin.Context) {
    // 请求时间
    startTime := time.Now()
    // 处理请求
    ctx.Next()
    // 结束时间
    endTime := time.Now()
    // 执行耗时
    execTime := startTime.Sub(endTime)
    // 请求方式
    method := ctx.Request.Method
    // 请求地址
    uri := ctx.Request.RequestURI
    // 状态码
    code := ctx.Writer.Status()
    // 来源 IP
    clientIP := ctx.ClientIP()

    // 完整的日志
    logStr := fmt.Sprintf("%s\t%s\t%d\t%s\t%s",
        method,
        uri,
        code,
        execTime.String(),
        clientIP,
    )

    // 打印日志，OPTIONS 请求使用 DEBUG
    if method == "OPTIONS" {
        fmt.Println(logStr)
    } else {
        fmt.Println(logStr)
    }
}
