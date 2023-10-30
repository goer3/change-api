package initialize

import (
    "change-api/common"
    "change-api/model"
)

// 数据结构同步
func Migrate() {
    _ = common.DB.AutoMigrate(new(model.Province))   // 省
    _ = common.DB.AutoMigrate(new(model.City))       // 市
    _ = common.DB.AutoMigrate(new(model.Area))       // 区
    _ = common.DB.AutoMigrate(new(model.Street))     // 街道
    _ = common.DB.AutoMigrate(new(model.Department)) // 部门
    _ = common.DB.AutoMigrate(new(model.User))       // 用户
    _ = common.DB.AutoMigrate(new(model.Menu))       // 菜单
    _ = common.DB.AutoMigrate(new(model.Role))       // 角色
}
