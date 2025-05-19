package port

import "context"

type TemplateContract interface {
	Temp(ctx context.Context, input any) (any, error)
}
