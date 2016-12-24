package grpc

import ()

type Response interface {
	GetRid() string
	GetCode() string
	GetDesc() string
}
