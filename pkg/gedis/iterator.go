package gedis

// 迭代器
type Iterator struct {
    Data  []interface{}
    Index int
}

// 构造函数
func NewIterator(data []interface{}) *Iterator {
    return &Iterator{Data: data}
}

// 判断是否还有下一个
func (i *Iterator) HasNext() bool {
    if i.Data == nil || len(i.Data) == 0 {
        return false
    }
    return true
}

// 获取下一个
func (i *Iterator) Next() (ret interface{}) {
    ret = i.Data[i.Index]
    i.Index += 1
    return ret
}
