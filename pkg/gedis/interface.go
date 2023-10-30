package gedis

import "change-api/common"

// interface{} 类型结果
type InterfaceResult struct {
    Result interface{}
    Error  error
}

// 构造函数
func NewInterfaceResult(result interface{}, err error) *InterfaceResult {
    return &InterfaceResult{Result: result, Error: err}
}

// 解析 interface 结果
func (r *InterfaceResult) Unwrap() interface{} {
    if r.Error != nil {
        common.Log.Debug("failed to query the redis cache")
        common.Log.Debug(r.Error)
    }
    return r.Result
}

// 解析 interface 结果，如果报错，则使用默认值
func (r *InterfaceResult) UnwrapWithDefault(v interface{}) interface{} {
    if r.Error != nil {
        common.Log.Debug("failed to query the redis cache, use the default value:", v)
        common.Log.Debug(r.Error)
        return v
    }
    return r.Result
}

// 解析 interface 结果，如果报错，则执行函数
func (r *InterfaceResult) UnwrapWithFunc(f func() interface{}) interface{} {
    if r.Error != nil {
        common.Log.Debug("failed to query the redis cache, , run funtion")
        common.Log.Debug(r.Error)
        return f()
    }
    return r.Result
}
