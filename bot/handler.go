package bot

import (
	"context"
	"errors"
	"html"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/lixin9311/vac-bot/ent"
	"github.com/lixin9311/vac-bot/ent/schema"
	"github.com/lixin9311/vac-bot/ent/vacuser"
	"github.com/lixin9311/vac-bot/tokyovacapi"
	"github.com/lixin9311/vac-bot/vacwatcher"
	tb "gopkg.in/tucnak/telebot.v2"
)

func nctx() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}

func (b *Bot) onStart(m *tb.Message) {
	ctx, cancel := nctx()
	defer cancel()
	// only handle private message
	if !m.Private() {
		return
	}
	_, err := b.db.VacUser.Create().SetSnsID("tg:" + strconv.Itoa(m.Sender.ID)).Save(ctx)
	if err != nil && !ent.IsConstraintError(err) {
		log.Printf("failed to create user: %v", err)
		b.sendError(m.Sender, err)
		return
	}
	b.bot.Send(m.Sender, html.EscapeString("Hello! You can start from '/help'"), menu)
}

func (b *Bot) onHelp(m *tb.Message) {
	b.bot.Send(m.Sender, "Help message")
}

func (b *Bot) onSetPartition(m *tb.Message) {
	ctx, cancel := nctx()
	defer cancel()
	parsed := strings.Split(m.Text, " ")
	if len(parsed) != 2 {
		b.bot.Send(m.Sender, "<code>"+html.EscapeString("/set_partition <municipal_code>")+"</code>")
		return
	}

	partition := parsed[1]
	if err := b.watcher.ValidateAndInitPartition(ctx, partition); err != nil {
		if errors.Is(err, vacwatcher.ErrPartitionNotExist) {
			b.bot.Send(m.Sender, "The partition does not exist, try a correct one.")
		} else {
			b.sendError(m.Sender, err)
		}
		return
	}

	err := b.db.VacUser.Update().Where(vacuser.SnsID("tg:" + strconv.Itoa(m.Sender.ID))).SetPartition(parsed[1]).Exec(ctx)
	if err != nil {
		b.sendError(m.Sender, err)
		return
	}
	b.bot.Send(m.Sender, "You have set partition to "+tokyovacapi.Partition(partition))
}

func (b *Bot) onLogin(m *tb.Message) {
	ctx, cancel := nctx()
	defer cancel()
	parsed := strings.Split(m.Text, " ")
	if len(parsed) != 3 {
		b.bot.Send(m.Sender, "<code>"+html.EscapeString("/login <username> <password>")+"</code>")
		return
	}
	user, err := b.db.VacUser.Query().Where(vacuser.SnsID("tg:" + strconv.Itoa(m.Sender.ID))).Only(ctx)
	if err != nil {
		b.sendError(m.Sender, err)
		return
	}
	if user.RangeKey != "" {
		b.bot.Send(m.Sender, "You have already logged in, please logout first.")
		return
	}
	if user.Partition == "" {
		b.bot.Send(m.Sender, "You have not set an partition yet, please use <code>"+html.EscapeString("/set_partition <municipal_code>")+"</code> first.")
		return
	}
	resp, err := b.watcher.ValidateLogin(ctx, user.Partition, parsed[1], parsed[2])
	if err != nil {
		b.bot.Send(m.Sender, "Failed to login: "+html.EscapeString(err.Error()))
		return
	}
	err = b.db.VacUser.UpdateOne(user).SetToken(resp.AccessToken).SetReservations(&tokyovacapi.ReservationList{
		Reservations: resp.Person.Reservations,
	}).Exec(ctx)
	if err != nil {
		b.sendError(m.Sender, err)
		return
	}
	b.bot.Send(m.Sender, "Login successed, go to setting to check it out.")
}

func (b *Bot) onLogout(m *tb.Message) {
	ctx, cancel := nctx()
	defer cancel()
	err := b.db.VacUser.Update().Where(vacuser.SnsID("tg:" + strconv.Itoa(m.Sender.ID))).
		ClearPassword().
		ClearRangeKey().
		ClearToken().
		SetReserveConfig(nil).
		SetReserveEnabled(false).
		Exec(ctx)
	if err != nil {
		b.sendError(m.Sender, err)
		return
	}
	b.bot.Send(m.Sender, "Your credentials are successfully removed.")
}

func (b *Bot) onPrune(m *tb.Message) {
	ctx, cancel := nctx()
	defer cancel()
	_, err := b.db.VacUser.Delete().Where(vacuser.SnsID("tg:" + strconv.Itoa(m.Sender.ID))).Exec(ctx)
	if err != nil {
		b.sendError(m.Sender, err)
		return
	}
	b.bot.Send(m.Sender, html.EscapeString("Your data has been pruned, I'm leaving the chat, you may need to '/start' conversation with me again."))
	b.bot.Leave(m.Chat)
}

func (b *Bot) onToggleAlert(m *tb.Message) {
	ctx, cancel := nctx()
	defer cancel()
	user, err := b.db.VacUser.Query().Where(vacuser.SnsID("tg:" + strconv.Itoa(m.Sender.ID))).Only(ctx)
	if err != nil {
		b.sendError(m.Sender, err)
		return
	}
	err = b.db.VacUser.UpdateOne(user).SetWatcherEnabled(!user.WatcherEnabled).Exec(ctx)
	if err != nil {
		b.sendError(m.Sender, err)
		return
	}
	if user.WatcherEnabled {
		b.bot.Send(m.Sender, "You have disabled Watcher")
	} else {
		b.bot.Send(m.Sender, "You have enabled Watcher")
	}
}

func (b *Bot) onReservationConfig(m *tb.Message) {
	ctx, cancel := nctx()
	defer cancel()
	parsed := strings.Split(m.Text, " ")
	if len(parsed) != 3 || len(parsed) != 4 {
		b.bot.Send(m.Sender, "<code>"+html.EscapeString("/set_reservation_config <start_offset> <end_offset> <start_date*>")+"</code>")
		return
	}
	startOffset, err := ParseDuration(parsed[1])
	if err != nil {
		b.bot.Send(m.Sender, "❌ start offset is invalid")
		return
	}
	endOffset, err := ParseDuration(parsed[2])
	if err != nil {
		b.bot.Send(m.Sender, "❌ end offset is invalid")
		return
	}
	startDate := time.Now()
	if len(parsed) == 4 {
		startDate, err = time.Parse("2006-01-02", parsed[3])
		if err != nil {
			b.bot.Send(m.Sender, "❌ start date is invalid")
			return
		}
	}
	config := &schema.ReserveConfig{
		StartOffset: startOffset,
		EndOffset:   endOffset,
		StartDate:   startDate,
	}
	_, err = b.db.VacUser.Update().Where(vacuser.SnsID("tg:" + strconv.Itoa(m.Sender.ID))).SetReserveConfig(config).Save(ctx)
	if err != nil {
		b.sendError(m.Sender, err)
		return
	}
	b.bot.Send(m.Sender, "successfully save your reservation config")
}

func (b *Bot) onToggleReserve(m *tb.Callback) {
	ctx, cancel := nctx()
	defer cancel()
	user, err := b.db.VacUser.Query().Where(vacuser.SnsID("tg:" + strconv.Itoa(m.Sender.ID))).Only(ctx)
	if err != nil {
		b.sendError(m.Sender, err)
		return
	}
	if user.Reservations != nil && len(user.Reservations.Reservations) == 2 {
		b.bot.Send(m.Sender, "You have already reserved all 2 times.")
	}
	// TODO: update user info via API, disable toggle accordingly
	b.bot.Send(m.Sender, "toggle reserve: ")
}

func (b *Bot) sendError(u *tb.User, err error) {
	_, err = b.bot.Send(u, `❌ Opps! Some error happend:\n<code>`+err.Error()+`</code>`)
	if err != nil {
		log.Printf("failed to send error: %v", err)
	}
}

func (b *Bot) onSetting(m *tb.Message) {
	login := false
	alertEnabled := false
	reserveEnabled := false
	msg := ""
	userSettingInline := &tb.ReplyMarkup{}
	// TODO: update user info
	if login {
		var btnToggleAlert, btnToggleReserve, btnLogOut, btnPrune tb.Btn
		msg += "You are logged in, vaccine ticket number: 000000000"
		if alertEnabled {
			btnToggleAlert = userSettingInline.Data("🔕 Disable Alert", "btn_toggle_alert")
		} else {
			btnToggleAlert = userSettingInline.Data("🔔 Enable Alert", "btn_toggle_alert")
		}
		if reserveEnabled {
			btnToggleReserve = userSettingInline.Data("Enable Auto-Reserve", "btn_toggle_reserve")
		} else {
			btnToggleReserve = userSettingInline.Data("Disable Auto-Reserve", "btn_toggle_reserve")
		}
		btnLogOut = userSettingInline.Data("logout", "btn_logout")
		btnPrune = userSettingInline.Data("Prune Data", "btn_prune")
		userSettingInline.Inline(
			userSettingInline.Row(btnToggleAlert, btnToggleReserve),
			userSettingInline.Row(btnLogOut, btnPrune),
		)
	} else {
		msg += "You are not logged in, type '/login <username> <password>' to login"
		var btnToggleAlert tb.Btn
		if alertEnabled {
			btnToggleAlert = userSettingInline.Data("🔕 Disable Alert", "btn_toggle_alert")
		} else {
			btnToggleAlert = userSettingInline.Data("🔔 Enable Alert", "btn_toggle_alert")
		}
		userSettingInline.Reply()
		userSettingInline.Inline(
			userSettingInline.Row(btnToggleAlert),
		)
	}
	b.bot.Send(m.Sender, msg, userSettingInline)
}
