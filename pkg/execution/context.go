package execution

import "github.com/google/uuid"

type ExecutionContext struct {
	WorkflowID  uuid.UUID
	ExecutionID uuid.UUID
	NodeOutputs map[string]map[string]any
	Variables   map[string]any
	Meta        map[string]any
}

func NewExecutionContext(workflowID uuid.UUID) *ExecutionContext {
	return &ExecutionContext{
		WorkflowID:  workflowID,
		ExecutionID: uuid.New(),
		NodeOutputs: make(map[string]map[string]any),
		Variables:   make(map[string]any),
		Meta:        make(map[string]any),
	}
}

func (ctx *ExecutionContext) SetNodeOutput(nodeID string, output map[string]any) {
	ctx.NodeOutputs[nodeID] = output
}

func (ctx *ExecutionContext) SetVariable(key string, value any) {
	ctx.Variables[key] = value
}

func (ctx *ExecutionContext) SetMeta(key string, value any) {
	ctx.Meta[key] = value
}

func (ctx *ExecutionContext) GetNodeOutput(nodeID string) (bool, map[string]any) {
	output, ok := ctx.NodeOutputs[nodeID]
	return ok, output
}

func (ctx *ExecutionContext) GetVariable(key string) (bool, any) {
	output, ok := ctx.Variables[key]
	return ok, output
}

func (ctx *ExecutionContext) GetMeta(key string) (bool, any) {
	output, ok := ctx.Meta[key]
	return ok, output
}

func (ctx *ExecutionContext) GetAllOutputs() (bool, map[string]map[string]any) {
	return true, ctx.NodeOutputs
}
