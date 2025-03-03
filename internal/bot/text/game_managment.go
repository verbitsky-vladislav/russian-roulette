package text

import (
	"fmt"
	"russian-roulette/internal/entities/types"
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
