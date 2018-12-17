package api

import (
	"log"
	"net"

	api "github.com/augmentable-opensource/phraser/pkg/api"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"go.uber.org/zap"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	quit <-chan struct{}
}

// StartGRPC starts a gRPC service on the bind address
func StartGRPC(ctx context.Context, bind string, logger *zap.Logger) error {
	s := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_zap.StreamServerInterceptor(logger),
			grpc_recovery.StreamServerInterceptor(),
			grpc_prometheus.StreamServerInterceptor,
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_zap.UnaryServerInterceptor(logger),
			grpc_recovery.UnaryServerInterceptor(),
			grpc_prometheus.UnaryServerInterceptor,
		)),
	)
	defer logger.Sync()
	apiServer := &server{
		quit: ctx.Done(),
	}
	api.RegisterPhraserServer(s, apiServer)
	grpc_prometheus.Register(s)
	reflection.Register(s)

	listener, err := net.Listen("tcp", bind)
	if err != nil {
		return err
	}

	go func() {
		<-ctx.Done()
		log.Println("grpc service: stopping")
		s.GracefulStop()
		log.Println("grpc service: stopped")
		logger.Sync()
	}()

	return s.Serve(listener)
}

func (s *server) SetPhrase(ctx context.Context, params *api.SetPhraseRequest) (*api.Phrase, error) {
	return &api.Phrase{}, nil
}

func (s *server) GetPhrase(ctx context.Context, params *api.GetPhraseRequest) (*api.Phrase, error) {
	return &api.Phrase{}, nil
}

func (s *server) ListPhrases(params *api.ListPhrasesRequest, stream api.Phraser_ListPhrasesServer) error {
	return nil
}
