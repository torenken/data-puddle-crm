package agreement

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Storer interface {
	Create(ctx context.Context, agt Agreement) error
}

type Core struct {
	storer Storer
}

func NewCore(storer Storer) *Core {
	return &Core{
		storer: storer,
	}
}

func (c *Core) Create(ctx context.Context, na NewAgreement) (Agreement, error) {
	now := time.Now()

	agt := Agreement{
		ID:          uuid.UUID{},
		Name:        na.Name,
		DateCreated: now,
		DateUpdated: now,
	}

	if err := c.storer.Create(ctx, agt); err != nil {
		return Agreement{}, fmt.Errorf("create: %w", err)
	}

	return agt, nil
}
