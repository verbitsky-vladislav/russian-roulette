package text

import "fmt"

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

func SuccessfulPassMessage() {

}

func SuccessfulPullMessage() {

}

func UnsuccessfulPullMessage() {

}
