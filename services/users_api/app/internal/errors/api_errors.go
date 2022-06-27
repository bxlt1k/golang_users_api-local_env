package errors

import "net/http"

type ApiError struct {
	Status  int
	Message string
	Error   error
}

func NewApiError(status int, message string, error error) *ApiError {
	return &ApiError{
		Status:  status,
		Message: message,
		Error:   error,
	}
}

func UnauthorizedError() *ApiError {
	return NewApiError(http.StatusUnauthorized, "Пользователь не авторизован", nil)
}

func ForbiddenError() *ApiError {
	return NewApiError(http.StatusForbidden, "Доступ закрыт", nil)
}

func BadRequestError(message string, error error) *ApiError {
	return NewApiError(http.StatusBadRequest, message, error)
}

func InternalServerError(error error) *ApiError {
	return NewApiError(http.StatusInternalServerError, "Упс... Что-то пошло не так...", error)
}
