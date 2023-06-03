package bot

import (
	"github.com/google/wire"
	"github.com/yanzay/tbot/v2"
	"sync"
	"telegram-bot-reminder/domain"
)

var (
	hdl     *handler
	hdlOnce sync.Once

	svc     *service
	svcOnce sync.Once

	ProviderSet wire.ProviderSet = wire.NewSet(
		ProvideHandler,
		ProvideService,

		//bind each one of the interfaces
		wire.Bind(new(domain.BotHandler), new(*handler)),
		wire.Bind(new(domain.BotService), new(*service)),
	)
)

func ProvideHandler(client *tbot.Client, svc domain.BotService) *handler {
	hdlOnce.Do(func() {
		hdl = &handler{
			client: client,
			svc:    svc,
		}
	})

	return hdl
}

func ProvideService() *service {
	svcOnce.Do(func() {
		svc = &service{}
	})

	return svc
}
