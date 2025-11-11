package node

import "github.com/rahulSailesh-shah/ch8n_go/internal/constants"

type ExecutionResult struct {
	Status constants.ExecutionStatus `json:"status"`
	Data   map[string]any            `json:"data,omitempty"`
	Error  error                     `json:"error,omitempty"`
}

type WorkflowNode interface {
	Validate(params map[string]any) error
	Execute(params map[string]any) *ExecutionResult
}
