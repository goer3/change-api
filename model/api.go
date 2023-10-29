package model

// 接口模型
type API struct {
    Id       uint   `gorm:"primaryKey;comment:自增编号" json:"id"`
    Name     string `gorm:"uniqueIndex:uidx_name;comment:接口名称" json:"name"`
    API      string `gorm:"uniqueIndex:uidx_api;comment:接口地址" json:"api"`
    Method   string `gorm:"comment:请求方法，如：GET,POST,DELETE" json:"method"`
    ParentId uint   `gorm:"comment:父id" json:"parent_id"`
    Children []API  `gorm:"-" json:"children"`
}

// 自定义表名
func (API) TableName() string {
    return "api"
}
