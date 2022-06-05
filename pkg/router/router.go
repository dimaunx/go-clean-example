package router

import (
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var apikey = os.Getenv("API_KEY")

type router struct {
	router *echo.Echo
}

type PostValidator struct {
	validator *validator.Validate
}

type Router interface {
	GetDevice(path string, f func(c echo.Context) error)
	AddNewDevice(path string, f func(c echo.Context) error)
	Start(address string)
}

func NewRouter() Router {
	return &router{router: echo.New()}
}

func (v PostValidator) Validate(i interface{}) error {
	if err := v.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func (r router) GetDevice(path string, f func(c echo.Context) error) {
	r.router.GET(path, f)
}

func (r router) AddNewDevice(path string, f func(c echo.Context) error) {
	r.router.Validator = &PostValidator{validator: validator.New()}
	r.router.POST(path, f)
}

func (r router) Start(address string) {
	r.router.Use(middleware.Logger())
	r.router.Use(middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		KeyLookup: "header:X-Api-Key",
		Validator: func(key string, c echo.Context) (bool, error) {
			return key == apikey, nil
		},
	}))
	if err := r.router.Start(address); err != http.ErrServerClosed {
		r.router.Logger.Fatal(err)
	}
}
