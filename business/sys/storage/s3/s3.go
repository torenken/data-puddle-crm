package storage

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"go.uber.org/zap"
)

//todo only for prototyping

func NewS3() *s3.Client {
	cfg, _ := config.LoadDefaultConfig(context.TODO())
	return s3.NewFromConfig(cfg)
}

func ListBuckets(ctx context.Context, log *zap.SugaredLogger, s3c *s3.Client) {
	log.Infof("calling")

	_, err := s3c.ListBuckets(ctx, &s3.ListBucketsInput{})
	if err != nil {
		log.Errorf("error: %v", err.Error())
	}

}
