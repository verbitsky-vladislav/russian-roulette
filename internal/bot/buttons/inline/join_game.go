package inline

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
)

func CreateJoinGameButton(gameUuid string) [][]tgbotapi.InlineKeyboardButton {
	return [][]tgbotapi.InlineKeyboardButton{
		{
			tgbotapi.NewInlineKeyboardButtonData("📝 Присоединиться", "join:"+gameUuid),
		},
		{
			tgbotapi.NewInlineKeyboardButtonData("🚫 Отменить", "cancel:"+gameUuid),
		},
	}
}

func IsJoinGameCallback(data string) bool {
	return strings.HasPrefix(data, "join:")
}

func IsCancelGameCallback(data string) bool {
	return strings.HasPrefix(data, "cancel:")
}

func ExtractGameUuid(data string) (string, bool) {
	parts := strings.SplitN(data, ":", 2)
	if len(parts) < 2 {
		return "", false
	}
	return parts[1], true
}
