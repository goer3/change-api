package model

// 菜单模型
type Menu struct {
    Id        uint   `gorm:"primaryKey;comment:自增编号" json:"id"`
    Name      string `gorm:"uniqueIndex:uidx_name;comment:菜单名称" json:"name"`
    Icon      string `gorm:"comment:菜单图标" json:"icon"`
    Path      string `gorm:"comment:菜单路径" json:"path"`
    ParentId  uint   `gorm:"comment:父id" json:"parent_id"`
    Children  []Menu `gorm:"-" json:"children"`
    Roles     []Role `gorm:"many2many:role_menu_relation" json:"roles"` // 菜单和角色多对多
    BaseModel                                                           // 基础字段信息
}

// 自定义表名
func (Menu) TableName() string {
    return "menu"
}
