package model

import "github.com/golang-module/carbon/v2"

// 基础模型，为了覆盖本身的 time.Time，换成 carbon 的方式，原因如下：
// 1. carbon 提供了丰富的日期和时间操作方法，比如添加、减去时间间隔、格式化日期时间、获取星期几等
// 2. carbon 支持全球范围内的时区，可以轻松地进行时区转换和时区相关的操作
// 3. carbon 可以根据不同的日期时间格式，解析字符串并转换为日期时间对象
// 使用文档：https://juejin.cn/post/6925036967151288328
type BaseModel struct {
    Id        uint            `gorm:"primaryKey;comment:自增编号" json:"id"`
    CreatedAt carbon.DateTime `gorm:"comment:创建时间" json:"created_at"`
    UpdatedAt carbon.DateTime `gorm:"comment:更新时间" json:"updated_at"`
    DeletedAt DeletedAt       `gorm:"comment:删除时间" json:"deleted_at"`
}
