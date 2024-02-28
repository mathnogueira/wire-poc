//go:build wireinject
// +build wireinject

package startup

import (
	"github.com/google/wire"
	ht "github.com/mathnogueira/ioc/http"
	"github.com/mathnogueira/ioc/repository"
	"github.com/mathnogueira/ioc/server"
)

func GetServer() (*server.Server, error) {
	wire.Build(server.NewServer, ht.NewController, repository.NewUserRepository)
	return &server.Server{}, nil
}
