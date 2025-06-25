package action

import "github.com/ragolsnagol/ragol-cli/core/context"

type Action struct {
	Action func(ctx context.Context) error
}

func NewAction(action func(ctx context.Context) error) Action {
	return Action{
		Action: action,
	}
}
