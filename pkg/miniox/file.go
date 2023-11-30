package miniox

import "github.com/minio/minio-go/v7"

// 文件上传
func (m *Miniox) FileUpload(bucketName string, objectName string, filePath string, contextType string) (object minio.UploadInfo, err error) {
	object, err = m.Client.FPutObject(
		m.Context,
		bucketName,
		objectName,
		filePath,
		minio.PutObjectOptions{ContentType: contextType},
	)
	return
}

// 文件下载
func (m *Miniox) FileDownload(bucketName string, objectName string, filePath string) (err error) {
	err = m.Client.FGetObject(m.Context, bucketName, objectName, filePath, minio.GetObjectOptions{})
	return
}
