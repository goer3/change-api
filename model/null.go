package model

import (
    "change-api/common"
    "database/sql/driver"
    "fmt"
    "github.com/golang-module/carbon/v2"
    "time"
)

// 类型转换
func ConvertToCarbonDateTime(value interface{}) (carbon.DateTime, error) {
    switch v := value.(type) {
    case carbon.DateTime:
        // 已经是 carbon.DateTime，则直接返回
        return v, nil
    case time.Time:
        // 将 time.Time 转换为 carbon.DateTime 类型
        return carbon.CreateFromStdTime(v).ToDateTimeStruct(), nil
    case string:
        // 将字符串转换为 carbon.DateTime 类型
        return carbon.Parse(v).ToDateTimeStruct(), nil
    default:
        // 无法转换则报错
        return carbon.DateTime{}, fmt.Errorf("type %T can not convert to carbon.DateTime", value)
    }
}

// 软删除时间类型，参考官方实现方式重写，直接复制使用即可
type CarbonNullTime struct {
    Time  carbon.DateTime
    Valid bool // Valid is true if Time is not NULL
}

// 实现 Scanner 接口
func (n *CarbonNullTime) Scan(value any) error {
    if value == nil {
        n.Time, n.Valid = carbon.DateTime{}, false
        return nil
    }

    // 转换类型
    t, err := ConvertToCarbonDateTime(value)

    // 转换失败
    if err != nil {
        common.Log.Error(err)
        n.Time, n.Valid = carbon.DateTime{}, false
        return err
    }

    // 转换成功
    n.Valid = true
    n.Time = t
    return nil
}

// 实现 Valuer 接口
func (n CarbonNullTime) Value() (driver.Value, error) {
    if !n.Valid {
        return nil, nil
    }
    return n.Time, nil
}
