package model

import (
    "change-api/common"
    "fmt"
    "github.com/matoous/go-nanoid"
)

// Nanoid 自定义类型
type Nanoid string

func (n *Nanoid) Scan(value interface{}) error {
    *n = Nanoid(fmt.Sprintf("%s", value))
    return nil
}

func (n Nanoid) Value() (interface{}, error) {
    return string(n), nil
}

// 自定义生成 Nanoid 的方法
func GenerateNanoid() Nanoid {
    id, err := gonanoid.Nanoid()
    if err != nil {
        common.Log.Error("Nanoid 生成失败", err.Error())
    }
    return Nanoid(id)
}
