package model

import "github.com/golang-module/carbon/v2"

// 用户模型
type User struct {
    Id                 uint            `gorm:"primaryKey;comment:自增编号" json:"id"`
    JobId              string          `gorm:"uniqueIndex:uidx_job_id;comment:工号（Username）" json:"job_id"`
    Name               string          `gorm:"not null;comment:姓名" json:"name"`
    Landline           string          `gorm:"comment:固话号码" json:"landline"`
    Mobile             string          `gorm:"uniqueIndex:uidx_mobile;comment:手机号" json:"mobile"`
    Email              string          `gorm:"uniqueIndex:uidx_email;comment:邮箱" json:"email"`
    Password           string          `gorm:"not null;comment:密码" json:"-"` // json 中不显示 password 字段
    JobName            string          `gorm:"not null;comment:岗位名称" json:"job_name"`
    JoinTime           carbon.DateTime `gorm:"comment:入职日期" json:"join_time"`
    OfficeProvinceId   uint            `gorm:"comment:省Id" json:"office_province_id"`
    OfficeProvince     Region          `gorm:"foreignKey:OfficeProvinceId;comment:省" json:"office_province"`
    OfficeCityId       uint            `gorm:"comment:市Id" json:"office_city_id"`
    OfficeCity         Region          `gorm:"foreignKey:OfficeCityId;comment:市" json:"office_city"`
    OfficeDistrictId   uint            `gorm:"comment:区Id" json:"office_district_id"`
    OfficeDistrict     Region          `gorm:"foreignKey:OfficeDistrictId;comment:区" json:"office_district"`
    OfficeAddress      string          `gorm:"comment:详细地址" json:"office_address"`
    NativeProvinceId   uint            `gorm:"comment:省Id" json:"native_province_id"`
    NativeProvince     Region          `gorm:"foreignKey:NativeProvinceId;comment:省" json:"native_province"`
    NativeCityId       uint            `gorm:"comment:市Id" json:"native_city_id"`
    NativeCity         Region          `gorm:"foreignKey:NativeCityId;comment:市" json:"native_city"`
    Gender             *uint           `gorm:"type:tinyint(1);default:1;comment:性别(1: 男, 2: 女, 3: 未知)" json:"gender"`
    Avatar             string          `gorm:"comment:头像" json:"avatar"`
    Birthday           carbon.DateTime `gorm:"comment:生日" json:"birthday"`
    CreatorId          uint            `gorm:"comment:创建人Id" json:"creator_id"`
    Creator            *User           `gorm:"foreignKey:CreatorId;comment:创建人;-" json:"creator"`
    FirstLogin         *uint           `gorm:"type:tinyint(1);default:1;comment:是否第一次登录(0: 否, 1: 是)" json:"first_login"`
    LastLogin          carbon.DateTime `gorm:"comment:最后一次登录时间" json:"last_login"`
    LastLoginIP        string          `gorm:"comment:最后一次登录IP" json:"last_login_ip"`
    LastChangePassword carbon.DateTime `gorm:"comment:最后一次修改密码时间" json:"last_change_password"`
    WrongTimes         uint            `gorm:"comment:登录错误次数" json:"wrong_times"`
    Status             *uint           `gorm:"type:tinyint(1);default:1;comment:用户状态(0: 禁用, 1: 正常, 2: 未激活, 3: 锁定)" json:"status"`
    BaseModel                          // 基础字段信息
}

// 自定义表名
func (User) TableName() string {
    return "user"
}
