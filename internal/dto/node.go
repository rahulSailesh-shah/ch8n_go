package dto

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/rahulSailesh-shah/ch8n_go/internal/constants"
)

type Position struct {
	X float64 `json:"x" binding:"required,min=0"`
	Y float64 `json:"y" binding:"required,min=0"`
}

type UpdateNodeRequest struct {
	ID       uuid.UUID          `json:"id" validate:"required"`
	Type     constants.NodeType `json:"type" validate:"required"`
	Name     string             `json:"name" validate:"required"`
	Position Position           `json:"position" binding:"required"`
	Data     json.RawMessage    `json:"data"`
}
