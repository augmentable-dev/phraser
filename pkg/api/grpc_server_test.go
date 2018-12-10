package api_test

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	"github.com/augmentable-opensource/phraser/pkg/api"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

const (
	GRPCAddr = "127.0.0.1:50060"
)

var GRPCClient api.PhraserClient

func TestMain(m *testing.M) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalln(err)
	}

	go func() {
		if err := api.StartGRPC(ctx, GRPCAddr, logger); err != nil {
			logger.Fatal(err.Error())
		}
	}()

	time.Sleep(3 * 1000)

	conn, err := grpc.Dial(GRPCAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	GRPCClient = api.NewPhraserClient(conn)

	os.Exit(m.Run())
}

func TestGetPhrase(t *testing.T) {
	phrase, err := GRPCClient.GetPhrase(context.Background(), &api.GetPhraseRequest{Collection: "test-collection"})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(phrase)
}
