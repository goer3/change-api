package initialize

import (
    "change-api/common"
    "change-api/model"
)

// 数据结构同步
func Migrate() {
    _ = common.DB.AutoMigrate(new(model.Region))     // 地区
    _ = common.DB.AutoMigrate(new(model.User))       // 用户
    _ = common.DB.AutoMigrate(new(model.Department)) // 部门
}
