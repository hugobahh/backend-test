package di

import (
	controllerRegister "backend-test/internal/app/controller"
	repositoryRegister "backend-test/internal/app/repository"
	serviceRegister "backend-test/internal/app/service"
	"backend-test/internal/health"
	"backend-test/pkg/logger"
	"sync"

	//Controller "test-api/internal/app/resume/controller"
	"backend-test/internal/config"
	"backend-test/internal/shutdown"
	"backend-test/pkg/http"

	//"backend-test/pkg/databases/mysql"

	"go.uber.org/dig"
)

var (
	container *dig.Container
	once      sync.Once
)

func GetContainer() *dig.Container {
	once.Do(func() {
		container = buildContainer()
	})
	return container
}

func buildContainer() *dig.Container {
	container := dig.New()
	container.Provide(logger.NewLog)
	container.Provide(config.NewConfiguration)
	//container.Provide(mysql.NewMySQLConnector)
	container.Provide(repositoryRegister.NewResumeRepository)
	container.Provide(controllerRegister.NewResumeController)
	container.Provide(serviceRegister.NewResumeService)

	container.Provide(health.NewHealthController)
	container.Provide(shutdown.NewShutdownManager)
	container.Provide(http.NewServer)

	return container
}
