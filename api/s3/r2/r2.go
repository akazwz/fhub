package r2

import (
	"os"
	"time"

	"github.com/akazwz/fhub/global"
	"github.com/akazwz/fhub/model/response"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
)

func ListBuckets(c *gin.Context) {
	buckets, err := global.R2Client.ListBuckets(c, &s3.ListBucketsInput{})
	if err != nil {
		response.BadRequest(400, nil, "获取 buckets 失败", c)
		return
	}
	response.Ok(200, buckets, "success", c)
}

func GetObjectURL(c *gin.Context) {
	key := c.Param("key")

	_, err := global.R2Client.HeadObject(c, &s3.HeadObjectInput{
		Bucket: aws.String(os.Getenv("R2_BUCKET_NAME")),
		Key:    aws.String(key),
	})
	if err != nil {
		response.BadRequest(400, nil, err.Error(), c)
		return
	}
	psClient := s3.NewPresignClient(global.R2Client, func(options *s3.PresignOptions) {
		options.Expires = 1 * time.Hour
	})

	object, err := psClient.PresignGetObject(c, &s3.GetObjectInput{
		Bucket: aws.String(os.Getenv("R2_BUCKET_NAME")),
		Key:    aws.String(key),
	})
	if err != nil {
		response.BadRequest(400, nil, err.Error(), c)
		return
	}
	response.Ok(200, object, "success", c)
}

func GetUploadURL(c *gin.Context) {
	psClient := s3.NewPresignClient(global.R2Client, func(options *s3.PresignOptions) {
		options.Expires = 1 * time.Hour
	})

	putObject, err := psClient.PresignPutObject(c, &s3.PutObjectInput{
		Bucket: aws.String(os.Getenv("R2_BUCKET_NAME")),
		Key:    aws.String("dog.png"),
	})
	if err != nil {
		response.BadRequest(400, nil, err.Error(), c)
		return
	}
	response.Ok(200, putObject, "success", c)
}
