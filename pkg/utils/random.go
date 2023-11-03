package utils

import (
    "change-api/common"
    "change-api/model"
    "github.com/matoous/go-nanoid"
    "math/rand"
)

// 生成指定长度的随机字符串
func RandString(n int) string {
    var letters = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
    b := make([]rune, n)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}

// 自定义生成 Nanoid 的方法
func GenerateNanoid() model.Nanoid {
    id, err := gonanoid.Nanoid()
    if err != nil {
        common.Log.Error("Nanoid 生成失败", err.Error())
    }
    return model.Nanoid(id)
}
