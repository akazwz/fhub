package utils

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

var S3Storage = s3Util{}

type s3Util struct{}

var ctx = context.TODO()

func (s3Util s3Util) NewS3Client(key, secret, url string) (*s3.Client, error) {
	credentialsOptions := config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(key, secret, ""))
	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL: url,
		}, nil
	})
	endPointOptions := config.WithEndpointResolverWithOptions(customResolver)
	cfg, err := config.LoadDefaultConfig(ctx, credentialsOptions, endPointOptions)
	if err != nil {
		return nil, err
	}
	client := s3.NewFromConfig(cfg)
	return client, nil
}

func (s3Util s3Util) CreateMultipartUpload(client *s3.Client, bucket, key string) *s3.CreateMultipartUploadOutput {
	upload, err := client.CreateMultipartUpload(context.TODO(), &s3.CreateMultipartUploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		//ChecksumAlgorithm: types.ChecksumAlgorithmSha256,
	})
	if err != nil {
		return nil
	}
	return upload
}

func (s3Util s3Util) CreatePresignUploadPart(client *s3.Client, bucket, key, uploadId string, partNumber int32) *v4.PresignedHTTPRequest {
	psClient := s3.NewPresignClient(client, func(options *s3.PresignOptions) {
		options.Expires = 1 * time.Hour
	})

	presignUploadPart, err := psClient.PresignUploadPart(context.TODO(), &s3.UploadPartInput{
		Bucket:     aws.String(bucket),
		Key:        aws.String(key),
		PartNumber: partNumber,
		UploadId:   aws.String(uploadId),
	})
	if err != nil {
		return nil
	}
	return presignUploadPart
}

func (s3Util s3Util) GetPresignGetObjectURL(client *s3.Client, bucket, key, contentDisposition string) (*v4.PresignedHTTPRequest, error) {
	psClient := s3.NewPresignClient(client, func(options *s3.PresignOptions) {
		options.Expires = 1 * time.Hour
	})
	log.Println(bucket)
	object, err := psClient.PresignGetObject(context.TODO(), &s3.GetObjectInput{
		Bucket:                     aws.String(bucket),
		Key:                        aws.String(key),
		ResponseContentDisposition: aws.String(contentDisposition),
	})
	if err != nil {
		return nil, err
	}
	return object, nil
}

func (s3Util s3Util) CompleteUpload(client *s3.Client, bucket, key, uploadId, contentHash string) (*s3.CompleteMultipartUploadOutput, error) {
	parts := make([]types.CompletedPart, 0)

	listParts, err := client.ListParts(ctx, &s3.ListPartsInput{
		Bucket:   aws.String(bucket),
		Key:      aws.String(key),
		UploadId: aws.String(uploadId),
	})
	if err != nil {
		return nil, err
	}
	for _, part := range listParts.Parts {
		parts = append(parts, types.CompletedPart{
			ChecksumCRC32:  part.ChecksumCRC32,
			ChecksumCRC32C: part.ChecksumCRC32C,
			ChecksumSHA1:   part.ChecksumSHA1,
			ChecksumSHA256: part.ChecksumSHA256,
			ETag:           part.ETag,
			PartNumber:     part.PartNumber,
		})
	}

	complete, err := client.CompleteMultipartUpload(context.TODO(), &s3.CompleteMultipartUploadInput{
		Bucket:   aws.String(bucket),
		Key:      aws.String(key),
		UploadId: aws.String(uploadId),
		MultipartUpload: &types.CompletedMultipartUpload{
			Parts: parts,
		},
	})

	if err != nil {
		log.Println(err)
		return nil, err
	}
	return complete, nil
}

func (s3Util s3Util) FindUploadPart(client *s3.Client, bucket, key, uploadId string) (*s3.ListPartsOutput, error) {
	listMultipartUploads, err := client.ListParts(context.TODO(), &s3.ListPartsInput{
		Bucket:   aws.String(bucket),
		Key:      aws.String(key),
		UploadId: aws.String(uploadId),
	})
	if err != nil {
		return nil, err
	}
	return listMultipartUploads, nil
}
