// pkg/node/registry/registry.go
package registry

import (
	"fmt"

	"github.com/rahulSailesh-shah/ch8n_go/pkg/node"
)

type NodeFactory func() (node.WorkflowNode, error)

var factoryRegistry = make(map[string]NodeFactory)

func RegisterFactory(name string, factory NodeFactory) {
	factoryRegistry[name] = factory
}

func InitializeAllNodes(registry *NodeRegistry) error {
	for name, factory := range factoryRegistry {
		node, err := factory()
		if err != nil {
			return fmt.Errorf("failed to create node %s: %w", name, err)
		}
		if err := registry.Register(node); err != nil {
			return fmt.Errorf("failed to register node %s: %w", name, err)
		}
	}

	return nil
}

type NodeRegistry struct {
	nodes map[string]node.WorkflowNode
}

func NewNodeRegistry() *NodeRegistry {
	return &NodeRegistry{
		nodes: make(map[string]node.WorkflowNode),
	}
}

func (r *NodeRegistry) Register(node node.WorkflowNode) error {
	name := node.GetName()
	if _, exists := r.nodes[name]; exists {
		return fmt.Errorf("node %s already registered", name)
	}
	r.nodes[name] = node
	return nil
}

func (r *NodeRegistry) Get(name string) (node.WorkflowNode, error) {
	node, exists := r.nodes[name]
	if !exists {
		return nil, fmt.Errorf("node %s not found", name)
	}
	return node, nil
}
