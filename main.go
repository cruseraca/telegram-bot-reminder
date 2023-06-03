package main

import (
	"github.com/joho/godotenv"
	cron2 "github.com/robfig/cron/v3"
	"github.com/yanzay/tbot/v2"
	"log"
	"os"
	"sync"
	bot2 "telegram-bot-reminder/bot"
)

var (
	bot   *tbot.Server
	token string

	cron     *cron2.Cron
	cronOnce sync.Once
)

func init() {
	e := godotenv.Load()
	if e != nil {
		log.Println(e)
	}
	token = os.Getenv("TELEGRAM_TOKEN")
}
func main() {
	// initiate telegram bot server
	bot = tbot.New(token)

	// initiate cron
	initiateCron()

	// define command handler
	botHandler := bot2.Wire(bot.Client())
	bot.HandleMessage("/ping", botHandler.Ping)
	bot.HandleMessage("/notify .+", botHandler.Notify)

	log.Fatal(bot.Start())
}

func initiateCron() {
	cronOnce.Do(func() {
		cron = cron2.New()
	})
}
