package main

import (
	"log/slog"

	"google.golang.org/grpc"

	"github.com/tesserical/cicada/grpc/cicadaapiv1pb"
)

func main() {
	srv := grpc.NewServer()
	cicadaapiv1pb.RegisterChatFetcherServer(srv, nil)
	slog.Info("some message")
}
