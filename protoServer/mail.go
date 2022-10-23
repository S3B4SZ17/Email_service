package protoServer

import (
	"context"

	"github.com/S3B4SZ17/Email_service/management"
	pbEmail "github.com/S3B4SZ17/Email_service/proto/email_user"
	"github.com/S3B4SZ17/Email_service/services"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
)

type EmailServer struct {
	pbEmail.UnimplementedSendEmailServer
}

type GetAuthenticatedUser struct {
	pbEmail.UnimplementedGetAuthenticatedUserServer
}

func (s *EmailServer) SendEmail(ctx context.Context, in *pbEmail.EmailMessage) (res *pbEmail.EmailResponse, err error) {
	management.Log.Info("Received: email =>", zap.String("email", in.Email))

	res, err = services.SendEmail(in)
	if err != nil {
		management.Log.Error("An error occurred while sending the email", zap.String("error", err.Error()))
		return
	}

	return
}

func (s *GetAuthenticatedUser) GetAuthenticatedUser(ctx context.Context, empty *emptypb.Empty) (res *pbEmail.EmailUser, err error) {

	// res, err = services.GetAuthenticatedUser()
	if err != nil {
		management.Log.Error("An error occurred while retrieving user", zap.String("error", err.Error()))
		return
	}

	return
}
