package miniox

import (
	"change-api/common"
	"context"
	"github.com/minio/minio-go/v7"
)

// 客户端方法封装
type Miniox struct {
	Client  *minio.Client
	Context context.Context
}

// 构造函数
func NewMiniox() *Miniox {
	return &Miniox{Client: common.MinioClient, Context: context.Background()}
}
