package custom_errors

const (
	ErrInsertUser            = "error inserting user"
	ErrUserNotFound          = "user not found"
	ErrUpdateUser            = "error updating user"
	ErrFetchUsers            = "error fetching users"
	ErrGetUserByUUID         = "error getting user by uuid"
	ErrGetUserByChatID       = "error getting user by chat_id"
	ErrUserAlreadyExists     = "user already exists"
	ErrUserNotFoundInContext = "user not found in context"
)
