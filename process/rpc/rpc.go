package rpc

import (
	"context"
	"net"
	"zrw/goMin/process/controller"
	"zrw/goMin/process/rpc/server"
)
import "google.golang.org/grpc"

type Server struct{}

// StartRpcServer 启动rpc服务
func StartRpcServer(addr string) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
	}
	s := grpc.NewServer()
	server.RegisterServerServer(s, &Server{})
	if err := s.Serve(lis); err != nil {
	}
}

func (rp *Server) GetServerTime(ctx context.Context, request *server.ServerTimeRequest) (*server.ServerTimeResponse, error) {
	data, code, msg := controller.GetServerTime()
	resp := &server.ServerTimeResponse{}
	respData := &server.ServerTimeResponseData{}
	resp.Msg = msg
	resp.Code = uint32(code)
	respData.ServerTime = uint64(data.ServerTime)
	resp.Data = respData
	return resp, nil
}
