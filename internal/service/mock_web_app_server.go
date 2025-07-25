package service

import (
	"context"
)

// MockWebAppServer is a mock implementation of the WebAppServiceServer interface to be used in tests.
type MockWebAppServer struct {
	UnimplementedWebAppServiceServer
	DataFunc func(ctx context.Context, req *GetDataRequest) (*GetDataResponse, error)
}

func (m *MockWebAppServer) GetData(ctx context.Context, req *GetDataRequest) (*GetDataResponse, error) {
	if m.DataFunc != nil {
		return m.DataFunc(ctx, req)
	}
	return nil, status.Errorf(codes.Unimplemented, "method GetData not implemented")
}