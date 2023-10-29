package grpcgateway

import (
	"context"
	"github.com/LCY2013/grpc-cloud/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"time"
)

func Dial(ctx context.Context, address string) (*grpc.ClientConn, error) {
	dialTime := 30 * time.Second
	ctx, cancel := context.WithTimeout(ctx, dialTime)
	defer cancel()
	var opts []grpc.DialOption
	timeout := time.Duration(30 * float64(time.Second))
	opts = append(opts, grpc.WithKeepaliveParams(keepalive.ClientParameters{
		Time:    timeout,
		Timeout: timeout,
	}))
	opts = append(opts, grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(1024*1024*1024)))

	UA := "grpc-gateway/fufeng"
	opts = append(opts, grpc.WithUserAgent(UA))

	network := "tcp"
	cc, err := BlockingDial(ctx, network, address, nil, opts...)
	if err != nil {
		logger.Log.Error(err, "Failed to Dial target host %q", address)
	}
	return cc, err
}
