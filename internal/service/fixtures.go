package service

// GetDataRequest is a fixture to create a request for GetData method.
func GetDataRequestFixture() *GetDataRequest {
	return &GetDataRequest{}
}

// GetDataResponse is a fixture to create a response for GetData method.
func GetDataResponseFixture() *GetDataResponse {
	return &GetDataResponse{
		Data: "expected data",
	}
}