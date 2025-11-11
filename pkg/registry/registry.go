// pkg/node/registry/registry.go
package registry

import (
	"fmt"

	"github.com/rahulSailesh-shah/ch8n_go/internal/constants"
	"github.com/rahulSailesh-shah/ch8n_go/pkg/node"
)

type NodeFactory func() (node.WorkflowNode, error)

var factoryRegistry = make(map[constants.NodeType]NodeFactory)

func RegisterFactory(name constants.NodeType, factory NodeFactory) {
	factoryRegistry[name] = factory
}

type NodeRegistry struct{}

func NewNodeRegistry() *NodeRegistry {
	return &NodeRegistry{}
}

func (r *NodeRegistry) Get(name constants.NodeType) (node.WorkflowNode, error) {
	if factory, exists := factoryRegistry[name]; exists {
		return factory()
	}
	return nil, fmt.Errorf("node %s not found", name)
}
