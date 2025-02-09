package remote

import (
	"crypto/tls"
	"regexp"
	"strconv"

	"google.golang.org/grpc/credentials"

	grpctypes "github.com/cosmos/cosmos-sdk/types/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var (
	HTTPProtocols = regexp.MustCompile("https?://")
)

// GetHeightRequestHeader returns the grpc.CallOption to query the state at a given height
func GetHeightRequestHeader(height int64) grpc.CallOption {
	header := metadata.New(map[string]string{
		grpctypes.GRPCBlockHeightHeader: strconv.FormatInt(height, 10),
	})
	return grpc.Header(&header)
}

// MustCreateGrpcConnection creates a new gRPC connection using the provided configuration and panics on error
func MustCreateGrpcConnection(cfg *GRPCConfig) *grpc.ClientConn {
	grpConnection, err := CreateGrpcConnection(cfg)
	if err != nil {
		panic(err)
	}
	return grpConnection
}

// CreateGrpcConnection creates a new gRPC client connection from the given configuration
func CreateGrpcConnection(cfg *GRPCConfig) (*grpc.ClientConn, error) {
	var grpcOpts []grpc.DialOption
	if cfg.Insecure {
		grpcOpts = append(grpcOpts, grpc.WithInsecure())
	} else {
		grpcOpts = append(grpcOpts, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})))
	}

	address := HTTPProtocols.ReplaceAllString(cfg.Address, "")
	return grpc.Dial(address, grpcOpts...)
}
