//go:build wireinject
// +build wireinject

package bot

import (
	"github.com/google/wire"
	"github.com/yanzay/tbot/v2"
)

func Wire(client *tbot.Client) *handler {
	panic(wire.Build(ProviderSet))
}
