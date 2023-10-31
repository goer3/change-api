package gedis

import (
    "change-api/common"
    "context"
    "fmt"
    "github.com/redis/go-redis/v9"
    "time"
)

// 参数处理
type OperationAttr struct {
    Name  string
    Value interface{}
}

// 构造函数
func NewOperationAttr(name string, value interface{}) *OperationAttr {
    return &OperationAttr{Name: name, Value: value}
}

// 多个参数
type OperationAttrs []*OperationAttr

// 查找参数
func (attrs OperationAttrs) Find(name string) *InterfaceResult {
    for _, attr := range attrs {
        if attr.Name == name {
            return NewInterfaceResult(attr.Value, nil)
        }
    }
    return NewInterfaceResult(nil, fmt.Errorf("operation attribute %s not found", name))
}

// 设置 Key 过期时间
func WithExpire(t time.Duration) *OperationAttr {
    return &OperationAttr{
        Name:  "expire",
        Value: t,
    }
}

// 设置 NX 锁，Key 不存在才能设置
func WithNX() *OperationAttr {
    return &OperationAttr{
        Name:  "nx",
        Value: struct{}{},
    }
}

// 设置 XX 锁，Key 存在才能设置
func WithXX() *OperationAttr {
    return &OperationAttr{
        Name:  "xx",
        Value: struct{}{},
    }
}

// 操作
type Operation struct {
    Redis   *redis.Client
    Context context.Context
}

// 构造函数
func NewOperation() *Operation {
    return &Operation{common.Cache, context.Background()}
}

// 获取单个字符串 Key
func (s *Operation) GetString(key string) *StringResult {
    return NewStringResult(s.Redis.Get(s.Context, key).Result())
}

// 获取单个 int Key
func (s *Operation) GetInt(key string) *IntResult {
    return NewIntResult(s.Redis.Get(s.Context, key).Int())
}

// 获取多个 Key
func (s *Operation) MGet(keys ...string) *SliceResult {
    return NewSliceResult(s.Redis.MGet(s.Context, keys...).Result())
}

// 删除单个 Key
func (s *Operation) Del(key string) (int64, error) {
    return s.Redis.Del(s.Context, key).Result()
}

// 设置 Key / Value
// 使用方法：
// gedis.Set("key", "value", gedis.WithExpire(time.Second * 10), gedis.WithNX())
func (s *Operation) Set(key string, value interface{}, attrs ...*OperationAttr) *InterfaceResult {
    // 生成参数列表
    oas := OperationAttrs(attrs)

    // 过期参数，如果没传递，默认设置永不不过期
    expire := oas.Find("expire").UnwrapWithDefault(time.Second * 0).(time.Duration)

    // 判断是否有 nx 锁
    nx := oas.Find("nx").UnwrapWithDefault(nil)
    if nx != nil {
        return NewInterfaceResult(s.Redis.SetNX(s.Context, key, value, expire).Result())
    }

    // 判断是否有 xx 锁
    xx := oas.Find("xx").UnwrapWithDefault(nil)
    if xx != nil {
        return NewInterfaceResult(s.Redis.SetXX(s.Context, key, value, expire).Result())
    }

    // 只设置超时
    return NewInterfaceResult(s.Redis.Set(s.Context, key, value, expire).Result())
}
