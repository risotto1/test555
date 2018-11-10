package client

import (
	"context"

	pb "github.com/GoingFast/test6/protobuf"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

type mockConns struct {
	OnRead func(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*pb.ReadResponse, error)
}

func (m mockConns) Read(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*pb.ReadResponse, error) {
	return m.OnRead(ctx, in)
}

func TestRg.T) {
// tests := []struct {
	// 	wantBody   interface{}
	// 	wantStatus int
	// }{
	// 	{
	// 		map[string][]map[string]interface{}{
	// 			"msg": []map[string]interface{}{
	// 				map[string]interface{}{
	// 					"message": "hello",
	// 				},
	// 				map[string]interface{}{},
	// 			},
	// 		},
	// 		200,
	// 	},
	// }

	// for _, test := range tests {
	// 	s := service{conns{
	// 		mockConns{
	// 			OnRead: func(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*pb.ReadResponse, error) {
	// 				return &pb.ReadResponse{Data: []*pb.Request{&pb.Request{Message: "hello"}}}, nil

	// 			},
	// 		},
	// 	}}
	// 	r, err := http.NewRequest("GET", "/", nil)
	// 	if err != nil {
	// 		t.Error(err)
	// 	}
	// 	w := httptest.NewRecorder()
	// 	handler := http.HandlerFunc(s.Read())
	// 	handler.ServeHTTP(w, r)

	// 	if w.Code != test.wantStatus {
	// 		t.Errorf("\ngot status: %v\nwant status: %v", w.Code, test.wantStatus)
	// 	}
	// 	got := make(map[string][]map[string]interface{})
	// 	err = json.NewDecoder(w.Body).Decode(&got)
	// 	if err != nil {
	// 		t.Error(err)
	// 	}
	// 	ok := reflect.DeepEqual(got, test.wantBody)
	// 	if !ok {
	// 		t.Errorf("\ngot body: %v\nwant body: %v", got, test.wantBody)
	// 	}
	// }
}
