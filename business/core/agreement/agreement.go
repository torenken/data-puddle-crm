package agreement

import "context"

type Storer interface {
	Create(ctx context.Context, agreement Agreement) error
}

type Core struct {
	storer Storer
}

func NewCore(storer Storer) *Core {
	return &Core{
		storer: storer,
	}
}

func (c Core) Create(ctx context.Context, agreement Agreement) error {
	return nil //todo
}
