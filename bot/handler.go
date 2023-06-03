package bot

import (
	"github.com/yanzay/tbot/v2"
	"telegram-bot-reminder/domain"
)

type handler struct {
	client *tbot.Client
	svc    domain.BotService
}

func (h *handler) Ping(m *tbot.Message) {
	h.client.SendMessage(m.Chat.ID, h.svc.Ping())
}

func (h *handler) Notify(m *tbot.Message) {

	h.client.SendMessage(m.Chat.ID, m.Text)
}
