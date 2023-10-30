package gedis

import "change-api/common"

// Slice / []interface{} 类型结果
type SliceResult struct {
    Result []interface{}
    Error  error
}

// 构造函数
func NewSliceResult(result []interface{}, err error) *SliceResult {
    return &SliceResult{Result: result, Error: err}
}

// 解析 []interface 结果
func (r *SliceResult) Unwrap() []interface{} {
    if r.Error != nil {
        common.Log.Debug("failed to query the redis cache")
        common.Log.Debug(r.Error)
    }
    return r.Result
}

// 解析 []interface 结果，如果报错，则使用默认值
func (r *SliceResult) UnwrapWithDefault(v []interface{}) []interface{} {
    if r.Error != nil {
        common.Log.Debug("failed to query the redis cache, use the default value:", v)
        common.Log.Debug(r.Error)
        return v
    }
    return r.Result
}

// 解析 []interface 结果，如果报错，则执行函数
func (r *SliceResult) UnwrapWithFunc(f func() []interface{}) []interface{} {
    if r.Error != nil {
        common.Log.Debug("failed to query the redis cache, , run funtion")
        common.Log.Debug(r.Error)
        return f()
    }
    return r.Result
}

// Iter 迭代方法，用法：
//
//  func demo() {
//      var conn = NewStringOperation()
//      var res = conn.Mget("name", "age", "gender").Iter()
//      for res.HasNext() {
//          fmt.Println(res.Next())
//      }
//  }
func (r *SliceResult) Iter() *Iterator {
    return NewIterator(r.Result)
}
