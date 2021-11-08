package myerror

import (
	"net/http"

	"github.com/cnson19700/pkg/apperror"
)

func ErrContentFormat(err error) apperror.AppError {
	return apperror.AppError{
		Raw:       err,
		ErrorCode: 200070,
		HTTPCode:  http.StatusNotAcceptable,
		Info:      "Content format is not valid",
		Message:   "Content format is not valid",
	}
}

func ErrInsertComment(err error) apperror.AppError {
	return apperror.AppError{
		Raw:       err,
		ErrorCode: 200070,
		HTTPCode:  http.StatusNotAcceptable,
		Info:      "Content format is not valid",
		Message:   "Content format is not valid",
	}
}

func ErrFindComment(err error) apperror.AppError {
	return apperror.AppError{
		Raw:       err,
		ErrorCode: 200080,
		HTTPCode:  http.StatusNotAcceptable,
		Info:      "Comment not found",
		Message:   "Comment not found",
	}
}
