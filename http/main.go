package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	pb "github.com/GoingFast/test6/protobuf"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type conns struct {
	crud pb.CRUDServiceClient
}

func fallbackEnv(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}

func newConns() conns {
	var (
		conn *grpc.ClientConn
		err  error
	)
	if os.Getenv("TLS") != "" {
		creds, err := credentials.NewClientTLSFromFile("/home/certs/tls.crt", "")
		if err != nil {
			log.Fatal(err)
		}
		conn, err = grpc.Dial(fallbackEnv("DIAL_ADDR", ":50051"), grpc.WithTransportCredentials(creds))

	} else {
		conn, err = grpc.Dial(fallbackEnv("DIAL_ADDR", ":50051"), grpc.WithInsecure())
	}
	if err != nil {
		log.Fatal(err)
	}
	client := pb.NewCRUDServiceClient(conn)
	return conns{
		crud: client,
	}
}

type service struct {
	c conns
}

func JSON(w http.ResponseWriter, code int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(&v)
}

func (s service) Read() http.HandlerFunc {
	type res struct {
		Err string `json:"err,omitempty"`
		Msg string `json:"msg,omitempty"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		resp, err := s.c.crud.Read(r.Context(), &empty.Empty{})
		if err != nil {
			JSON(w, 500, res{Err: err.Error()})
			return
		}
		if len(resp.Data) <= 0 {
			JSON(w, http.StatusNotFound, res{Err: "couldn't find any"})
			return
		}

		JSON(w, 200, res{Msg: fmt.Sprintf("server hostname: %s\ngateway hostname: %s", resp.Data, os.Getenv("HOSTNAME"))})
	}
}

func newService() service {
	return service{newConns()}
}
func main() {
	r := chi.NewRouter()
	r.Use(middleware.Heartbeat("/healthz"))

	svc := newService()
	r.Get("/read", svc.Read())
	http.ListenAndServe(fallbackEnv("GATEWAY_LISTEN_ADDR", ":8081"), r)
}
