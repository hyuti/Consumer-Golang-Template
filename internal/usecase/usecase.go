package usecase

import (
	"context"
	"github.com/hyuti/Consumer-Golang-Template/pkg/kafka"
	"github.com/hyuti/Consumer-Golang-Template/pkg/model"
	"golang.org/x/exp/slog"
)

type uc struct {
	log *slog.Logger
}

var _ kafka.Consumer[*model.Model] = (*uc)(nil)

func (u *uc) Consume(ctx context.Context, msg *model.Model) error {
	return nil
}

func NewUseCase(
	log *slog.Logger,
) kafka.Consumer[*model.Model] {
	return &uc{
		log: log,
	}
}
