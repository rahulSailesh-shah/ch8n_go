package template

import (
	"encoding/json"
	"fmt"

	"github.com/expr-lang/expr"
	"github.com/rahulSailesh-shah/ch8n_go/pkg/execution"
)

type TemplateEngine struct{}

func NewTemplateEngine() *TemplateEngine {
	return &TemplateEngine{}
}

// ResolveParams evaluates all template strings in the params map using expr
func (t *TemplateEngine) ResolveParams(
	params map[string]any,
	globalCtx *execution.ExecutionContext,
) (map[string]any, error) {
	data := map[string]any{
		"$json": globalCtx.NodeOutputs,
		"$vars": globalCtx.Variables,
		"$execution": map[string]any{
			"id":          globalCtx.ExecutionID,
			"workflow_id": globalCtx.WorkflowID,
		},
	}

	resolved := make(map[string]any)

	for key, value := range params {
		// Special handling for "body" - try to parse as JSON first
		if key == "body" {
			if strVal, ok := value.(string); ok {
				var bodyObj any
				if err := json.Unmarshal([]byte(strVal), &bodyObj); err == nil {
					// Successfully parsed as JSON, resolve the parsed structure
					resolvedValue, err := t.resolveValue(bodyObj, data)
					if err != nil {
						return nil, fmt.Errorf("failed to resolve param %s: %w", key, err)
					}
					resolved[key] = resolvedValue
					continue
				}
			}
		}

		// Standard resolution for non-body params or unparseable body
		resolvedValue, err := t.resolveValue(value, data)
		if err != nil {
			return nil, fmt.Errorf("failed to resolve param %s: %w", key, err)
		}
		resolved[key] = resolvedValue
	}

	return resolved, nil
}

func (t *TemplateEngine) resolveValue(value any, data map[string]any) (any, error) {
	switch v := value.(type) {
	case string:
		return t.evaluateExpr(v, data)
	case map[string]any:
		m := make(map[string]any)
		for k, val := range v {
			resolved, err := t.resolveValue(val, data)
			if err != nil {
				return nil, err
			}
			m[k] = resolved
		}
		return m, nil
	case []any:
		arr := make([]any, len(v))
		for i, val := range v {
			resolved, err := t.resolveValue(val, data)
			if err != nil {
				return nil, err
			}
			arr[i] = resolved
		}
		return arr, nil
	default:
		return v, nil
	}
}

// evaluateExpr runs expr on the string
func (t *TemplateEngine) evaluateExpr(str string, data map[string]any) (any, error) {
	trimmed := str
	if len(trimmed) >= 4 && trimmed[:2] == "{{" && trimmed[len(trimmed)-2:] == "}}" {
		exprStr := trimmed[2 : len(trimmed)-2] // remove {{ }}
		prog, err := expr.Compile(exprStr, expr.Env(data))
		if err != nil {
			return nil, fmt.Errorf("compile error: %w", err)
		}
		result, err := expr.Run(prog, data)
		if err != nil {
			return nil, fmt.Errorf("evaluation error: %w", err)
		}
		return result, nil
	}
	return str, nil
}
