package miniox

import (
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/lifecycle"
)

// 判断 Bucket 是否存在
func (m *Miniox) BucketExists(bucketName string) (exists bool, err error) {
	exists, err = m.Client.BucketExists(m.Context, bucketName)
	return
}

// 创建 Bucket
func (m *Miniox) BucketCreate(bucketName string) (err error) {
	// 判断桶是否存在再去创建桶
	if exists, _ := m.Client.BucketExists(m.Context, bucketName); exists {
		return fmt.Errorf("BucketExistError")
	}
	err = m.Client.MakeBucket(m.Context, bucketName, minio.MakeBucketOptions{})
	return
}

// 设置 Bucket 的生命周期
func (m *Miniox) BucketLifecycleSet(bucketName string, days lifecycle.ExpirationDays) (err error) {
	config := lifecycle.NewConfiguration()
	config.Rules = []lifecycle.Rule{
		{
			ID:     "expire-bucket",
			Status: "Enabled",
			Expiration: lifecycle.Expiration{
				Days: days,
			},
		},
	}
	err = m.Client.SetBucketLifecycle(m.Context, bucketName, config)
	return
}

// 列出 Bucket
func (m *Miniox) BucketList() (buckets []minio.BucketInfo, err error) {
	buckets, err = m.Client.ListBuckets(m.Context)
	return
}

// 删除 Bucket
func (m *Miniox) BucketDelete(bucketName string) (err error) {
	// 判断桶是否存在再去删除桶
	if exists, _ := m.Client.BucketExists(m.Context, bucketName); !exists {
		return fmt.Errorf("BucketNotExistError")
	}
	err = m.Client.RemoveBucket(m.Context, bucketName)
	return
}

// 设置 bucket 默认可读
func (m *Miniox) BucketAnonymousReadonlySet(bucketName string) (err error) {
	// 策略规则，固定的格式
	policyJSON := fmt.Sprintf(`{
		"Version": "2012-10-17",
		"Statement": [
			{
				"Effect": "Allow",
				"Principal": "*",
				"Action": ["s3:GetObject"],
				"Resource": ["arn:aws:s3:::%s/*"]
			}
		]
	}`, bucketName)

	// 配置 bucket 策略，设置后页码上显示的 access 是 none，并不影响
	err = m.Client.SetBucketPolicy(m.Context, bucketName, policyJSON)
	return
}
