package error

import (
	"fmt"
)

var (
	//通用错误
	ErrorIllegalParams       = NEWError(CommonIllegalParams)
	ErrorResourceExist       = NEWError(CommonResourceExist)
	ErrorNotFound            = NEWError(CommonResourceNoExist)
	ErrorInternalServerError = NEWError(CommonInternalServerError)
	ErrorTokenInvalidated    = NEWError(CommonTokenInvalidated)
)

//模块码[01]
const (
	//通用状态模块码 [000]
	CommonSuccess             = "00000001" //成功的码各模块通用,都是00000001
	CommonIllegalParams       = "01000002"
	CommonResourceNoExist     = "01000003"
	CommonResourceExist       = "01000004"
	CommonInternalServerError = "01000005"
	CommonTokenInvalidated    = "01000006"
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
	IMError struct {
		Code string
		Desc string
	}
)

func NEWError(code string) *IMError {
	return &IMError{Code: code, Desc: ErrorCodeToText(code)}
}

func (err *IMError) Error() string {
	errString := fmt.Sprintf("code = %s,desc = %s", err.Code, err.Desc)
	return errString
}
