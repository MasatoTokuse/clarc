package server

import (
	"context"

	"github.com/go-chi/chi"
	"github.com/mtoku/clarc/app/controllers/user"
	"github.com/mtoku/clarc/app/infrastructure"
)

func getRouter() (chi.Router, error) {
	r := chi.NewRouter()
	userController, err := user.InitializeUserController(infrastructure.NewTestMySQLConnectionString(), context.TODO())
	if err != nil {
		return nil, err
	}
	r.Route("/user", func(r chi.Router) {
		r.Post("/create", userController.Create)
	})
	return r, nil
}
