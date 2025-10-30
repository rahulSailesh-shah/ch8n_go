package dto

import (
	"github.com/google/uuid"
)

type UpdateConnectionRequest struct {
	ID           string    `json:"id"`
	Source       uuid.UUID `json:"source"`
	SourceHandle string    `json:"sourceHandle"`
	Target       uuid.UUID `json:"target"`
	TargetHandle string    `json:"targetHandle"`
}
