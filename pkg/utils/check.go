package utils

import (
    "fmt"
    "net"
    "os"
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
