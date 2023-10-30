package gedis

import (
    "change-api/common"
)

// 字符串类型结果
type StringResult struct {
    Result string
    Error  error
}

// 构造函数
func NewStringResult(result string, err error) *StringResult {
    return &StringResult{Result: result, Error: err}
}

// 解析字符串结果
func (r *StringResult) Unwrap() string {
    if r.Error != nil {
        common.Log.Debug("failed to query the redis cache")
        common.Log.Debug(r.Error)
    }
    return r.Result
}

// 解析字符串结果，如果报错，则使用默认值
func (r *StringResult) UnwrapWithDefault(v string) string {
    if r.Error != nil {
        common.Log.Debug("failed to query the redis cache, use the default value:", v)
        common.Log.Debug(r.Error)
        return v
    }
    return r.Result
}

// 解析字符串结果，如果报错，则执行函数
func (r *StringResult) UnwrapWithFunc(f func() string) string {
    if r.Error != nil {
        common.Log.Debug("failed to query the redis cache, run funtion")
        common.Log.Debug(r.Error)
        return f()
    }
    return r.Result
}
