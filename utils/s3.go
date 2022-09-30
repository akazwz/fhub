package utils

import (
	"context"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

var S3Storage = s3Util{}

type s3Util struct{}

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

	listMultipartUploads, err := client.ListParts(context.TODO(), &s3.ListPartsInput{
		Bucket:   aws.String(bucket),
		Key:      aws.String(key),
		UploadId: aws.String(uploadId),
	})

	for _, part := range listMultipartUploads.Parts {
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
		Bucket:         aws.String(bucket),
		Key:            aws.String(key),
		UploadId:       aws.String(uploadId),
		ChecksumSHA256: aws.String(contentHash),
		MultipartUpload: &types.CompletedMultipartUpload{
			Parts: parts,
		},
	})

	if err != nil {
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
