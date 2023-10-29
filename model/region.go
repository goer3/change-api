package model

// 国内省市区模型
// 数据来源：https://github.com/modood/Administrative-divisions-of-China
type Region struct {
    Id       uint     `gorm:"primaryKey;comment:id" json:"id"`
    Name     string   `gorm:"comment:省市区名称" json:"name"`
    Level    uint     `gorm:"comment:级别" json:"level"`
    ParentId uint     `gorm:"comment:父id" json:"parent_id"`
    Children []Region `gorm:"-" json:"children"` // 下属区域关联
}

// 自定义表名
func (Region) TableName() string {
    return "region"
}
