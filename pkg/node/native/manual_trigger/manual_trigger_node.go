package http_node

import (
	"github.com/rahulSailesh-shah/ch8n_go/pkg/node"
	"github.com/rahulSailesh-shah/ch8n_go/pkg/registry"
)

type ManualTriggerNode struct{}

func init() {
	registry.RegisterFactory("manual_trigger", func() (node.WorkflowNode, error) {
		return NewManualTriggerNode()
	})
}

func NewManualTriggerNode() (*ManualTriggerNode, error) {
	return &ManualTriggerNode{}, nil
}

func (n *ManualTriggerNode) GetName() string {
	return "MANUAL_TRIGGER"
}

func (n *ManualTriggerNode) Validate(params map[string]any) error {
	return nil
}

func (n *ManualTriggerNode) Execute(params map[string]any) (map[string]any, error) {
	res := make(map[string]any)
	res["data"] = "Node Triggered Manually"
	res["test data"] = map[string]any{
		"var": "https://rahulshah.com",
	}

	return res, nil
}
