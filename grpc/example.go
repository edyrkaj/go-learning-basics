package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"time"

	"golearn.com/grpc/hello"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const addr = "localhost:50051"

// greeterServer implements the generated GreeterServer interface.
// Embedding UnimplementedGreeterServer keeps the code forward-compatible
// when new RPCs are added to the .proto file.
type greeterServer struct {
	hello.UnimplementedGreeterServer
}

// SayHello is a unary RPC: one request in, one response out.
func (s *greeterServer) SayHello(_ context.Context, req *hello.HelloRequest) (*hello.HelloReply, error) {
	return &hello.HelloReply{
		Message: "Hello, " + req.GetName(),
	}, nil
}

// ListGreetings is a server-streaming RPC: one request in, many responses out.
func (s *greeterServer) ListGreetings(req *hello.HelloRequest, stream hello.Greeter_ListGreetingsServer) error {
	greetings := []string{"Hello", "Hi", "Hey"}
	for _, greeting := range greetings {
		reply := &hello.HelloReply{
			Message: greeting + ", " + req.GetName(),
		}
		if err := stream.Send(reply); err != nil {
			return err
		}
	}
	return nil
}

func startServer() *grpc.Server {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	hello.RegisterGreeterServer(s, &greeterServer{})

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	return s
}

func runClient() {
	// Dial over HTTP/2. insecure.NewCredentials() skips TLS — fine for localhost demos,
	// never for production. Real services use TLS (credentials.NewClientTLSFromFile).
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := hello.NewGreeterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// 1. Unary call
	unary, err := client.SayHello(ctx, &hello.HelloRequest{Name: "Eledi"})
	if err != nil {
		log.Fatalf("SayHello failed: %v", err)
	}
	fmt.Println("unary:", unary.GetMessage())

	// 2. Server-streaming call — Recv() until io.EOF
	stream, err := client.ListGreetings(ctx, &hello.HelloRequest{Name: "Eledi"})
	if err != nil {
		log.Fatalf("ListGreetings failed: %v", err)
	}
	for {
		reply, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("stream recv failed: %v", err)
		}
		fmt.Println("stream:", reply.GetMessage())
	}
}

func main() {
	// 1. Start the gRPC server (HTTP/2) in a background goroutine
	server := startServer()
	defer server.GracefulStop()

	// 2. Dial as a client and exercise unary + streaming RPCs
	runClient()
}
