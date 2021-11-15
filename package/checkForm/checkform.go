package checkform

import (
	"strconv"
	"strings"

	"github.com/cnson19700/auth_service/package/auth"

	errmsg "github.com/cnson19700/auth_service/package/logMessages"

	"github.com/cnson19700/auth_service/package/namestand"

	"github.com/labstack/echo/v4"
)

var appErr = errmsg.InitErrMsg()

func IsFormFullFill(c echo.Context, array []string) bool {

	isFullFill := true
	for i := 0; i < len(array); i++ {
		if c.FormValue(array[i]) == "" {
			isFullFill = false
			break
		}
	}
	return isFullFill
}

func CheckFormatValue(formAtributeName string, value string) (bool, string) {
	value = namestand.RemoveDoubleSpace(value)
	if value == "" {
		return false, appErr.AuthMsg.EmptyValue + formAtributeName
	}
	switch formAtributeName {
	case "email":
		if !auth.ValidEmail(value) {
			return false, appErr.AuthMsg.WrongMailFormat
		}
		return true, value
	case "age":
		age, err := strconv.Atoi(value)
		minAge := 0
		maxAge := 112

		if err != nil || age < minAge || age > maxAge {
			return false, appErr.AuthMsg.AgeNotTrue
		}

		return true, value
	case "full_name":
		str := namestand.Check(value)
		if str == "" {
			return false, appErr.AuthMsg.WrongFullName
		}
		return true, str
	case "page":
		_, err := strconv.Atoi(value)
		if err != nil {
			return false, appErr.QueryMsg.PageNotNumber
		}
		return true, value
	case "search", "overview", "content":
		str := namestand.FormatText(value, true, true)
		if str == "" {
			return false, formAtributeName + appErr.QueryMsg.WrongFomat
		}
		return true, str
	case "min_rating":
		num, err := strconv.Atoi(value)
		if err != nil || num < 0 || num > 10 {
			return false, appErr.QueryMsg.MinRatingWrong
		}
		return true, value
	case "cate_id", "author_id", "book_id", "parent_id":
		_, err := strconv.Atoi(value)
		if err != nil {
			return false, formAtributeName + appErr.QueryMsg.MustBeNumber
		}
		return true, value
	case "original_language", "original_title", "spoken_language":
		str := namestand.FormatText(value, false, false)
		if str == "" {
			return false, formAtributeName + appErr.QueryMsg.WrongFomat
		}
		return true, str
	case "release_date": //dung rfc339 check o delivery
		//str := namestand.IsDate(value)
		// if str == "" {
		// 	return false, formAtributeName + appErr.QueryMsg.WrongFomat
		// }
		return true, value
	case "rating_everage":
		num, err := strconv.ParseFloat(value, 64)
		if err != nil || num > 10 || num < 0 {
			return false, formAtributeName + appErr.QueryMsg.MustBeNumber
		}
		return true, value
	case "status":
		num, err := strconv.Atoi(value)
		if err != nil || num < 0 || num > 2 {
			return false, appErr.QueryMsg.StatusWrong
		}
		return true, value
	case "password":
		str := namestand.FormatText(value, true, true)
		if str == "" || strings.Contains(str, " ") {
			return false, formAtributeName + appErr.QueryMsg.WrongFomat
		}
		// } else if len(str) <= 8 {
		// 	return false, appErr.AuthMsg.Min8Character
		// }
		return true, str
	default:
		return true, value
	}
}
