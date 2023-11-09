package dto

// 分页默认限制
var (
    DefaultMaxPageSize int64 = 100 // 默认最多显示 100 条数据
    DefaultPageSize    int64 = 1   // 默认每页显示 10 条数据
)

// 数据分页请求参数
type Pagination struct {
    NoPagination bool  `json:"no_pagination" form:"no_pagination"` // 默认分页
    PageNumber   int64 `json:"page_number" form:"page_number"`     // 请求页码
    PageSize     int64 `json:"page_size" form:"page_size"`         // 每页数据量
    Total        int64 `json:"total" form:"total"`                 // 数据总量
}

// 分页数据响应体
type PageResponse struct {
    Pagination `json:"pagination"`       // 分页信息
    List       interface{} `json:"list"` // 分页数据
}

// 通过 Limit 和 Offset 实现 Gorm 数据分页功能
func (p *Pagination) GetLimitAndOffset() (limit, offset int) {
    // 前端传递的数据接收到会是 int64 类型
    pageSize := p.PageSize
    pageNumber := p.PageNumber

    // 请求每页数量小于 1 或者大于配置最大允许数量都返回默认的配置
    if p.PageSize < 1 || p.PageSize > DefaultMaxPageSize {
        pageSize = DefaultPageSize
    }

    // 请求页码小于 1，则返回第一页
    if p.PageNumber < 1 {
        pageNumber = 1
    }

    // 统计最大页码，不能整除需要 +1，能整除则 -1
    maxPageNumber := p.Total/pageSize + 1
    if p.Total%pageSize == 0 {
        maxPageNumber -= 1
    }

    // 如果没数据，或者刚好只有一页
    if maxPageNumber <= 1 {
        pageNumber = 1
    }

    // 页码大于最大页码则显示最后一页
    if pageNumber > maxPageNumber {
        pageNumber = maxPageNumber
    }

    // 限制和偏移
    limit = int(pageSize)
    offset = int(pageSize * (pageNumber - 1))

    // 处理原始数据
    p.PageNumber = pageNumber
    p.PageSize = pageSize
    return
}
