package myerror

import (
	"net/http"

	"github.com/cnson19700/pkg/apperror"
)

func ErrInsertBook(err error) apperror.AppError {
	return apperror.AppError{
		Raw:       err,
		ErrorCode: 200001,
		HTTPCode:  http.StatusNotAcceptable,
		Info:      "Create book is not valid",
		Message:   "Create book is not valid",
	}
}

func ErrGetBook(err error) apperror.AppError {
	return apperror.AppError{
		Raw:       err,
		ErrorCode: 200010,
		HTTPCode:  http.StatusNotAcceptable,
		Info:      "Fail to get Book",
		Message:   "Fail to get Book",
	}
}

func ErrDeleteBook(err error) apperror.AppError {
	return apperror.AppError{
		Raw:       err,
		ErrorCode: 200020,
		HTTPCode:  http.StatusNotAcceptable,
		Info:      "Fail to delete Book",
		Message:   "Fail to delete Book",
	}
}

func ErrTitleFormat(err error) apperror.AppError {
	return apperror.AppError{
		Raw:       err,
		ErrorCode: 200030,
		HTTPCode:  http.StatusNotAcceptable,
		Info:      "Title format is not valid",
		Message:   "Title format is not valid",
	}
}

func ErrLanguageFormat(err error) apperror.AppError {
	return apperror.AppError{
		Raw:       err,
		ErrorCode: 200040,
		HTTPCode:  http.StatusNotAcceptable,
		Info:      "Title format is not valid",
		Message:   "Title format is not valid",
	}
}

func ErrDescriptionFormat(err error) apperror.AppError {
	return apperror.AppError{
		Raw:       err,
		ErrorCode: 200050,
		HTTPCode:  http.StatusNotAcceptable,
		Info:      "Title format is not valid",
		Message:   "Title format is not valid",
	}
}

func ErrRatingAvgFormat(err error) apperror.AppError {
	return apperror.AppError{
		Raw:       err,
		ErrorCode: 200060,
		HTTPCode:  http.StatusNotAcceptable,
		Info:      "Rating Average format is not valid",
		Message:   "Rating Average format is not valid",
	}
}

func ErrReleaseDateFormat(err error) apperror.AppError {
	return apperror.AppError{
		Raw:       err,
		ErrorCode: 200060,
		HTTPCode:  http.StatusNotAcceptable,
		Info:      "Release Date format is not valid",
		Message:   "Release Date format is not valid",
	}
}

func ErrSearchTextFormat(err error) apperror.AppError {
	return apperror.AppError{
		Raw:       err,
		ErrorCode: 200070,
		HTTPCode:  http.StatusNotAcceptable,
		Info:      "Search book format is not valid",
		Message:   "Search book format is not valid",
	}
}


