package data

import (
    "change-api/common"
    "change-api/model"
    "errors"
    "gorm.io/gorm"
)

// 测试部门数据
var departments = []model.Department{
    {
        Id:       100000,
        Name:     "集团总部",
        LeaderId: 1,
        ParentId: 0,
        Children: []model.Department{
            {
                Id:       101000,
                Name:     "总裁办",
                LeaderId: 1,
                ParentId: 100000,
            },
            {
                Id:       102000,
                Name:     "人力资源中心",
                LeaderId: 1,
                ParentId: 100000,
            },
            {
                Id:       103000,
                Name:     "行政中心",
                LeaderId: 1,
                ParentId: 100000,
            },
            {
                Id:       104000,
                Name:     "财务部",
                LeaderId: 1,
                ParentId: 100000,
            },
            {
                Id:       105000,
                Name:     "商务部",
                LeaderId: 1,
                ParentId: 100000,
            },
            {
                Id:       106000,
                Name:     "品牌公关部",
                LeaderId: 1,
                ParentId: 100000,
            },
            {
                Id:       107000,
                Name:     "销售部",
                LeaderId: 1,
                ParentId: 100000,
            },
            {
                Id:       108000,
                Name:     "产品中心",
                LeaderId: 1,
                ParentId: 100000,
                Children: []model.Department{
                    {
                        Id:       108100,
                        Name:     "UI设计部",
                        LeaderId: 1,
                        ParentId: 108000,
                    },
                    {
                        Id:       108200,
                        Name:     "产品部",
                        LeaderId: 1,
                        ParentId: 108000,
                    },
                },
            },
            {
                Id:       109000,
                Name:     "研发中心",
                LeaderId: 1,
                ParentId: 100000,
                Children: []model.Department{
                    {
                        Id:       109100,
                        Name:     "前端开发部",
                        LeaderId: 1,
                        ParentId: 109000,
                    },
                    {
                        Id:       109200,
                        Name:     "后端开发部",
                        LeaderId: 1,
                        ParentId: 109000,
                    },
                    {
                        Id:       109300,
                        Name:     "架构部",
                        LeaderId: 1,
                        ParentId: 109000,
                    },
                    {
                        Id:       109400,
                        Name:     "运维部",
                        LeaderId: 1,
                        ParentId: 109000,
                        Children: []model.Department{
                            {
                                Id:       109410,
                                Name:     "系统组",
                                LeaderId: 1,
                                ParentId: 109400,
                            },
                            {
                                Id:       109420,
                                Name:     "运维开发组",
                                LeaderId: 1,
                                ParentId: 109400,
                            },
                            {
                                Id:       109430,
                                Name:     "监控组",
                                LeaderId: 1,
                                ParentId: 109400,
                            },
                            {
                                Id:       109440,
                                Name:     "中间件组",
                                LeaderId: 1,
                                ParentId: 109400,
                            },
                            {
                                Id:       109450,
                                Name:     "网络组",
                                LeaderId: 1,
                                ParentId: 109400,
                            },
                        },
                    },
                    {
                        Id:       109500,
                        Name:     "测试部",
                        LeaderId: 1,
                        ParentId: 109000,
                    },
                },
            },
        },
    },
}

// 递归插入数据方法
func insertDepartmentsData(departments []model.Department) {
    var department model.Department
    for _, item := range departments {
        // 查看数据是否存在，如果不存在才执行创建
        err := common.DB.Where("id = ? or name = ?", item.Id, item.Name).First(&department).Error
        if errors.Is(err, gorm.ErrRecordNotFound) {
            common.DB.Create(&item)
        }
        // 递归插入子部门
        if len(item.Children) > 0 {
            insertDepartmentsData(item.Children)
        }
    }
}

// 部门数据初始化
func Department() {
    insertDepartmentsData(departments)
}
