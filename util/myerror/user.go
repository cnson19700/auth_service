package myerror

import (
	"net/http"

	"github.com/cnson19700/pkg/apperror"
)

func ErrUserFormat(err error) apperror.AppError {
	return apperror.AppError{
		Raw:       err,
		ErrorCode: 100001,
		HTTPCode:  http.StatusNotAcceptable,
		Info:      "Create user is not valid",
		Message:   "Create user is not valid",
	}
}

func ErrGetUser(err error) apperror.AppError {
	return apperror.AppError{
		Raw:       err,
		ErrorCode: 100010,
		HTTPCode:  http.StatusNotAcceptable,
		Info:      "Fail to get User",
		Message:   "Fail to get USer",
	}
}

func ErrUpdatePassword(err error) apperror.AppError {
	return apperror.AppError{
		Raw:       err,
		ErrorCode: 100010,
		HTTPCode:  http.StatusNotAcceptable,
		Info:      "Fail to update password",
		Message:   "Fail to update password",
	}
}

func ErrInvalidPassword(err error) apperror.AppError {
	return apperror.AppError{
		Raw:       err,
		ErrorCode: 100020,
		HTTPCode:  http.StatusNotAcceptable,
		Info:      "Password is not correct",
		Message:   "Password is not correct",
	}
}

func ErrMatchPassword(err error) apperror.AppError {
	return apperror.AppError{
		Raw:       err,
		ErrorCode: 100030,
		HTTPCode:  http.StatusNotAcceptable,
		Info:      "2 Password is not match",
		Message:   "2 Password is not match",
	}
}

func ErrUpdateUser(err error) apperror.AppError {
	return apperror.AppError{
		Raw:       err,
		ErrorCode: 100070,
		HTTPCode:  http.StatusNotAcceptable,
		Info:      "Fail to update User",
		Message:   "Fail to update USer",
	}
}

func ErrAgeFormat(err error) apperror.AppError {
	return apperror.AppError{
		Raw:       err,
		ErrorCode: 30000021,
		HTTPCode:  http.StatusNotAcceptable,
		Info:      "age format is wrong",
		Message:   "age format is wrong.",
	}
}

func ErrFileOver5MB(err error) apperror.AppError {
	return apperror.AppError{
		Raw:       err,
		ErrorCode: 30000031,
		HTTPCode:  http.StatusNotAcceptable,
		Info:      "File cannot over 5Mb",
		Message:   "File cannot over 5Mb",
	}
}

func ErrOpenFile(err error) apperror.AppError {
	return apperror.AppError{
		Raw:       err,
		ErrorCode: 30000040,
		HTTPCode:  http.StatusNotAcceptable,
		Info:      "cannot open file",
		Message:   "cannot open file",
	}
}

func ErrReadBufferFail(err error) apperror.AppError {
	return apperror.AppError{
		Raw:       err,
		ErrorCode: 30000050,
		HTTPCode:  http.StatusNotAcceptable,
		Info:      "cannot read buffer of File",
		Message:   "cannot read buffer of File",
	}
}

func ErrNotImageFile(err error) apperror.AppError {
	return apperror.AppError{
		Raw:       err,
		ErrorCode: 30000060,
		HTTPCode:  http.StatusNotAcceptable,
		Info:      "not an image file.",
		Message:   "not an image file.",
	}
}
