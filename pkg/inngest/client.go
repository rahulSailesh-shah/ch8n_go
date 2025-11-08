package inngest

import (
	"github.com/gin-gonic/gin"
	"github.com/inngest/inngestgo"
)

type Inngest struct {
	client inngestgo.Client
	XXX    string
}

func NewInngest() (*Inngest, error) {
	client, err := inngestgo.NewClient(inngestgo.ClientOpts{
		AppID: "core",
		Dev:   inngestgo.BoolPtr(true),
	})
	if err != nil {
		return nil, err
	}

	i := &Inngest{
		client: client,
		XXX:    "",
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
