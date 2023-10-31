package utils

import (
    "fmt"
    "net"
    "os"
    "regexp"
)

// 判断文件是否存在
func FileExists(filename string) (bool, error) {
    stat, err := os.Stat(filename)
    if !os.IsNotExist(err) {
        if !stat.IsDir() {
            return true, nil
        }
        return false, fmt.Errorf("the path %s exists, but it is a directory", filename)
    }
    return false, fmt.Errorf("the file %s not exists", filename)
}

// 判断目录是否存储
func DirExists(dirname string) (bool, error) {
    stat, err := os.Stat(dirname)
    if !os.IsNotExist(err) {
        if stat.IsDir() {
            return true, nil
        }
        return false, fmt.Errorf("the path %s exists, but it is a file", dirname)
    }
    return false, fmt.Errorf("the directory %s not exists", dirname)
}

// 判断 IP 地址是否合法
func IsIPAddress(ip string) bool {
    result := net.ParseIP(ip)
    return result != nil
}

// 判断端口是否合法
func IsPort(port string) bool {
    _, err := net.ResolveTCPAddr("tcp", ":"+port)
    return err == nil
}

// 验证邮箱合法性
func IsEmail(email string) bool {
    pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$`
    reg := regexp.MustCompile(pattern)
    return reg.MatchString(email)
}

// 验证手机号合法性
func IsMobile(mobile string) bool {
    pattern := `^((13[0-9])|(14[5|7])|(15([0-3]|[5-9]))|(18[0,5-9]))\d{8}$`
    reg := regexp.MustCompile(pattern)
    return reg.MatchString(mobile)
}

// 密码复杂度层级
const (
    Invalid = iota // 非法密码，密码小于 8 位或者大于 20 位，或者不在指定字符串中
    Weak           // 弱密码，只包含数字或小写字母或大写字母中的一种
    Medium         // 中等密码，只包含数字，大小写字母
    Strong         // 强密码，8-12 位以上，包含数字，大小写字母，特殊符号
    Final          // 究极密码，12 位以上，包含数字，大小写字母，特殊符号
)

// 密码复杂度
func PasswordComplexity(password string) int {
    // 类型
    number := "0-9"           // 数字
    uppercase := "A-Z"        // 大写字母
    lowercase := "a-z"        // 小写字母
    symbol := "~!@#$%^&*()_+" // 特殊符号

    // 不合法
    if len(password) < 8 || len(password) > 20 ||
        !regexp.MustCompile("^["+number+uppercase+lowercase+symbol+"]+$").MatchString(password) {
        return Invalid
    }

    // 弱密码
    if regexp.MustCompile("^["+number+"]+$").MatchString(password) ||
        regexp.MustCompile("^["+uppercase+"]+$").MatchString(password) ||
        regexp.MustCompile("^["+lowercase+"]+$").MatchString(password) {
        return Weak
    }

    // 中等密码
    if regexp.MustCompile("^[" + number + uppercase + lowercase + "]+$").MatchString(password) {
        return Medium
    }

    // 强密码
    if len(password) <= 12 {
        return Strong
    }

    // 究极密码
    return Final
}
