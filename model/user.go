package model

import "github.com/golang-module/carbon/v2"

// 用户模型
type User struct {
    Id                 uint            `gorm:"primaryKey;comment:自增编号" json:"id"`
    Name               string          `gorm:"not null;comment:姓名" json:"name"`
    Landline           string          `gorm:"comment:固话号码" json:"landline"`
    Mobile             string          `gorm:"uniqueIndex:uidx_mobile;comment:手机号" json:"mobile"`
    Email              string          `gorm:"uniqueIndex:uidx_email;comment:邮箱" json:"email"`
    Password           string          `gorm:"not null;comment:密码" json:"-"` // json 中不显示 password 字段
    JobId              string          `gorm:"uniqueIndex:uidx_job_id;comment:工号" json:"job_id"`
    JobName            string          `gorm:"not null;comment:岗位名称" json:"job_name"`
    JoinTime           carbon.DateTime `gorm:"comment:入职日期" json:"join_time"`
    DepartmentId       uint            `gorm:"comment:部门id" json:"department_id"` // 关联部门
    Department         Department      `gorm:"foreignKey:DepartmentId;comment:部门" json:"department"`
    OfficeProvinceId   uint            `gorm:"comment:办公地点省id" json:"office_province_id"` // 关联省市区
    OfficeProvince     Province        `gorm:"foreignKey:OfficeProvinceId;comment:省" json:"office_province"`
    OfficeCityId       uint            `gorm:"comment:办公地点市id" json:"office_city_id"`
    OfficeCity         City            `gorm:"foreignKey:OfficeCityId;comment:市" json:"office_city"`
    OfficeAreaId       uint            `gorm:"comment:办公地点区id" json:"office_area_id"`
    OfficeArea         Area            `gorm:"foreignKey:OfficeAreaId;comment:区" json:"office_area"`
    OfficeStreetId     uint            `gorm:"comment:办公地点街道id" json:"office_street_id"`
    OfficeStreet       Street          `gorm:"foreignKey:OfficeStreetId;comment:街道" json:"office_street"`
    OfficeAddress      string          `gorm:"comment:办公地点详细地址" json:"office_address"`
    OfficeStation      string          `gorm:"comment:办公地点工位" json:"office_station"`
    NativeProvinceId   uint            `gorm:"comment:籍贯省id" json:"native_province_id"`
    NativeProvince     Province        `gorm:"foreignKey:NativeProvinceId;comment:省" json:"native_province"`
    NativeCityId       uint            `gorm:"comment:籍贯市id" json:"native_city_id"`
    NativeCity         City            `gorm:"foreignKey:NativeCityId;comment:市" json:"native_city"`
    Gender             *uint           `gorm:"type:tinyint(1);default:1;comment:性别(1=男,2=女,3=未知)" json:"gender"`
    Avatar             string          `gorm:"comment:头像" json:"avatar"`
    Birthday           carbon.DateTime `gorm:"comment:生日" json:"birthday"`
    CreatorId          uint            `gorm:"comment:创建人id" json:"creator_id"` // 关联用户自身
    Creator            *User           `gorm:"foreignKey:CreatorId;-" json:"creator"`
    FirstLogin         *uint           `gorm:"type:tinyint(1);default:1;comment:是否第一次登录(0=否,1=是)" json:"first_login"`
    LastLogin          carbon.DateTime `gorm:"comment:最后一次登录时间" json:"last_login"`
    LastLoginIP        string          `gorm:"comment:最后一次登录IP" json:"last_login_ip"`
    LastChangePassword carbon.DateTime `gorm:"comment:最后一次修改密码时间" json:"last_change_password"`
    WrongTimes         uint            `gorm:"comment:登录错误次数" json:"wrong_times"`
    Status             *uint           `gorm:"type:tinyint(1);default:1;comment:用户状态(0=禁用,1=正常,2=未激活,3=锁定)" json:"status"`
    RoleId             uint            `gorm:"comment:角色id" json:"role_id"` // 关联角色
    Role               Role            `gorm:"foreignKey:RoleId;comment:角色" json:"role"`
    BaseModel                          // 基础字段信息
}

// 自定义表名
func (User) TableName() string {
    return "user"
}
