package model

import (
    "database/sql/driver"
    "fmt"
    "github.com/matoous/go-nanoid"
    "gorm.io/gorm"
)

// Nanoid 自定义类型
type Nanoid string

func (n *Nanoid) Scan(value interface{}) error {
    *n = Nanoid(fmt.Sprintf("%s", value))
    return nil
}

func (n *Nanoid) Value() (driver.Value, error) {
    return string(*n), nil
}

func (n *Nanoid) BeforeCreate(tx *gorm.DB) error {
    id, _ := gonanoid.Nanoid()
    *n = Nanoid(id)
    return nil
}
