package initialize

import "change-api/initialize/data"

// 数据初始化
func Data() {
    // data.Region()     // 省市区街道数据
    // data.Department() // 部门数据
    data.Role() // 角色数据
    data.Menu() // 菜单数据
    data.User() // 用户数据
}
