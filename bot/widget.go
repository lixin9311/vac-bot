package bot

import tb "gopkg.in/tucnak/telebot.v2"

var (
	menu                 = &tb.ReplyMarkup{ResizeReplyKeyboard: true}
	btnHelp              = menu.Text("ℹ Help")
	btnSetting           = menu.Text("⚙ Setting")
	btnStubToggleAlert   = &tb.Btn{Unique: "btn_toggle_alert"}
	btnStubToggleReserve = &tb.Btn{Unique: "btn_toggle_reserve"}
	btnStubLogout        = &tb.Btn{Unique: "btn_logout"}
	btnStubPrune         = &tb.Btn{Unique: "btn_prune"}
)

func init() {
	menu.Reply(
		menu.Row(btnHelp, btnSetting),
	)
}
