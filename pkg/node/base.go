package node

type ExecutionStatus string

const (
	ExecutionStatusPending ExecutionStatus = "pending"
	ExecutionStatusSuccess ExecutionStatus = "success"
	ExecutionStatusFailed  ExecutionStatus = "failed"
	ExecutionStatusSkipped ExecutionStatus = "skipped"
)

type ExecutionResult struct {
	Status ExecutionStatus `json:"status"`
	Data   map[string]any  `json:"data,omitempty"`
	Error  error           `json:"error,omitempty"`
}

type WorkflowNode interface {
	GetName() string
	Validate(params map[string]any) error
	Execute(params map[string]any) *ExecutionResult
}
