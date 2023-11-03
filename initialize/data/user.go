package data

import (
    "change-api/common"
    "change-api/model"
    "change-api/pkg/utils"
    "errors"

    "github.com/golang-module/carbon/v2"
    "gorm.io/gorm"
)

// 初始化密码
var password = "123456"

// 用户初始化数据
var users = []model.User{
    {
        Id:       "admin",
        Name:     "超管",
        Mobile:   "18888888888",
        Email:    "admin@ezops.cn",
        Password: utils.CryptoPassword(password),
        JobId:    "ez000001",
        JobName:  "高级运维工程师",
        JoinTime: carbon.DateTime{
            Carbon: carbon.Now(),
        },
        DepartmentId:     109400,
        OfficeProvinceId: 44,
        OfficeCityId:     4403,
        OfficeAreaId:     440304,
        OfficeStreetId:   440304005,
        OfficeAddress:    "下沙社区",
        OfficeStation:    "T1-9F-A-11",
        NativeProvinceId: 51,
        NativeCityId:     5105,
        Gender:           common.Male,
        Avatar:           "/images/avatar/default.png",
        Birthday: carbon.DateTime{
            Carbon: carbon.Now(),
        },
        CreatorId: 0,
        RoleId:    1,
    },
    {
        Id:       "guest",
        Name:     "访客",
        Mobile:   "19999999999",
        Email:    "guest@ezops.cn",
        Password: utils.CryptoPassword(password),
        JobId:    "ez000002",
        JobName:  "访客",
        JoinTime: carbon.DateTime{
            Carbon: carbon.Now(),
        },
        DepartmentId:     109400,
        OfficeProvinceId: 44,
        OfficeCityId:     4403,
        OfficeAreaId:     440304,
        OfficeStreetId:   440304005,
        OfficeAddress:    "下沙社区",
        OfficeStation:    "T1-9F-A-12",
        NativeProvinceId: 51,
        NativeCityId:     5105,
        Gender:           common.Female,
        Avatar:           "/images/avatar/default1.png",
        Birthday: carbon.DateTime{
            Carbon: carbon.Now(),
        },
        CreatorId: 0,
        RoleId:    2,
    },
}

// 用户数据初始化
func User() {
    var user model.User
    for _, item := range users {
        // 查看数据是否存在，如果不存在才执行创建
        err := common.DB.Where("id = ? or mobile = ? or email = ? or job_id = ?",
            item.Id,
            item.Mobile,
            item.Email,
            item.JobId).First(&user).Error
        if errors.Is(err, gorm.ErrRecordNotFound) {
            common.DB.Create(&item)
        }
    }
}
