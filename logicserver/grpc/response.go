package grpc

import (
// imError "im/logicserver/error"
// "strings"
)

type Response interface {
	GetRid() string
	GetCode() string
	GetDesc() string
	// Code string      `json:"code"`
	// Desc string      `json:"desc"`
	// Data interface{} `json:"data"`
}

// func (r *Response) IsSuccesss() bool {
// 	return strings.EqualFold(r.Code, imError.CommonSuccess)
// }

// func (r *Response) IsFail() bool {
// 	return !r.IsSuccesss()
// }

// func (r *Response) ResponseToError() *imError.IMError {
// 	if r.IsSuccesss() {
// 		return nil
// 	}

// 	err := &imError.IMError{
// 		Code: r.Code,
// 		Desc: r.Desc,
// 	}

// 	return err
// }

// func ResponseToError(response *Response) *imError.IMError {

// 	if response.IsSuccesss() {
// 		return nil
// 	}

// 	err := &imError.IMError{
// 		Code: response.Code,
// 		Desc: response.Desc,
// 	}

// 	return err
// }
