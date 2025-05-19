package service

import (
	"context"
	"pongsatorn/go-http-template/internal/core/port"
)

type TemplateService struct {
	tplContract port.TemplateContract
}

func NewTemplateService(tplContract port.TemplateContract) *TemplateService {
	return &TemplateService{
		tplContract: tplContract,
	}
}

func (s *TemplateService) Temp(ctx context.Context, input any) (any, error) {
	return s.tplContract.Temp(ctx, input)
}
