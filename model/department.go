package model

// 部门模型
type Department struct {
    Id        uint         `gorm:"primaryKey;comment:自增编号" json:"id"`
    Name      string       `gorm:"comment:部门名称" json:"name"`
    LeaderId  uint         `gorm:"comment:负责人id" json:"leader_id"`
    ParentId  uint         `gorm:"comment:父id" json:"parent_id"`
    Users     []uint       `gorm:"-" json:"users"`    // 用户
    Children  []Department `gorm:"-" json:"children"` // 子部门关联
    BaseModel                                         // 基础字段信息
}

// 自定义表名
func (Department) TableName() string {
    return "department"
}
