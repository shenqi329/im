package error

import (
	"fmt"
	"strings"
)

var (
	//通用错误
	ErrorIllegalParams       = NewAccessError(CommonIllegalParams)
	ErrorResourceExist       = NewAccessError(CommonResourceExist)
	ErrorNotFound            = NewAccessError(CommonResourceNoExist)
	ErrorInternalServerError = NewAccessError(CommonInternalServerError)
	ErrorTokenInvalidated    = NewAccessError(CommonTokenInvalidated)
)

//模块码[01]
const (
	//通用状态模块码 [000]
	CommonSuccess             = "00000001" //成功的码各模块通用,都是00000001
	CommonIllegalParams       = "02000002"
	CommonResourceNoExist     = "02000003"
	CommonResourceExist       = "02000004"
	CommonInternalServerError = "02000005"
	CommonTokenInvalidated    = "02000006"
)

var codeText = map[string]string{
	//通用状态
	CommonSuccess:             "success",
	CommonIllegalParams:       "illegal parameter",
	CommonResourceNoExist:     "resource doesn't exist",
	CommonResourceExist:       "resource already exists",
	CommonInternalServerError: "internal server wrong",
	CommonTokenInvalidated:    "token invalidated",
}

func ErrorCodeToText(code string) string {
	return codeText[code]
}

type (
	AccessError struct {
		Code string
		Desc string
	}
)

func CodeIsSuccess(code string) bool {
	return strings.EqualFold(CommonSuccess, code)
}

func NewAccessError(code string) *AccessError {
	return &AccessError{Code: code, Desc: ErrorCodeToText(code)}
}

func (err *AccessError) Error() string {
	errString := fmt.Sprintf("code = %s,desc = %s", err.Code, err.Desc)
	return errString
}
