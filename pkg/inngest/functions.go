package inngest

import (
	"context"

	"github.com/inngest/inngestgo"
)

type userCreatedEvent struct {
	UserID string `json:"user_id"`
	Name   string `json:"name"`
}

func (i *Inngest) RegisterFunctions() error {
	err := i.registerWelcomeEmail()
	return err
}

func (i *Inngest) SendWelcomeEmail(ctx context.Context, event userCreatedEvent) error {
	_, err := i.client.Send(ctx, inngestgo.Event{
		Name: "user/created",
		Data: map[string]any{
			"user_id": event.UserID,
			"name":    event.Name,
		},
	})
	return err
}

func (i *Inngest) registerWelcomeEmail() error {
	_, err := inngestgo.CreateFunction(
		i.client,
		inngestgo.FunctionOpts{
			ID:   "send-welcome-email",
			Name: "Send Welcome Email",
		},
		inngestgo.EventTrigger("user.created", nil),
		func(ctx context.Context, input inngestgo.Input[userCreatedEvent]) (any, error) {
			return map[string]string{"status": input.Event.Data.Name}, nil
		},
	)
	return err
}
