package main

import (
	"github.com/ruifrodrigues/ecooda/config"
)

func main() {
	conf, err := config.NewConfig(
		config.WithMySQL,
		config.WithCache,
		config.WithQueue,
		config.WithLogger,
		config.WithElastic,
	)

	if err != nil {
		panic("Initialize configuration failed! >> " + err.Error())
	}

	//config.RunFactories(conf.Database, challenge.Seeders)

	go runGrpcServer(conf)
	runHttpProxy(conf)
}
