package dto

type NodeType string

const (
	NodeTypeInitial   NodeType = "INITIAL"
	NodeManualTrigger NodeType = "MANUAL_TRIGGER"
	NodeHTTPTrigger   NodeType = "HTTP_TRIGGER"
)
