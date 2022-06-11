package controller

import (
	"context"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"

	"github.com/dimaunx/go-clean-example/pkg/entity"
	"github.com/dimaunx/go-clean-example/pkg/usecase"
)

type DeviceController interface {
	FindAll(c echo.Context) error
	FindById(c echo.Context) error
	AddDevice(c echo.Context) error
}

type Controller struct {
	deviceUseCase *usecase.DeviceUseCase
	ctx           context.Context
}

func NewDeviceController(ctx context.Context, u *usecase.DeviceUseCase) *Controller {
	return &Controller{deviceUseCase: u, ctx: ctx}
}

func (c Controller) FindAll(e echo.Context) error {
	devices, err := c.deviceUseCase.FindAll(c.ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if len(devices) > 0 {
		return e.JSON(http.StatusOK, devices)
	}
	return e.JSON(http.StatusOK, struct {
		Message string `json:"message"`
	}{
		Message: "No devices found",
	})
}

func (c Controller) FindById(e echo.Context) error {
	device, err := c.deviceUseCase.FindById(c.ctx, e.Param("id"))
	if err != nil {
		if strings.Contains(err.Error(), "mongo: no documents in result") {
			return echo.NewHTTPError(http.StatusNotFound, "Device not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusOK, device)
}

func (c Controller) AddDevice(e echo.Context) error {
	device := new(entity.Device)
	if err := e.Bind(device); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := e.Validate(device); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	id, err := c.deviceUseCase.Add(c.ctx, device)
	if err != nil {
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}
	return e.JSON(http.StatusOK, struct {
		Message string `json:"message"`
		Id      string `json:"id"`
	}{
		Message: "New device added successfully",
		Id:      id,
	})
}
