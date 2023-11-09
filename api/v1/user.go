package v1

import (
    "change-api/common"
    "change-api/dto"
    "change-api/model"
    "change-api/pkg/response"
    "github.com/gin-gonic/gin"
)

// 用户列表
func UserListHandler(ctx *gin.Context) {
    // 获取用户的分页信息
    var req dto.Pagination
    if err := ctx.ShouldBindQuery(&req); err != nil {
        response.FailedWithCode(response.ParamError)
    }

    // 搜索条件
    dbt := common.DB.Preload("Department").
        Preload("Role")

    // 根据是否需要分页来解决查询方法
    var users []model.User
    // 更新请求信息中的 Total
    dbt.Find(&model.User{}).Count(&req.Total)
    // 不分页
    if req.NoPagination {
        dbt.Find(&users)
    } else {
        // 获取分页限制和偏移量
        limit, offset := req.GetLimitAndOffset()
        // 根据偏移量查询
        dbt.Limit(limit).Offset(offset).Find(&users)
    }

    // 请求响应
    response.SuccessWithData(dto.PageResponse{
        Pagination: req,
        List:       users,
    })
}
