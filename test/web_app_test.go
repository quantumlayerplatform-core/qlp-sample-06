package service

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	RegisterWebAppServiceServer(s, &mockWebAppServer{})
	go func() {
		if err := s.Serve(lis); err != nil {
			panic("Server exited with error: " + err.Error())
		}
	}()
}

func bufDialer(context.Context, string) (grpc.ClientConnInterface, error) {
	return grpc.DialContext(context.Background(), "", grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
		return lis.Dial()
	}), grpc.WithInsecure())
}

func TestWebAppService_GetData(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()

	client := NewWebAppServiceClient(conn)

	resp, err := client.GetData(ctx, &GetDataRequest{})
	if err != nil {
		t.Fatalf("GetData failed: %v", err)
	}

	expected := "expected data"
	if resp.Data != expected {
		t.Errorf("Expected data %q, got %q", expected, resp.Data)
	}
}

type mockWebAppServer struct {
	UnimplementedWebAppServiceServer
}

func (m *mockWebAppServer) GetData(ctx context.Context, req *GetDataRequest) (*GetDataResponse, error) {
	return &GetDataResponse{Data: "expected data"}, nil
}