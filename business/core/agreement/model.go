package agreement

import (
	"github.com/google/uuid"
	"time"
)

type Agreement struct {
	ID          uuid.UUID
	Name        string
	DateCreated time.Time
	DateUpdated time.Time
}

type NewAgreement struct {
	Name string
}
