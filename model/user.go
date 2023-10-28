package model

// 用户模型
type User struct {
    BaseModel
    Name string `gorm:"comment:姓名" json:"name"`
}
