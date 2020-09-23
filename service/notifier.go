package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	pb "notifier/api/gen/go"
	beeep "github.com/gen2brain/beeep"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"golang.org/x/net/context"

	"github.com/rakyll/statik/fs"
	_ "notifier/service/statik"
)

const (
	GRPC_CONTENT_TYPE = "application/grpc"
	GRPCWEB_CONTENT_TYPE = "application/grpc-web+proto"
)

type server struct {
}

func (s *server) ShowNotification(ctx context.Context, in *pb.ShowNotificationRequest) (*pb.ShowNotificationResponse, error) {
	err := beeep.Alert(in.Title, in.Message, "")
	if err != nil {
		panic(err)	
	}
	return &pb.ShowNotificationResponse{}, nil
}

func main() {
	cfg, err := configFromEnv()
	if err != nil {
		log.Printf("error reading configuration from env: %v", err)
		os.Exit(1)
	}

	initServer(cfg)
}

func initStaticServer(cfg *config) (http.Handler) {
	staticServer := http.FileServer(http.Dir("../ui/build"))
	if !cfg.Development {
		statikFS, err := fs.New()
		if err != nil {
			log.Fatal(err)
		}
		staticServer = http.FileServer(statikFS)
	}
	return staticServer
}

func initServer(cfg *config) {
	h2s := &http2.Server{}
	grpcServer := grpc.NewServer()
	pb.RegisterNotifierAPIServer(grpcServer, &server{})
	grpclog.SetLogger(log.New(os.Stdout, "notifier server: ", log.LstdFlags))

	wrappedServer := grpcweb.WrapServer(grpcServer)

	staticServer := initStaticServer(cfg)

	handler := http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		contentType := req.Header.Get("Content-Type")
		if contentType == GRPC_CONTENT_TYPE || contentType == GRPCWEB_CONTENT_TYPE {
			wrappedServer.ServeHTTP(res, req)
		} else {
			staticServer.ServeHTTP(res, req)
		}
	})

	server := &http.Server{
		Addr:    fmt.Sprintf("127.0.0.1:%s", cfg.Port),
		Handler: h2c.NewHandler(handler, h2s),
	}

	log.Printf("Starting HTTP/2 server on port: %s", cfg.Port)

	err := server.ListenAndServe()

	if err != nil {
		grpclog.Fatalf("failed starting HTTP/2 server: %v", err)
	}
}
