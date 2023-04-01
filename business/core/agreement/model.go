package agreement

import (
	"time"

	"github.com/google/uuid"
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
