package agreementdb

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/torenken/data-puddle-crm/business/core/agreement"
	"go.uber.org/zap"
)

type Store struct {
	log *zap.SugaredLogger
	db  *dynamodb.Client
}

func NewStore(log *zap.SugaredLogger, db *dynamodb.Client) *Store {
	return &Store{
		log: log,
		db:  db,
	}
}

func (s *Store) Create(ctx context.Context, agreement agreement.Agreement) error {
	return nil
}
