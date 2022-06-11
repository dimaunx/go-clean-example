package usecase

import (
	"context"

	"github.com/google/uuid"

	"github.com/dimaunx/go-clean-example/pkg/entity"
	"github.com/dimaunx/go-clean-example/pkg/repository"
)

type DeviceService interface {
	Add(ctx context.Context, d *entity.Device) (string, error)
	FindAll(ctx context.Context) ([]entity.Device, error)
	FindById(ctx context.Context, id string) (*entity.Device, error)
}

type DeviceUseCase struct {
	repo repository.DeviceRepository
}

func NewDeviceUseCase(r repository.DeviceRepository) *DeviceUseCase {
	return &DeviceUseCase{repo: r}
}

func (s DeviceUseCase) Add(ctx context.Context, d *entity.Device) (string, error) {
	d.Id = uuid.NewString()
	id, err := s.repo.Save(ctx, d)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (s DeviceUseCase) FindAll(ctx context.Context) ([]entity.Device, error) {
	data, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s DeviceUseCase) FindById(ctx context.Context, id string) (*entity.Device, error) {
	data, err := s.repo.FindById(ctx, id)
	if err != nil {
		return nil, err
	}
	return data, nil
}
