package bot

import (
	"time"

	"github.com/lixin9311/vac-bot/ent"
	"github.com/lixin9311/vac-bot/vacwatcher"
	tb "gopkg.in/tucnak/telebot.v2"
)

var (
	commands = []tb.Command{
		{
			Text:        "help",
			Description: "print help",
		},
		{
			Text:        "set_partition",
			Description: "/set_partition <municipal_code>",
		},
		{
			Text:        "set_reservation_config",
			Description: "/set_reservation_config <start_offset> <end_offset> <start_date*>\nex: /set_reservation_config 2d 60d 2021-07-09",
		},
		{
			Text:        "login",
			Description: "/login <username> <password>",
		},
		{
			Text:        "setting",
			Description: "show setting and infos",
		},
		{
			Text:        "prune",
			Description: "prune all the data stored in the bot",
		},
	}
)

type Bot struct {
	bot     *tb.Bot
	db      *ent.Client
	watcher *vacwatcher.Watcher
}

func NewBot(token string) (*Bot, error) {
	b, err := tb.NewBot(tb.Settings{
		Token: token,
		Poller: &tb.MiddlewarePoller{
			Poller: &tb.LongPoller{Timeout: 10 * time.Second},
			Filter: func(up *tb.Update) bool {
				if up.Message != nil && !up.Message.Private() {
					return false
				} else if up.EditedMessage != nil && !up.EditedMessage.Private() {
					return false
				}
				return true
			},
		},
		ParseMode: tb.ModeHTML,
	})
	if err != nil {
		return nil, err
	}
	b.SetCommands(commands)
	bot := &Bot{
		bot: b,
	}
	bot.init()
	return bot, nil
}

func (b *Bot) init() {
	b.bot.Handle("/start", b.onStart)
	b.bot.Handle("/set_partition", b.onSetPartition)
	b.bot.Handle("/login", b.onLogin)
	b.bot.Handle("/help", b.onHelp)
	b.bot.Handle("/setting", b.onSetting)
	b.bot.Handle("/prune", b.onPrune)
	b.bot.Handle("/set_reservation_config", b.onReservationConfig)

	b.bot.Handle(&btnHelp, b.onHelp)
	b.bot.Handle(&btnSetting, b.onSetting)
	b.bot.Handle(&btnStubToggleAlert, b.onToggleAlert)
	b.bot.Handle(&btnStubToggleReserve, b.onToggleReserve)
	b.bot.Handle(&btnStubLogout, b.onLogout)
	b.bot.Handle(&btnStubPrune, b.onPrune)
}
