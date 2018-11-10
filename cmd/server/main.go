package main

import (
	"flag"
	"log"
	"net"

	"github.com/GoingFast/test6/internal/server"
	"github.com/GoingFast/test6/pkg/env"
	pb "github.com/GoingFast/test6/protobuf"
	"google.golang.org/grpc"
)

var (
	sqlMode     = flag.Bool("sql", false, "use PostgresSQL repository")
	sqlAddr     = flag.String("sqladdr", "", "PostgresSQL address")
	sqlDatabase = flag.String("sqldb", "", "PostgresSQL database")
	sqlUsername = flag.String("sqluser", "", "PostgresSQL username")
	sqlPassword = flag.String("sqlpass", "", "PostgresSQL password")
)

func main() {
	flag.Parse()
	ln, err := net.Listen("tcp", env.FallbackEnv("LISTEN_ADDR", ":50051"))
	if err != nil {
		log.Fatal(err)
	}
	srv := grpc.NewServer()
	svc := server.NewService(*sqlMode, server.SQLConn{
		SQLAddr: *sqlAddr,
		SQLUser: *sqlUsername,
		SQLPass: *sqlPassword,
		SQLDB:   *sqlDatabase,
	})
	pb.RegisterCRUDServiceServer(srv, svc)
	srv.Serve(ln)
}
