package grpcgateway

import (
	"context"
	"github.com/LCY2013/grpc-cloud/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/keepalive"
	"sync"
	"time"
)

// conn pool for grpc
var (
	pool  = &connPool{}
	gLock sync.Mutex
)

// connPool is a pool of gRPC connections.
type connPool struct {
	mut   sync.RWMutex
	conns map[string]*grpc.ClientConn
}

// conn returns a cached connection to the given address.
func (cp *connPool) conn(address string) *grpc.ClientConn {
	cp.mut.RLock()
	defer cp.mut.RUnlock()
	return cp.conns[address]
}

// putConn caches the given connection.
func (cp *connPool) putConn(address string, cc *grpc.ClientConn) {
	cp.mut.Lock()
	defer cp.mut.Unlock()
	cp.conns[address] = cc
}

// Dial dials to the given address with the given options.
func Dial(ctx context.Context, address string) (*grpc.ClientConn, error) {
	gLock.Lock()
	defer gLock.Unlock()
	if cc := pool.conn(address); cc != nil && cc.GetState() == connectivity.Ready {
		return cc, nil
	} else {
		err := cc.Close()
		if err != nil {
			return nil, err
		}
	}
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

	pool.putConn(address, cc)
	return cc, err
}
