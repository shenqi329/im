package grpc

type Request interface {
	GetRid() uint64
}
