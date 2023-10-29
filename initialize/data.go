package initialize

import (
    "change-api/common"
    "change-api/model"
    "errors"
    "gorm.io/gorm"
)

// 数据初始化
func Data() {
    var users = []model.User{
        {
            Id:   1,
            Name: "张三",
        },
        {
            Id:   2,
            Name: "李四",
        },
        {
            Id:   3,
            Name: "王五",
        },
    }

    // 创建用户
    var user model.User
    for _, item := range users {
        // 查看数据是否存在，如果不存在才执行创建
        err := common.DB.Where("id = ?", item.Id).First(&user).Error
        if errors.Is(err, gorm.ErrRecordNotFound) {
            common.DB.Create(&item)
        }
    }
}
