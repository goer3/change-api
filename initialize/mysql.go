package initialize

import (
    "change-api/common"
    "fmt"
    "time"

    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "gorm.io/gorm/schema"
)

// MySQL 初始化
func MySQL() {
    // 数据库连接串
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&collation=%s&timeout=%dms&%s",
        common.Config.MySQL.Username,
        common.Config.MySQL.Password,
        common.Config.MySQL.Host,
        common.Config.MySQL.Port,
        common.Config.MySQL.Database,
        common.Config.MySQL.Charset,
        common.Config.MySQL.Collation,
        common.Config.MySQL.Timeout,
        common.Config.MySQL.ExtraParam)

    // 连接数据库
    db, err := gorm.Open(mysql.New(mysql.Config{
        DSN:               dsn, // 数据库连接字符串
        DefaultStringSize: 170, // varchar 默认长度，太长影响查询
    }), &gorm.Config{
        NamingStrategy: schema.NamingStrategy{
            SingularTable: true, // 单数表名
            // TablePrefix:   "tb_", // 表名前缀
        },
        DisableForeignKeyConstraintWhenMigrating: true,  // 禁用外键
        IgnoreRelationshipsWhenMigrating:         false, // 开启会导致 many2many 的表创建失败
        QueryFields:                              true,  // 解决查询索引失效问题
    })

    // 错误处理
    if err != nil {
        common.Log.Error("database connect failed")
        common.Log.Error(err)
        panic("database connect failed")
    }

    // 设置数据库连接池
    sqlDB, _ := db.DB()
    sqlDB.SetMaxOpenConns(common.Config.MySQL.MaxOpenConns)
    sqlDB.SetMaxIdleConns(common.Config.MySQL.MaxIdleConns)
    sqlDB.SetConnMaxIdleTime(time.Duration(common.Config.MySQL.MaxIdleTime) * time.Minute)

    // 设置全局数据库连接，方便后续使用
    common.DB = db
    common.Log.Info("database initialize success: ", fmt.Sprintf("%s:%d", common.Config.MySQL.Host, common.Config.MySQL.Port))
}
