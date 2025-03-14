package text

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"russian-roulette/internal/entities/types"
	"strings"
)

func WrongRouletteCommandMessage() string {
	return "Использование: /roulette [кол-во игроков (2-6)] [ставка].\n Пример: /roulette 4 100"
}

func WrongRouletteParamsMessage() string {
	return "Некорректные параметры. Кол-во игроков: 2-6. Ставка должна быть > 0."
}

func GameAlreadyExistsMessage() string {
	return "Вы уже создали игру. Чтобы создать новую, отмените или сыграйте предыдущую игру."
}

func NewRouletteGameMessage(players, bullets int, bet *types.Decimal) string {
	return fmt.Sprintf(
		"Новая игра русской рулетки!\nИгроков: %d\nСтавка: %s\nПатронов в магазине: %d\n\nЖмите кнопку, чтобы присоединиться!\n\n"+
			"📜 /players – список игроков\n"+
			"🔫 /pull – выстрел (обязательное действие)\n"+
			"⏭ /pass – передать револьвер",
		players, bet.Round(2).String(), bullets,
	)
}

func StartGameMessage(playersName []string, firstPlayer string) string {
	escapedPlayers := make([]string, len(playersName))
	for i, name := range playersName {
		escapedPlayers[i] = "@" + tgbotapi.EscapeText(tgbotapi.ModeMarkdownV2, name)
	}

	escapedFirstPlayer := "@" + tgbotapi.EscapeText(tgbotapi.ModeMarkdownV2, firstPlayer)

	return fmt.Sprintf(`🎰 *Игра началась!* 🎰
Добро пожаловать в русскую рулетку! ☠️🔫

🔫 *Игроки:*
%s

📜 *Правила игры:*
- В барабане револьвера *6 слотов*, случайно распределены *патроны* и *пустые гильзы*.
- Игрок *обязан стрелять* хотя бы *один раз*, но может рискнуть и повторить.
- *Передать ход* можно *только после первого выстрела*.
- Если выстрел был холостым – ход передаётся дальше.
- Если попался патрон – игрок выбывает, игра продолжается.

Первым совершает действие: %s

🎮 *Доступные команды:*
🔫 /pull – выстрелить (можно несколько раз).
🔄 /pass – передать револьвер (только после выстрела).
👥 /players – список участников.

🔥 *Удачи, бойцы!* Пусть фортуна будет на вашей стороне! 💀`, strings.Join(escapedPlayers, " "), escapedFirstPlayer)
}

func SuccessfulCancelGameMessage() string {
	return "Игра успешно отменена. Все ставки возвращены игрокам."
}

func FinishGameMessage(winner string, players []string, betAmount, bulletsUsed int) string {
	escapedPlayers := make([]string, len(players))
	for i, name := range players {
		escapedPlayers[i] = "@" + tgbotapi.EscapeText(tgbotapi.ModeMarkdownV2, name)
	}

	return fmt.Sprintf(
		"🏆 *Игра окончена!* 🎉\n\n"+
			"🎯 Использовано патронов: %d/6\n"+
			"💰 Ставка: *%d USDT*\n"+
			"👥 Участники: %s\n\n"+
			"🥇 *Победитель:* %s 🎊\n\n"+
			"Спасибо за игру! 🚀 Хотите сыграть еще раз? /roulette",
		bulletsUsed,
		betAmount,
		strings.Join(escapedPlayers, ", "),
		"@"+tgbotapi.EscapeText(tgbotapi.ModeMarkdownV2, winner),
	)
}
