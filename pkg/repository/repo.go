package repository

import (
	"context"

	"github.com/dimaunx/go-clean-example/pkg/entity"
)

type DeviceRepository interface {
	Save(ctx context.Context, d *entity.Device) (string, error)
	FindById(ctx context.Context, id string) (*entity.Device, error)
	FindAll(ctx context.Context) ([]entity.Device, error)
}
