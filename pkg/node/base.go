package node

type WorkflowNode interface {
	GetName() string
	Validate(params map[string]any) error
	Execute(params map[string]any) (map[string]any, error)
}
