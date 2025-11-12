package constants

type NodeType string

const (
	INITIAL         NodeType = "INITIAL"
	MANUAL_TRIGGER  NodeType = "MANUAL_TRIGGER"
	HTTP_TRIGGER    NodeType = "HTTP_TRIGGER"
	WEBHOOK_TRIGGER NodeType = "WEBHOOK_TRIGGER"
)

type ExecutionStatus string

const (
	ExecutionStatusPending ExecutionStatus = "pending"
	ExecutionStatusSuccess ExecutionStatus = "success"
	ExecutionStatusFailed  ExecutionStatus = "failed"
	ExecutionStatusSkipped ExecutionStatus = "skipped"
)
