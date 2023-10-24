package dto

// Runtime 服务运行信息
type Runtime struct {
    Listen string // 监听地址
    Port   string // 监听端口
    Config string // 指定配置
}

// Developer 开发者信息
type Developer struct {
    Name  string // 开发者名字
    Email string // 开发者邮箱
}

// Version 系统信息
type Version struct {
    SystemVersion string // 系统版本信息
    GoVersion     string // Golang 版本信息
}
