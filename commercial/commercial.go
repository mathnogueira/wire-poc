//go:build wireinject
// +build wireinject

package commercial

import (
	"github.com/google/wire"
	ht "github.com/mathnogueira/ioc/http"
	"github.com/mathnogueira/ioc/paidrepository"
	"github.com/mathnogueira/ioc/server"
)

func GetServer() (*server.Server, error) {
	wire.Build(server.NewServer, ht.NewController, paidrepository.NewUserRepository)
	return &server.Server{}, nil
}
