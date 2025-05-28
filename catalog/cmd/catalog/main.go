package main

import (
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/murodkadirkhanoff/go-grpc-graphql-microservice/catalog"
)

type Config struct {
	DatabaseURL string `envconfig:"DATABASE_URL"`
}

func main() {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	var r catalog.Repository

	for {
		rNew, err := catalog.NewElasticRepository(cfg.DatabaseURL)
		if err != nil {
			log.Println("Ошибка подключения к базе данных:", err)
			time.Sleep(2 * time.Second)
			continue
		}

		r = rNew // только после успешного подключения
		break
	}

	defer r.Close()
	log.Println(":istening on port 8080...")
	s := catalog.NewService(r)
	log.Fatal(catalog.ListenGRPC(s, 8080))
}
