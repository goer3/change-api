package initialize

import (
    "change-api/common"
    "change-api/model"
)

// 数据结构同步
func Migrate() {
    _ = common.DB.AutoMigrate(new(model.User))
}
