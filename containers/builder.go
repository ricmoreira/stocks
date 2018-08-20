// Inspiration to create dependency injection came from this post: https://blog.drewolson.org/dependency-injection-in-go/

package containers

import (
	"stocks/config"
	controllers "stocks/controllers/v1"
	"stocks/handlers"
	"stocks/repositories"
	"stocks/server"
	"stocks/services"

	"go.uber.org/dig"
)

// BuildContainer returns a container with all app dependencies built in
func BuildContainer() *dig.Container {
	container := dig.New()

	// config
	err := container.Provide(config.NewConfig)
	if err != nil {
		panic(err)
	}

	// persistance layer
	err = container.Provide(repositories.NewDBCollections)
	if err != nil {
		panic(err)
	}
	err = container.Provide(repositories.NewStockMovRepository)
	if err != nil {
		panic(err)
	}

	// services
	err = container.Provide(services.NewStockMovService)
	if err != nil {
		panic(err)
	}
	err = container.Provide(services.NewKafkaConsumer)
	if err != nil {
		panic(err)
	}

	// controllers
	err = container.Provide(controllers.NewStockMovController)
	if err != nil {
		panic(err)
	}

	// generic http layer
	err = container.Provide(handlers.NewHttpHandlers)
	if err != nil {
		panic(err)
	}

	// server
	err = container.Provide(server.NewServer)
	if err != nil {
		panic(err)
	}

	return container
}
