//go:build wireinject

package main

import (
	"github.com/blackhorseya/patrick/cmd"
	"github.com/blackhorseya/patrick/internal/app/patrick/biz"
	"github.com/blackhorseya/patrick/internal/pkg/infra/log"
	"github.com/google/wire"
	"github.com/spf13/cobra"
)

var providerSet = wire.NewSet(
	log.ProviderSet,
	cmd.ProviderSet,
	biz.ProviderSet,
)

func CreateApp() (*cobra.Command, error) {
	panic(wire.Build(providerSet))
}
