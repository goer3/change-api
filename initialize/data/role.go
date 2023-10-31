package data

import (
    "change-api/common"
    "change-api/model"
    "errors"

    "gorm.io/gorm"
)

// 角色数据
var roles = []model.Role{
    {
        Id:          1,
        Name:        "超级管理员",
        Description: "System Administrator",
    },
    {
        Id:          2,
        Name:        "访客",
        Description: "Guest",
    },
}

// 角色数据初始化
func Role() {
    var role model.Role
    for _, item := range roles {
        // 查看数据是否存在，如果不存在才执行创建
        err := common.DB.Where("id = ? or name = ?", item.Id, item.Name).First(&role).Error
        if errors.Is(err, gorm.ErrRecordNotFound) {
            common.DB.Create(&item)
        }
    }
}
