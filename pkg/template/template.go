package template

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"text/template"

	"github.com/rahulSailesh-shah/ch8n_go/pkg/execution"
)

/*
Usage Examples:

1. Access node output:
   "{{ .json.node1.data.user.name }}"
   "{{ index .json.node1.items 0 }}"

2. Access variables:
   "{{ .vars.api_key }}"
   "{{ .variables.base_url }}"

3. Access execution data:
   "{{ .execution.id }}"
   "{{ .execution.workflow_id }}"

4. Complex expressions:
   "{{ if eq .json.node1.status "success" }}active{{ else }}inactive{{ end }}"
   "{{ range .json.node1.items }}{{ .name }}{{ end }}"

5. Mixed content:
   "Bearer {{ .vars.api_token }}"
   "User: {{ .json.user.name }} ({{ .json.user.email }})"

6. Type preservation:
   "{{ .json.node1.count }}" -> returns number, not string
   "Total: {{ .json.node1.count }}" -> returns string "Total: 123"
*/

type TemplateEngine struct {
	funcMap template.FuncMap
}

func NewTemplateEngine() *TemplateEngine {
	return &TemplateEngine{
		funcMap: template.FuncMap{
			"json": jsonHelper,
			"toJson": func(v interface{}) string {
				b, _ := json.Marshal(v)
				return string(b)
			},
			"fromJson": func(s string) interface{} {
				var v interface{}
				json.Unmarshal([]byte(s), &v)
				return v
			},
		},
	}
}

func (t *TemplateEngine) ResolveParams(
	params map[string]any,
	globalCtx *execution.ExecutionContext,
) (map[string]any, error) {
	data := map[string]any{
		"json":      globalCtx.NodeOutputs,
		"vars":      globalCtx.Variables,
		"variables": globalCtx.Variables,
		"execution": map[string]any{
			"id":          globalCtx.ExecutionID,
			"workflow_id": globalCtx.WorkflowID,
		},
	}

	resolved := make(map[string]any)

	for key, value := range params {
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
		return t.resolveString(v, data)
	case map[string]any:
		return t.resolveMap(v, data)
	case []any:
		return t.resolveArray(v, data)
	default:
		return value, nil
	}
}

func (t *TemplateEngine) resolveString(str string, data map[string]any) (any, error) {
	if !strings.Contains(str, "{{") {
		return str, nil
	}

	trimmed := strings.TrimSpace(str)
	if strings.HasPrefix(trimmed, "{{") && strings.HasSuffix(trimmed, "}}") && strings.Count(trimmed, "{{") == 1 {
		return t.evaluateSingleExpression(trimmed, data)
	}

	tmpl, err := template.New("param").Funcs(t.funcMap).Parse(str)
	if err != nil {
		return nil, fmt.Errorf("template parse error: %w", err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return nil, fmt.Errorf("template execution error: %w", err)
	}

	return buf.String(), nil
}

func (t *TemplateEngine) evaluateSingleExpression(expr string, data map[string]any) (any, error) {
	wrapperTemplate := `{{- toJson (` + strings.TrimSpace(expr[2:len(expr)-2]) + `) -}}`

	tmpl, err := template.New("expr").Funcs(t.funcMap).Parse(wrapperTemplate)
	if err != nil {
		return nil, fmt.Errorf("template parse error: %w", err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return nil, fmt.Errorf("template execution error: %w", err)
	}

	var result any
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		return buf.String(), nil
	}

	return result, nil
}

func (t *TemplateEngine) resolveMap(m map[string]any, data map[string]any) (map[string]any, error) {
	resolved := make(map[string]any)
	for key, value := range m {
		resolvedValue, err := t.resolveValue(value, data)
		if err != nil {
			return nil, err
		}
		resolved[key] = resolvedValue
	}
	return resolved, nil
}

func (t *TemplateEngine) resolveArray(arr []any, data map[string]any) ([]any, error) {
	resolved := make([]any, len(arr))
	for i, value := range arr {
		resolvedValue, err := t.resolveValue(value, data)
		if err != nil {
			return nil, err
		}
		resolved[i] = resolvedValue
	}
	return resolved, nil
}

func jsonHelper(v any) string {
	b, _ := json.MarshalIndent(v, "", "  ")
	return string(b)
}
