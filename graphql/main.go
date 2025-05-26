package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/kelseyhightower/envconfig"
)

type AppConfig struct {
	AccountURL  string `envconfig:"ACCOUNT_SERVUCE_URL"`
	CatalogtURL string `envconfig:"CATALOG_SERVUCE_URL"`
	OrderURL    string `envconfig:"ORDER_SERVUCE_URL"`
}

func main() {
	var cfg AppConfig
	err := envconfig.Process("", &cfg)

	if err != nil {
		log.Fatalf("%v", err)
	}

	s, err := NewGraphQLServer(cfg.AccountURL, cfg.CatalogtURL, cfg.OrderURL)
	if err != nil {
		log.Fatalf("%v", err)
	}

	http.Handle("/graphql", handler.New(s.ToExecutableSchema()))
	http.Handle("playground", playground.Handler("mk", "/graphql"))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
