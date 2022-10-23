package protoServer

import (
	"net"

	pbEmail "github.com/S3B4SZ17/Email_service/proto/email_user"
	"google.golang.org/grpc"
)

func RegisterServices(s *grpc.Server) {
	pbEmail.RegisterSendEmailServer(s, &EmailServer{})
	pbEmail.RegisterGetAuthenticatedUserServer(s, &GetAuthenticatedUser{})
}

func StartServer(s *grpc.Server, l net.Listener) error {
	return s.Serve(l)
}
