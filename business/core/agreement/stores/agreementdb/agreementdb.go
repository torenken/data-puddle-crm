package agreementdb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/torenken/data-puddle-mock-env/business/core/agreement"
	"go.uber.org/zap"
)

//todo create own interface for selected dynamo.Client methods

type Storer interface {
	PutItem(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error)
}

// Store manages the set of APIs for product database access.
type Store struct {
	log *zap.SugaredLogger
	db  *Storer
}

// NewStore constructs the api for data access.
func NewStore(log *zap.SugaredLogger, db *Storer) *Store {
	return &Store{
		log: log,
		db:  db,
	}
}

func (s Store) Create(ctx context.Context, agt agreement.Agreement) {
}
