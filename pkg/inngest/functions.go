package inngest

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/inngest/inngestgo"
	"github.com/inngest/inngestgo/step"
	"github.com/rahulSailesh-shah/ch8n_go/internal/db/repo"
	"github.com/rahulSailesh-shah/ch8n_go/pkg/execution"
)

type ExecuteWorkflowRequest struct {
	WorkflowID uuid.UUID         `json:"workflow_id"`
	Nodes      []repo.Node       `json:"nodes"`
	Edges      []repo.Connection `json:"edges"`
}

func (i *Inngest) RegisterFunctions() error {
	err := i.executeWorkflow()
	return err
}

func (i *Inngest) ExecuteWorkflow(ctx context.Context, data *ExecuteWorkflowRequest) error {
	_, err := i.client.Send(ctx, inngestgo.Event{
		Name: "workflows/execute.workflow",
		Data: map[string]any{
			"nodes":       data.Nodes,
			"edges":       data.Edges,
			"workflow_id": data.WorkflowID,
		},
	})
	return err
}

func (i *Inngest) executeWorkflow() error {
	_, err := inngestgo.CreateFunction(
		i.client,
		inngestgo.FunctionOpts{
			ID:   "execute-workflow",
			Name: "Execute Workflow",
		},
		inngestgo.EventTrigger("workflows/execute.workflow", nil),
		func(ctx context.Context, input inngestgo.Input[ExecuteWorkflowRequest]) (any, error) {
			// Build and validate DAG
			nodes, err := step.Run(ctx, "build-dag", func(ctx context.Context) ([][]repo.Node, error) {
				dag := NewDAG(input.Event.Data.Nodes, input.Event.Data.Edges)
				if !dag.ValidateGraph() {
					return nil, fmt.Errorf("graph is not valid")
				}
				levelOrder := dag.GetLevelOrder()
				return levelOrder, nil
			})
			if err != nil {
				return nil, err
			}

			// Create execution context
			executionContext := execution.NewExecutionContext(input.Event.Data.WorkflowID)

			// Execute nodes in level order
			for _, level := range nodes {
				for _, node := range level {
					ec, err := step.Run(ctx, node.Type, func(ctx context.Context) (*execution.ExecutionContext, error) {
						return i.executeNode(executionContext, node)
					})
					if err != nil {
						return nil, err
					}
					executionContext = ec
				}
			}
			return executionContext, nil
		},
	)
	return err
}

func (i *Inngest) executeNode(
	executionContext *execution.ExecutionContext,
	node repo.Node,
) (*execution.ExecutionContext, error) {
	// Get node from registry
	workflowNode, err := i.nodeRegistry.Get(node.Type)
	if err != nil {
		return nil, inngestgo.NoRetryError(fmt.Errorf("failed to get node from registry: %w", err))
	}

	// Resolve params using the template engine
	var nodeParams map[string]any
	if node.Data != nil {
		if err := json.Unmarshal(node.Data, &nodeParams); err != nil {
			return nil, inngestgo.NoRetryError(fmt.Errorf("failed to unmarshal node params: %w", err))
		}
	} else {
		nodeParams = make(map[string]any)
	}

	resolvedParams, err := i.templateEngine.ResolveParams(nodeParams, executionContext)
	if err != nil {
		return nil, inngestgo.NoRetryError(fmt.Errorf("failed to resolve params: %w", err))
	}

	// Validate params
	if err := workflowNode.Validate(resolvedParams); err != nil {
		return nil, inngestgo.NoRetryError(fmt.Errorf("validation failed: %w", err))
	}

	// Execute node
	result, err := workflowNode.Execute(resolvedParams)
	if err != nil {
		return nil, fmt.Errorf("failed to execute node: %w", err)
	}

	// Store result
	executionContext.SetNodeOutput(node.Type, result)
	return executionContext, nil
}

// For testing purpose
func (i *Inngest) ExecuteDirect(data *ExecuteWorkflowRequest) (*execution.ExecutionContext, error) {
	dag := NewDAG(data.Nodes, data.Edges)
	if !dag.ValidateGraph() {
		return nil, fmt.Errorf("graph is not valid")
	}
	levelOrder := dag.GetLevelOrder()

	executionContext := execution.NewExecutionContext(data.WorkflowID)
	var err error

	for _, level := range levelOrder {
		for _, node := range level {
			executionContext, err = i.executeNode(executionContext, node)
			if err != nil {
				return nil, err
			}
		}
	}
	return executionContext, nil
}
