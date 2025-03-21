package app

import (
	v1 "unhashService/app/controller/http/v1"
	"unhashService/app/usecase/hasher"
	"unhashService/entity"
	"unhashService/pkg/logger"
)

type App struct {
	web        *v1.HashController
	logger     *logger.Logger
	HTTPConfig entity.HTTP
}

func New(cfg *entity.Config) (*App, error) {
	lg := logger.NewConsoleLogger(logger.DebugLevel)
	secret := cfg.GetHMACSecret()
	usecase := hasher.New(lg, secret)

	web := v1.NewHashController(usecase, lg)

	return &App{
		web:        web,
		logger:     lg,
		HTTPConfig: cfg.HTTP,
	}, nil
}

func (a *App) Run() error {
	a.logger.Info("Application running")
	mux := a.web.SetupRoutes()
	endpoint := a.HTTPConfig.Host + ":" + a.HTTPConfig.Port
	err := a.web.ListenAndServe(endpoint, mux)
	if err != nil {
		a.logger.Error(err.Error())
	}
	return err
}
