package custom_errors

import "fmt"

// CustomError - структура для ошибок с кодом и сообщением.
type CustomError struct {
	Code    string // Уникальный код ошибки
	Message string // Сообщение об ошибке
}

// Error - метод для преобразования ошибки в строку.
func (e CustomError) Error() string {
	return e.Message
}

// Common errors
var (
	ErrMessageSending = CustomError{
		Code:    "ErrMessageSending",
		Message: fmt.Sprintf("Произошла ошибка при отправке сообщения. Обратитесь в поддержку или попробуйте позже /help"),
	}
	ErrUserNotFound = CustomError{
		Code:    "ErrUserNotFound",
		Message: "Кажется, вы еще не зарегистрировались. Пожалуйста, зарегистрируйтесь, чтобы начать /start",
	}
)

// Commands errors
var (
	ErrNoCommandFound = CustomError{
		Code:    "ErrNoCommandFound",
		Message: fmt.Sprintf("Не удалось найти такую команду. Список команд тут /help"),
	}
	ErrChatOnlyCommand = CustomError{
		Code:    "ErrChatOnlyCommand",
		Message: fmt.Sprintf("Эту команду можно вызвать только внутри диалога с ботом @eth_russian_roulette_bot"),
	}
	ErrGroupOnlyCommand = CustomError{
		Code:    "ErrChatOnlyCommand",
		Message: fmt.Sprintf("Эту команду можно вызвать только внутри группового чата @eth_russian_roulette"),
	}
)
