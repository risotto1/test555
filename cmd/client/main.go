package main

import (
	"net/http"

	"github.com/GoingFast/test6/internal/client"
	"github.com/GoingFast/test6/pkg/env"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Heartbeat("/healthz"))
	svc := client.NewService()
	r.Get("/read", svc.Read())
	http.ListenAndServe(env.FallbackEnv("GATEWAY_LISTEN_ADDR", ":8081"), r)
}
