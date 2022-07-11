package main

import (
	"net"
	"os"
	"strings"

	//core "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	discovery "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"github.com/zirain/envoy-xds-server/pkg"
)

func main() {
	initLogging()

	listener, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		logrus.Fatalf("Failed to start listener on port 8080: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	// Register v3 server
	s := &pkg.DiscoveryServer{}
	discovery.RegisterAggregatedDiscoveryServiceServer(grpcServer, s)
	logrus.Println("Starting envoy XDS Server")
	if err := grpcServer.Serve(listener); err != nil {
		logrus.Fatalf("grpc serve err: %v", err)
	}
}

func initLogging() {
	levelEnv := os.Getenv("LOGGING_LEVEL")
	if levelEnv == "" {
		levelEnv = "info"
	}
	level, err := logrus.ParseLevel(strings.ToLower(levelEnv))
	if err != nil {
		logrus.Errorf("parse logging level, use info level")
		level = logrus.InfoLevel
	}

	logrus.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
	})
	logrus.SetLevel(level)
}
