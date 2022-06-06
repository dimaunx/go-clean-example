package main

import (
	"context"

	"github.com/dimaunx/go-clean-example/pkg/controller"
	"github.com/dimaunx/go-clean-example/pkg/repository"
	"github.com/dimaunx/go-clean-example/pkg/router"
	"github.com/dimaunx/go-clean-example/pkg/usecase"
)

func main() {
	ctx := context.Background()
	repo := repository.NewMongoRepository()
	// Switch to redis
	// repo := repository.NewRedisRepository()
	deviceUseCase := usecase.NewDeviceUseCase(repo)
	deviceController := controller.NewDeviceController(ctx, deviceUseCase)
	httpRouter := router.NewRouter()

	httpRouter.GetDevice("/device", deviceController.FindAll)
	httpRouter.GetDevice("/device/:id", deviceController.FindById)
	httpRouter.AddNewDevice("/device", deviceController.AddDevice)
	httpRouter.Start(":8000")
}
