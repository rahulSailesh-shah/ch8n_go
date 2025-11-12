package http_node

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/rahulSailesh-shah/ch8n_go/internal/constants"
	"github.com/rahulSailesh-shah/ch8n_go/pkg/node"
	"github.com/rahulSailesh-shah/ch8n_go/pkg/registry"
)

type HTTPRequestNode struct {
	client *http.Client
}

func init() {
	registry.RegisterFactory(constants.HTTP_TRIGGER, func() (node.WorkflowNode, error) {
		return NewHTTPRequestNode()
	})
}

func NewHTTPRequestNode() (*HTTPRequestNode, error) {
	return &HTTPRequestNode{
		client: &http.Client{
			Timeout: time.Second * 30,
		},
	}, nil
}

func (n *HTTPRequestNode) Validate(params map[string]any) error {
	if _, ok := params["variableName"]; !ok {
		return fmt.Errorf("variable name is required")
	}
	if _, ok := params["endpoint"]; !ok {
		return fmt.Errorf("endpoint is required")
	}

	endpoint, ok := params["endpoint"].(string)
	if !ok {
		return fmt.Errorf("endpoint must be a string")
	}

	_, err := url.ParseRequestURI(endpoint)
	if err != nil {
		return fmt.Errorf("endpoint must be a valid URL: %w", err)
	}

	if _, ok := params["method"]; !ok {
		return fmt.Errorf("method is required")
	}

	return nil
}

func (n *HTTPRequestNode) Execute(params map[string]any) *node.ExecutionResult {
	url := params["endpoint"].(string)
	method := params["method"].(string)

	// Set body
	var bodyReader io.Reader
	if body, ok := params["body"]; ok {
		switch v := body.(type) {
		case string:
			bodyReader = bytes.NewBufferString(v)
		case map[string]any, any:
			jsonBody, err := json.Marshal(v)
			if err != nil {
				return &node.ExecutionResult{
					Status: constants.ExecutionStatusFailed,
					Error:  err,
				}
			}
			bodyReader = bytes.NewBuffer(jsonBody)
		}
	}

	// Set body if required
	requiresBody := []string{"POST", "PUT", "PATCH"}
	if bodyReader == nil && slices.Contains(requiresBody, method) {
		bodyReader = bytes.NewReader([]byte{})
	}

	// Create request
	req, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		return &node.ExecutionResult{
			Status: constants.ExecutionStatusFailed,
			Error:  err,
		}
	}

	// Set default headers and custom headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	if _, ok := params["headers"]; ok {
		if headers, ok := params["headers"].(map[string]any); ok {
			for k, v := range headers {
				req.Header.Set(k, v.(string))
			}
		}
	}

	// Make request
	resp, err := n.client.Do(req)
	if err != nil {
		return &node.ExecutionResult{
			Status: constants.ExecutionStatusFailed,
			Error:  err,
		}
	}
	defer resp.Body.Close()

	// Read response
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &node.ExecutionResult{
			Status: constants.ExecutionStatusFailed,
			Error:  err,
		}
	}

	// Parse response and return as JSON if possible
	var jsonResp map[string]any
	if err := json.Unmarshal(respBody, &jsonResp); err == nil {
		return &node.ExecutionResult{
			Status: constants.ExecutionStatusSuccess,
			Data: map[string]any{
				"status":  resp.StatusCode,
				"data":    jsonResp,
				"headers": resp.Header,
			},
		}
	}

	// Return response as string if not JSON
	return &node.ExecutionResult{
		Status: constants.ExecutionStatusSuccess,
		Data: map[string]any{
			"status":  resp.StatusCode,
			"data":    string(respBody),
			"headers": resp.Header,
		},
	}
}
