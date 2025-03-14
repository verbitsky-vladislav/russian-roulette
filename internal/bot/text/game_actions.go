package text

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SuccessfulJoinGameMessage() string {
	return fmt.Sprintf("Вы успешно присоединились к игре! Ожидаем других игроков...")
}

func GameNotFoundMessage() string {
	return "Игра не найдена. Возможно, она уже завершилась или была удалена."
}

func GameIsAlreadyFullMessage() string {
	return "Игра уже набрала максимальное количество игроков. Попробуйте создать новую или присоединиться к другой."
}

func UserAlreadyJoinedMessage() string {
	return "Вы уже присоединились к этой игре. Ожидаем начала!"
}

func UserAlreadyHaveActiveGameMessage() string {
	return "Вы можете участвовать только в одной игре, сначала доиграйте её!"
}

func DefaultErrorMessage() string {
	return "Произошла неизвестная ошибка. Попробуйте снова или создайте новую игру."
}

func SuccessfulPassMessage(currentPlayer, nextPlayer string, bulletsLeft, roundsLeft int) string {
	return fmt.Sprintf(
		"🔄 *%s остаётся в игре и передаёт ход %s.*\n🎯 Осталось патронов: %d/6\n🕒 Осталось раундов: %d",
		"@"+tgbotapi.EscapeText(tgbotapi.ModeMarkdownV2, currentPlayer),
		"@"+tgbotapi.EscapeText(tgbotapi.ModeMarkdownV2, nextPlayer),
		bulletsLeft,
		roundsLeft,
	)
}

// todo fix экранирование
func SuccessfulPullMessage(currentPlayer string, bulletsLeft, roundsLeft int) string {
	return fmt.Sprintf(
		"💥 *%s остаётся в игре!*\nЧто будете делать дальше?\n👉 /pull (стрелять) или 🔄 /pass (передать ход)\n🎯 Осталось патронов: %d/6\n🔄 Раундов осталось: %d/6",
		"@"+tgbotapi.EscapeText(tgbotapi.ModeMarkdownV2, currentPlayer),
		bulletsLeft,
		roundsLeft,
	)
}

func UnsuccessfulPullMessage(currentPlayer, nextPlayer string, bulletsLeft, roundsLeft int) string {
	return fmt.Sprintf(
		"💀 *%s выбывает из игры!* 😵\nСледующий ход делает %s.\n🎯 Осталось патронов: %d/6\n🔄 Раундов осталось: %d/6\n\n🎮 *Доступные команды:*\n🔫 /pull – выстрелить\n🔄 /pass – передать револьвер",
		"@"+tgbotapi.EscapeText(tgbotapi.ModeMarkdownV2, currentPlayer),
		"@"+tgbotapi.EscapeText(tgbotapi.ModeMarkdownV2, nextPlayer),
		bulletsLeft,
		roundsLeft,
	)
}
