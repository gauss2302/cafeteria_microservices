package main

import (
	common "github.com/gauss2302/cafemania_commons"
	pb "github.com/gauss2302/cafemania_commons/api"
	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
)

var (
	httpAddr = common.EnvString("HTTP_ADDR", ":8080")
	//orderService = common.EnvString("ORDER_SERVICE", "localhost:50051")
	orderServiceAddr = "localhost:2000"
)

func main() {
	conn, err := grpc.NewClient(orderServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}(conn)

	c := pb.NewOrderServiceClient(conn)

	log.Println("Dialing order service at ", orderServiceAddr)

	mux := http.NewServeMux()
	handler := NewHandler(c)
	handler.RegisterRoutes(mux)

	log.Printf("starting server on %s", httpAddr)

	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}

}
