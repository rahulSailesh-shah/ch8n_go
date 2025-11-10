package inngest

import (
	"github.com/gin-gonic/gin"
	"github.com/inngest/inngestgo"
	"github.com/rahulSailesh-shah/ch8n_go/pkg/registry"
	"github.com/rahulSailesh-shah/ch8n_go/pkg/template"
)

type Inngest struct {
	client         inngestgo.Client
	templateEngine *template.TemplateEngine
	nodeRegistry   *registry.NodeRegistry
}

func NewInngest(
	nodeRegistry *registry.NodeRegistry,
	templateEngine *template.TemplateEngine,
) (*Inngest, error) {
	client, err := inngestgo.NewClient(inngestgo.ClientOpts{
		AppID: "core",
		Dev:   inngestgo.BoolPtr(true),
	})
	if err != nil {
		return nil, err
	}

	i := &Inngest{
		client:         client,
		templateEngine: templateEngine,
		nodeRegistry:   nodeRegistry,
	}

	err = i.RegisterFunctions()
	if err != nil {
		return nil, err
	}

	return i, nil
}

func (i *Inngest) Handler() gin.HandlerFunc {
	inngestHandler := i.client.Serve()
	return func(c *gin.Context) {
		inngestHandler.ServeHTTP(c.Writer, c.Request)
	}
}
