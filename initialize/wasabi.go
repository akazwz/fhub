package initialize

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/akazwz/fhub/global"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func InitWasabiClient() {
	accessKeyId := os.Getenv("WASABI_ACCESS_KEY")
	accessKeySecret := os.Getenv("WASABI_SECRET_KEY")

	// 生成 s3 client
	client, err := generateWasabiClient(accessKeyId, accessKeySecret)
	if err != nil {
		log.Fatalln("init wasabi client error:", err)
	}
	global.WasabiClient = client
}

func generateWasabiClient(accessKeyId, accessKeySecret string) (*s3.Client, error) {
	resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL: fmt.Sprintf("https://s3.eu-west-2.wasabisys.com"),
		}, nil
	})

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithEndpointResolverWithOptions(resolver),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKeyId, accessKeySecret, "")),
	)
	if err != nil {
		return nil, err
	}

	client := s3.NewFromConfig(cfg)
	return client, nil
}
