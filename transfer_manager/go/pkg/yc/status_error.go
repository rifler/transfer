package yc

import (
	grpcutil "github.com/doublecloud/tross/transfer_manager/go/pkg/util/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GetStatus(err error) *status.Status {
	if ok, statusErr := grpcutil.UnwrapStatusError(err); ok {
		return statusErr.GRPCStatus()
	}
	return status.New(codes.Unknown, err.Error())
}