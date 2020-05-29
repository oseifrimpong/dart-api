package dig

import (
	"dart-api/api/config"
	"dart-api/api/repository"
	"dart-api/api/service"

	"go.uber.org/dig"
)

var container = dig.New()

func BuildProject() *dig.Container {
	// container.Provide()
	container.Provide(config.NewDb)
	container.Provide(config.NewConfig)
	container.Provide(repository.NewUserRepo)
	container.Provide(service.NewUserService)

	return container
}

func Invoke(i interface{}) error {
	return container.Invoke(i)
}
