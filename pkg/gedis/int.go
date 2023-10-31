package gedis

import "change-api/common"

// int 类型结果
type IntResult struct {
    Result int
    Error  error
}

// 构造函数
func NewIntResult(result int, err error) *IntResult {
    return &IntResult{Result: result, Error: err}
}

// 解析 int 结果
func (r *IntResult) Unwrap() int {
    if r.Error != nil {
        common.Log.Debug("failed to query the redis cache")
        common.Log.Debug(r.Error)
    }
    return r.Result
}

// 解析 int 结果，如果报错，则使用默认值
func (r *IntResult) UnwrapWithDefault(v int) int {
    if r.Error != nil {
        common.Log.Debug("failed to query the redis cache, use the default value:", v)
        common.Log.Debug(r.Error)
        return v
    }
    return r.Result
}

// 解析 int 结果，如果报错，则执行函数
func (r *IntResult) UnwrapWithFunc(f func() int) int {
    if r.Error != nil {
        common.Log.Debug("failed to query the redis cache, run function")
        common.Log.Debug(r.Error)
        return f()
    }
    return r.Result
}
