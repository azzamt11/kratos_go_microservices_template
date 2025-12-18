//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"kratos_go_microservices_template/app/user/internal/biz"
	"kratos_go_microservices_template/app/user/internal/conf"
	"kratos_go_microservices_template/app/user/internal/data"
	"kratos_go_microservices_template/app/user/internal/server"
	"kratos_go_microservices_template/app/user/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
