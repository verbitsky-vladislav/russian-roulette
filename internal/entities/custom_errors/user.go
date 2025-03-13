package custom_errors

const (
	ErrInsertUser                = "error inserting user"
	ErrUserNotFound              = "error user not found"
	ErrUpdateUser                = "error updating user"
	ErrFetchUsers                = "error fetching users"
	ErrGetUserByUUID             = "error getting user by uuid"
	ErrGetUserByChatID           = "error getting user by chat_id"
	ErrUserAlreadyExists         = "error user already exists"
	ErrUserNotFoundInContext     = "error user not found in context"
	ErrUserAlreadyJoinToGame     = "error user is already joined to game"
	ErrUserAlreadyHaveActiveGame = "error user is already have active game"
)
