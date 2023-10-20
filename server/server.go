package main

import "context"
import "github.com/KDias-code/calculate/api"

type server struct {
	__.UnimplementedCalculatorServer
}

func (s *server) Add(ctx context.Context, req *__.AddRequest) (*__.AddResponse, error) {
	result := req.Num1 + req.Num2
	return &__.AddResponse{
		Result: result,
	}, nil
}
