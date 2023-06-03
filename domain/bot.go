package domain

import "github.com/yanzay/tbot/v2"

type (
	BotService interface {
		Ping() string
		Notify(cronjob, msg string)
	}

	BotHandler interface {
		Ping(m *tbot.Message)
		Notify(m *tbot.Message)
	}
)
