package api

import (
    "change-api/common"
    "change-api/pkg/response"
    "net/http"

    "github.com/gin-gonic/gin"
)

// 健康检测接口
func HealthHandler(ctx *gin.Context) {
    ctx.String(http.StatusOK, "OK")
}

// 系统信息接口
func InfoHandler(ctx *gin.Context) {
    response.SuccessWithData(map[string]interface{}{
        "developer": common.Developer.Name,
        "email":     common.Developer.Email,
    })
}

// 系统版本接口
func VersionHandler(ctx *gin.Context) {
    response.SuccessWithData(map[string]interface{}{
        "system": common.Version.SystemVersion,
        "golang": common.Version.GoVersion,
    })
}
