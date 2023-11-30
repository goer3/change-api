package initialize

import (
	"change-api/common"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// 初始化 Minio 连接
func Minio() {
	// 创建客户端连接
	client, err := minio.New(common.Config.Minio.URL, &minio.Options{
		Secure: common.Config.Minio.SSL,
		Creds:  credentials.NewStaticV4(common.Config.Minio.AccessKey, common.Config.Minio.AccessSecret, ""),
	})

	if err != nil {
		common.Log.Error("minio connect failed")
		common.Log.Error(err)
		panic("minio connect failed")
	}

	common.MinioClient = client
}

// 初始化 Minio Bucket
func MinioBucket() {

}
