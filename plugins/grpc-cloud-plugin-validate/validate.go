package grpc_cloud_plugin_validate

import (
	"context"
	"github.com/LCY2013/grpc-cloud/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

/*// Authorization unary interceptor function to handle authorize per RPC call
func serverInterceptor(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()
	// Skip authorize when GetJWT is requested
	if info.FullMethod != "/proto.EventStoreService/GetJWT" {
		if err := authorize(ctx); err != nil {
			return nil, err
		}
	}

	// Calls the handler
	h, err := handler(ctx, req)

	// Logging with grpclog (grpclog.LoggerV2)
	grpcLog.Infof("Request - Method:%s\tDuration:%s\tError:%v\n",
		info.FullMethod,
		time.Since(start),
		err)

	return h, err
}

// authorize function authorizes the token received from Metadata
func authorize(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(codes.InvalidArgument, "Retrieving metadata is failed")
	}

	authHeader, ok := md["authorization"]
	if !ok {
		return status.Errorf(codes.Unauthenticated, "Authorization token is not supplied")
	}

	token := authHeader[0]
	// validateToken function validates the token
	err := validateToken(token)

	if err != nil {
		return status.Errorf(codes.Unauthenticated, err.Error())
	}
	return nil
}*/

func WithValidateServerInterceptor(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()
	defer func() {
		logger.Log.Infof("Request - Method:%s\tDuration:%s\t\n",
			info.FullMethod,
			time.Since(start))
	}()

	// validate
	if validate, ok := req.(interface{ ValidateAll() error }); ok {
		if err := validate.ValidateAll(); err != nil {
			logger.Log.Infof("Request - Method:%s\tDuration:%s\tError:%v\n",
				info.FullMethod,
				time.Since(start),
				err)
			return nil, status.Errorf(codes.InvalidArgument, err.Error())
		}
	}

	// Calls the handler
	h, err := handler(ctx, req)
	if err != nil {
		logger.Log.Infof("Request - Method:%s\tDuration:%s\tError:%v\n",
			info.FullMethod,
			time.Since(start),
			err)
		return nil, status.Errorf(codes.Unavailable, err.Error())
	}
	return h, nil
}
