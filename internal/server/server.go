package server

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"

	pb "github.com/GoingFast/test6/protobuf"
	"github.com/golang/protobuf/ptypes/empty"
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
			"2": &pb.Request{Message: "ba"},
			"3": &pb.Request{Message: "addf"},
			"4": &pb.Request{Message: "fgh"},
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

func NewService(sql bool, s SQLConn) service {
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
