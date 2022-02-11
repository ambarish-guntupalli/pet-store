package main

import (
	"github.com/ambarish-guntupalli/pet-store/routes"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

const PetStoreIntegrationUrl = "petstore_url"

type PetStore struct {
	IntegrationUrl string
}

func main() {
	port := ":8080"

	vi := viper.New()
	vi.SetConfigFile("config.yaml")
	vi.ReadInConfig()

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	petStore := routes.PetStore{IntegrationUrl: vi.GetString(PetStoreIntegrationUrl)}
	r.Mount("/pets", petStore.Routes())

	log.Printf("Starting up on http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, r))
}
