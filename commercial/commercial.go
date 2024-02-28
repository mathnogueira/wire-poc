//go:build wireinject
// +build wireinject

package commercial

import (
	"github.com/google/wire"
	"github.com/mathnogueira/ioc/paidhttp"
	"github.com/mathnogueira/ioc/paidrepository"
	"github.com/mathnogueira/ioc/server"
)

func GetServer() (*server.Server, error) {
	wire.Build(server.NewServer, paidhttp.NewController, paidrepository.NewUserRepository)
	return &server.Server{}, nil
}
