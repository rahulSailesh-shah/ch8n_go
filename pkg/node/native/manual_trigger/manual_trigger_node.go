package manual_trigger_node

import (
	"github.com/rahulSailesh-shah/ch8n_go/internal/constants"
	"github.com/rahulSailesh-shah/ch8n_go/pkg/node"
	"github.com/rahulSailesh-shah/ch8n_go/pkg/registry"
)

type ManualTriggerNode struct{}

func init() {
	registry.RegisterFactory(constants.MANUAL_TRIGGER, func() (node.WorkflowNode, error) {
		return NewManualTriggerNode()
	})
}

func NewManualTriggerNode() (*ManualTriggerNode, error) {
	return &ManualTriggerNode{}, nil
}

func (n *ManualTriggerNode) Validate(params map[string]any) error {
	return nil
}

func (n *ManualTriggerNode) Execute(params map[string]any) *node.ExecutionResult {
	return &node.ExecutionResult{
		Status: constants.ExecutionStatusSuccess,
	}
}
