package agreementbucket

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/torenken/data-puddle-mock-env/business/core/agreement"
	storage "github.com/torenken/data-puddle-mock-env/business/sys/storage/s3"
	"go.uber.org/zap"
)

type Store struct {
	log *zap.SugaredLogger
	s3c *s3.Client
}

func NewStore(log *zap.SugaredLogger, s3c *s3.Client) *Store {
	return &Store{
		log: log,
		s3c: s3c,
	}
}

func (s Store) Create(ctx context.Context, agt agreement.Agreement) error {
	//todo
	storage.ListBuckets(ctx, s.log, s.s3c)
	return nil
}
