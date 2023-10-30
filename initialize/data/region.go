package data

import (
    "change-api/common"
    "fmt"
    "strings"
)

// 执行 SQL 文件
func ImportSQLFile(filename string, table string, truncate bool) {
    bs, err := common.FS.ReadFile(filename)
    if err != nil {
        common.Log.Error(err)
        panic(err)
    }
    // 判断是否需要提前清空表
    if truncate {
        common.DB.Exec(fmt.Sprintf("TRUNCATE TABLE %s;", table))
    }
    // 导入数据
    sqls := strings.Split(string(bs), ";")
    for _, sql := range sqls {
        common.DB.Table(table).Exec(sql)
    }
}

// 初始化省市区街道数据
func Region() {
    ImportSQLFile("initialize/data/sql/province.sql", "province", true)
    ImportSQLFile("initialize/data/sql/city.sql", "city", true)
    ImportSQLFile("initialize/data/sql/area.sql", "area", true)
    ImportSQLFile("initialize/data/sql/street.sql", "street", true)
}
