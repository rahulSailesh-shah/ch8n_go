package http_node

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/rahulSailesh-shah/ch8n_go/pkg/node"
	"github.com/rahulSailesh-shah/ch8n_go/pkg/registry"
)

type HTTPRequestNode struct {
	client *http.Client
}

func init() {
	registry.RegisterFactory("http_request", func() (node.WorkflowNode, error) {
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

func (n *HTTPRequestNode) GetName() string {
	return "HTTP_TRIGGER"
}

func (n *HTTPRequestNode) Validate(params map[string]any) error {
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

func (n *HTTPRequestNode) Execute(params map[string]any) (map[string]any, error) {
	res := make(map[string]any)
	res["endpoint"] = params["endpoint"]
	res["method"] = params["method"]
	res["headers"] = params["headers"]
	res["body"] = params["body"]

	return res, nil
}
