package appresp

import (
	"donkey-ucenter/apperror"
	"errors"
	"github.com/go-sql-driver/mysql"
)

type RespData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Ok() *RespData {
	return &RespData{
		Code:    0,
		Message: "success",
	}
}

func Data(data interface{}) *RespData {
	return &RespData{
		Code:    0,
		Message: "success",
		Data:    data,
	}
}

func CodeErr(code int, err error) *RespData {
	return &RespData{
		Code:    code,
		Message: err.Error(),
	}
}

func extractErr(err error) (int, string) {
	if err == nil {
		return 0, "success"
	}

	code := 9999
	var appErr apperror.AppError
	ok := errors.As(err, &appErr)
	if ok {
		return appErr.Code, appErr.Message
	}

	mysqlErr := new(mysql.MySQLError)
	ok = errors.As(err, &mysqlErr)
	if ok {
		code = int(mysqlErr.Number)
		return code, mysqlErr.Message
	}
	return code, err.Error()
}

func Err(err error) *RespData {
	if err == nil {
		return Ok()
	}

	code, msg := extractErr(err)

	return &RespData{
		Code:    code,
		Message: msg,
	}
}

func ErrMsg(errMsg string) *RespData {
	return &RespData{
		Code:    9999,
		Message: errMsg,
	}
}

func ErrData(data interface{}, err error) *RespData {
	if err != nil {
		return Data(data)
	}

	code, msg := extractErr(err)

	return &RespData{
		Code:    code,
		Message: msg,
		Data:    data,
	}
}

func Reps(data interface{}, err error) *RespData {
	if err != nil {
		return Err(err)
	}
	return Data(data)
}

func Raw(code int, data interface{}, err error) *RespData {
	return &RespData{
		Code:    code,
		Message: err.Error(),
		Data:    data,
	}
}
