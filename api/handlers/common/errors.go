package common

const (
	ErrFailedToParseJson = "failed to parse json body"
	ErrFailedToGetUserId = "failed to get user id from request context"
	ErrDatabaseError     = "something bad happened to database"
	ErrNotFound          = "resource not found"
	ErrResourceNotOwned  = "current user does not own such a resource"
)
