package main

import (
	"context"
	"crypto/tls"
	"database/sql"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"sync"

	pb "github.com/GoingFast/test6/protobuf"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	sqlMode     = flag.Bool("sql", false, "use PostgresSQL repository")
	sqlAddr     = flag.String("sqladdr", "", "PostgresSQL address")
	sqlDatabase = flag.String("sqldb", "", "PostgresSQL database")
	sqlUsername = flag.String("sqluser", "", "PostgresSQL username")
	sqlPassword = flag.String("sqlpass", "", "PostgresSQL password")
)

type repository interface {
	Read(ctx context.Context) ([]*pb.Request, error)
}

type inmemRepo struct {
	mu sync.RWMutex
	rs map[string]*pb.Request
}

func newInmemRepo() *inmemRepo {
	return &inmemRepo{
		rs: map[string]*pb.Request{
			"1": &pb.Request{Message: os.Getenv("HOSTNAME")},
			// "2": &pb.Request{Message: "bar"},
			// "3": &pb.Request{Message: "asd"},
			// "4": &pb.Request{Message: "fgh"},
		},
	}
}

func (i *inmemRepo) Read(ctx context.Context) ([]*pb.Request, error) {
	i.mu.RLock()
	defer i.mu.RUnlock()

	var rs []*pb.Request
	for _, r := range i.rs {
		rs = append(rs, r)
	}
	return rs, nil
}

type sqlRepo struct {
	*sql.DB
}

func newSqlRepo(db *sql.DB) sqlRepo {
	return sqlRepo{db}
}

func (s sqlRepo) Read(ctx context.Context) ([]*pb.Request, error) {
	var rs []*pb.Request
	rows, err := s.Query("SELECT message FROM requests")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var r pb.Request
		err := rows.Scan(&r.Message)
		if err != nil {
			return nil, err
		}

		rs = append(rs, &r)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return rs, nil
}

type service struct {
	repo repository
}

type SQLConn struct {
	SQLAddr string
	SQLUser string
	SQLPass string
	SQLDB   string
}

func db(s SQLConn) *sql.DB {
	conn, err := sql.Open("postgres", fmt.Sprintf("%s:%s@%s/%s?sslmode=disable",
		s.SQLUser,
		s.SQLPass,
		s.SQLAddr,
		s.SQLAddr,
	))
	if err != nil {
		log.Fatal(err)
	}

	var n int
	err = conn.QueryRow("SELECT 1").Scan(&n)
	if err != nil {
		log.Fatal(err)
	}

	return conn
}

func newService(sql bool, s SQLConn) service {
	if sql {
		return service{newSqlRepo(db(s))}
	}
	return service{newInmemRepo()}
}

func (s service) Read(ctx context.Context, _ *empty.Empty) (*pb.ReadResponse, error) {
	r, err := s.repo.Read(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.ReadResponse{Data: r}, nil
}

func fallbackEnv(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}

func main() {
	flag.Parse()
	ln, err := net.Listen("tcp", fallbackEnv("LISTEN_ADDR", ":50051"))
	if err != nil {
		log.Fatal(err)
	}
	var srv *grpc.Server
	if os.Getenv("TLS") != "" {
		BackendCert, _ := ioutil.ReadFile("/etc/certs/tls.crt")
		BackendKey, _ := ioutil.ReadFile("/etc/certs/tls.key")

		cert, err := tls.X509KeyPair(BackendCert, BackendKey)
		if err != nil {
			log.Fatalf("failed to parse certificate: %v", err)
		}
		creds := credentials.NewServerTLSFromCert(&cert)
		srv = grpc.NewServer(grpc.Creds(creds))
		fmt.Println("tls")
	} else {
		srv = grpc.NewServer()
	}
	svc := newService(*sqlMode, SQLConn{
		SQLAddr: *sqlAddr,
		SQLUser: *sqlUsername,
		SQLPass: *sqlPassword,
		SQLDB:   *sqlDatabase,
	})
	pb.RegisterCRUDServiceServer(srv, svc)
	srv.Serve(ln)
}
