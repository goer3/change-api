package data

import (
    "change-api/common"
    "change-api/model"
    "errors"
    "gorm.io/gorm"
)

// 测试菜单数据
var menus = []model.Menu{
    {
        Id:       1000,
        Name:     "工作空间",
        Icon:     "HomeOutlined",
        Path:     "/dashboard",
        Sort:     0,
        ParentId: 0,
    },
    {
        Id:       1010,
        Name:     "流程审批",
        Icon:     "AuditOutlined",
        Path:     "/workflow",
        Sort:     0,
        ParentId: 0,
        Children: []model.Menu{
            {
                Id:       1011,
                Name:     "流程创建",
                Icon:     "",
                Path:     "/workflow/create",
                Sort:     0,
                ParentId: 1010,
            },
            {
                Id:       1012,
                Name:     "流程列表",
                Icon:     "",
                Path:     "/workflow/list",
                Sort:     0,
                ParentId: 1010,
            },
            {
                Id:       1013,
                Name:     "待我审批",
                Icon:     "",
                Path:     "/workflow/approve",
                Sort:     0,
                ParentId: 1010,
            },
        },
    },
    {
        Id:       1020,
        Name:     "作业执行",
        Icon:     "ScanOutlined",
        Path:     "/job",
        Sort:     0,
        ParentId: 0,
        Children: []model.Menu{
            {
                Id:       1021,
                Name:     "定时任务",
                Icon:     "",
                Path:     "/job/cron",
                Sort:     0,
                ParentId: 1020,
            },
            {
                Id:       1022,
                Name:     "批量任务",
                Icon:     "",
                Path:     "/job/multiple",
                Sort:     0,
                ParentId: 1020,
            },
            {
                Id:       1023,
                Name:     "任务监控",
                Icon:     "",
                Path:     "/job/monitor",
                Sort:     0,
                ParentId: 1020,
            },
        },
    },
    {
        Id:       1030,
        Name:     "资产管理",
        Icon:     "CodeOutlined",
        Path:     "/asset",
        Sort:     0,
        ParentId: 0,
        Children: []model.Menu{
            {
                Id:       1031,
                Name:     "自建机房",
                Icon:     "",
                Path:     "/asset/owen",
                Sort:     0,
                ParentId: 1030,
            },
            {
                Id:       1032,
                Name:     "云联托管",
                Icon:     "",
                Path:     "/asset/cloud",
                Sort:     0,
                ParentId: 1030,
            },
            {
                Id:       1033,
                Name:     "云商接入",
                Icon:     "",
                Path:     "/asset/business",
                Sort:     0,
                ParentId: 1030,
            },
        },
    },
    {
        Id:       1040,
        Name:     "域名解析",
        Icon:     "ApiOutlined",
        Path:     "/dns",
        Sort:     0,
        ParentId: 0,
        Children: []model.Menu{
            {
                Id:       1041,
                Name:     "内网解析",
                Icon:     "",
                Path:     "/dns/in",
                Sort:     0,
                ParentId: 1040,
            },
            {
                Id:       1042,
                Name:     "自建接入",
                Icon:     "",
                Path:     "/dns/add",
                Sort:     0,
                ParentId: 1040,
            },
            {
                Id:       1043,
                Name:     "外网解析",
                Icon:     "",
                Path:     "/dns/out",
                Sort:     0,
                ParentId: 1040,
            },
            {
                Id:       1044,
                Name:     "网商接入",
                Icon:     "",
                Path:     "/dns/business",
                Sort:     0,
                ParentId: 1040,
            },
        },
    },
    {
        Id:       1050,
        Name:     "制品管理",
        Icon:     "BranchesOutlined",
        Path:     "/artifact",
        Sort:     0,
        ParentId: 0,
        Children: []model.Menu{
            {
                Id:       1051,
                Name:     "制品构建",
                Icon:     "",
                Path:     "/artifact/build",
                Sort:     0,
                ParentId: 1050,
            },
            {
                Id:       1052,
                Name:     "制品仓库",
                Icon:     "",
                Path:     "/artifact/repository",
                Sort:     0,
                ParentId: 1050,
            },
            {
                Id:       1053,
                Name:     "构建模板",
                Icon:     "",
                Path:     "/artifact/template",
                Sort:     0,
                ParentId: 1050,
            },
        },
    },
    {
        Id:       1060,
        Name:     "应用发布",
        Icon:     "FlagOutlined",
        Path:     "/deploy",
        Sort:     0,
        ParentId: 0,
        Children: []model.Menu{
            {
                Id:       1061,
                Name:     "主机发布",
                Icon:     "",
                Path:     "/deploy/build",
                Sort:     0,
                ParentId: 1060,
            },
            {
                Id:       1062,
                Name:     "容器发布",
                Icon:     "",
                Path:     "/deploy/container",
                Sort:     0,
                ParentId: 1060,
            },
            {
                Id:       1063,
                Name:     "发布模板",
                Icon:     "",
                Path:     "/deploy/template",
                Sort:     0,
                ParentId: 1060,
            },
        },
    },
    {
        Id:       1070,
        Name:     "监控告警",
        Icon:     "AlertOutlined",
        Path:     "/alert",
        Sort:     0,
        ParentId: 0,
        Children: []model.Menu{
            {
                Id:       1071,
                Name:     "监控大盘",
                Icon:     "",
                Path:     "/alert/dashboard",
                Sort:     0,
                ParentId: 1070,
            },
            {
                Id:       1072,
                Name:     "告警规则",
                Icon:     "",
                Path:     "/alert/rule",
                Sort:     0,
                ParentId: 1070,
            },
            {
                Id:       1073,
                Name:     "告警媒介",
                Icon:     "",
                Path:     "/alert/message",
                Sort:     0,
                ParentId: 1070,
            },
            {
                Id:       1074,
                Name:     "告警历史",
                Icon:     "",
                Path:     "/alert/history",
                Sort:     0,
                ParentId: 1070,
            },
        },
    },
    {
        Id:       1080,
        Name:     "项目维护",
        Icon:     "PartitionOutlined",
        Path:     "/project",
        Sort:     0,
        ParentId: 0,
        Children: []model.Menu{
            {
                Id:       1081,
                Name:     "项目列表",
                Icon:     "",
                Path:     "/project/list",
                Sort:     0,
                ParentId: 1080,
            },
            {
                Id:       1082,
                Name:     "应用列表",
                Icon:     "",
                Path:     "/project/applications",
                Sort:     0,
                ParentId: 1080,
            },
        },
    },
    {
        Id:       1090,
        Name:     "用户中心",
        Icon:     "TeamOutlined",
        Path:     "/users",
        Sort:     0,
        ParentId: 0,
        Children: []model.Menu{
            {
                Id:       1091,
                Name:     "用户管理",
                Icon:     "",
                Path:     "/users/list",
                Sort:     0,
                ParentId: 1090,
            },
            {
                Id:       1092,
                Name:     "分组管理",
                Icon:     "",
                Path:     "/users/group",
                Sort:     0,
                ParentId: 1090,
            },
            {
                Id:       1093,
                Name:     "角色管理",
                Icon:     "",
                Path:     "/users/role",
                Sort:     0,
                ParentId: 1090,
            },
        },
    },
    {
        Id:       1100,
        Name:     "系统配置",
        Icon:     "SettingOutlined",
        Path:     "/system",
        Sort:     0,
        ParentId: 0,
        Children: []model.Menu{
            {
                Id:       1101,
                Name:     "部门管理",
                Icon:     "",
                Path:     "/system/department",
                Sort:     0,
                ParentId: 1100,
            },
            {
                Id:       1102,
                Name:     "菜单管理",
                Icon:     "",
                Path:     "/system/menu",
                Sort:     0,
                ParentId: 1100,
            },
            {
                Id:       1103,
                Name:     "接口管理",
                Icon:     "",
                Path:     "/system/api",
                Sort:     0,
                ParentId: 1100,
            },
            {
                Id:       1104,
                Name:     "服务配置",
                Icon:     "",
                Path:     "/system/setting",
                Sort:     0,
                ParentId: 1100,
            },
        },
    },
    {
        Id:       1110,
        Name:     "日志审计",
        Icon:     "InsuranceOutlined",
        Path:     "/log",
        Sort:     0,
        ParentId: 0,
        Children: []model.Menu{
            {
                Id:       1111,
                Name:     "操作日志",
                Icon:     "",
                Path:     "/log/operation",
                Sort:     0,
                ParentId: 1110,
            },
            {
                Id:       1112,
                Name:     "登录日志",
                Icon:     "",
                Path:     "/log/login",
                Sort:     0,
                ParentId: 1110,
            },
            {
                Id:       1113,
                Name:     "改密日志",
                Icon:     "",
                Path:     "/log/password",
                Sort:     0,
                ParentId: 1110,
            },
            {
                Id:       1114,
                Name:     "机器日志",
                Icon:     "",
                Path:     "/log/machine",
                Sort:     0,
                ParentId: 1110,
            },
        },
    },
    {
        Id:       1120,
        Name:     "个人中心",
        Icon:     "UserOutlined",
        Path:     "/me",
        Sort:     0,
        ParentId: 0,
    },
    {
        Id:       1130,
        Name:     "获取帮助",
        Icon:     "FileProtectOutlined",
        Path:     "/help",
        Sort:     0,
        ParentId: 0,
    },
}

// 递归插入数据方法
func insertMenusData(menus []model.Menu) {
    var menu model.Menu
    for _, item := range menus {
        // 查看数据是否存在，如果不存在才执行创建
        err := common.DB.Where("id = ? or name = ? or path = ?", item.Id, item.Name, item.Path).First(&menu).Error
        if errors.Is(err, gorm.ErrRecordNotFound) {
            common.DB.Create(&item)
        }
        // 递归插入子菜单
        if len(item.Children) > 0 {
            insertMenusData(item.Children)
        }
    }
}

// 菜单数据初始化
func Menu() {
    insertMenusData(menus)
}
