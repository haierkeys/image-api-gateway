package cloudflare_r2

import (
	"bytes"
	"context"
	"fmt"
	"mime/multipart"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/pkg/errors"

	"github.com/haierspi/golang-image-upload-service/global"
)

type R2 struct {
	S3Client  *s3.Client
	S3Manager *manager.Uploader
}

func (p *R2) GetBucket(bucketName string) string {

	// Get bucket
	if len(bucketName) <= 0 {
		bucketName = global.Config.CloudfluR2.BucketName
	}

	return bucketName
}

// UploadByFile 上传文件
func (p *R2) SendFile(fileKey string, file multipart.File, h *multipart.FileHeader) (string, error) {

	ctx := context.Background()
	bucket := p.GetBucket("")

	if strings.HasSuffix(global.Config.CloudfluR2.CustomPath, "/") {
		fileKey = global.Config.CloudfluR2.CustomPath + fileKey
	} else {
		fileKey = global.Config.CloudfluR2.CustomPath + "/" + fileKey
	}

	k, _ := h.Open()

	_, err := p.S3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(fileKey),
		Body:        k,
		ContentType: aws.String(h.Header.Get("Content-Type")),
	})

	if err != nil {
		return "", err
	}

	return fileKey, nil
}

func (p *R2) SendContent(fileKey string, content []byte) (string, error) {

	ctx := context.Background()
	bucket := p.GetBucket("")

	if strings.HasSuffix(global.Config.CloudfluR2.CustomPath, "/") {
		fileKey = global.Config.CloudfluR2.CustomPath + fileKey
	} else {
		fileKey = global.Config.CloudfluR2.CustomPath + "/" + fileKey
	}

	input := &s3.PutObjectInput{
		Bucket:            aws.String(bucket),
		Key:               aws.String(fileKey),
		Body:              bytes.NewReader(content),
		ChecksumAlgorithm: types.ChecksumAlgorithmSha256,
	}
	output, err := p.S3Manager.Upload(ctx, input)
	if err != nil {
		var noBucket *types.NoSuchBucket
		if errors.As(err, &noBucket) {
			fmt.Printf("Bucket %s does not exist.\n", bucket)
			err = noBucket
		}
	} else {
		err := s3.NewObjectExistsWaiter(p.S3Client).Wait(ctx, &s3.HeadObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(fileKey),
		}, time.Minute)
		if err != nil {
			fmt.Printf("Failed attempt to wait for object %s to exist in %s.\n", fileKey, bucket)
		} else {
			_ = *output.Key
		}
	}

	return fileKey, err
}