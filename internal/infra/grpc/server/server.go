package server

import (
	"net"

	"github.com/Samuel-Ricardo/GPT-Chat_Service/internal/domain/usecase/chatcompletionstream"
	"github.com/Samuel-Ricardo/GPT-Chat_Service/internal/infra/grpc/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type GRPCServer struct {
	ChatCompletionStreamUseCase chatcompletionstream.ChatCompletionUseCase
  ChatConfigStream            chatcompletionstream.ChatCompletionConfigInputDTO
  ChatService                 service.ChatService
  Port                        string
  AuthToken                   string
  StreamChannel               chan chatcompletionstream.ChatCompletionOutputDTO
}

func NewGRPCServer(chatCompletionStreamUseCase chatcompletionstream.ChatCompletionUseCase, chatConfigStream chatcompletionstream.ChatCompletionConfigInputDTO, port, authToken string, streamChannel chan chatcompletionstream.ChatCompletionOutputDTO) *GRPCServer {
	chatService := service.NewChatService(chatCompletionStreamUseCase, chatConfigStream, streamChannel)
	return &GRPCServer{
		ChatCompletionStreamUseCase: chatCompletionStreamUseCase,
		ChatConfigStream:            chatConfigStream,
		Port:                        port,
		AuthToken:                   authToken,
		StreamChannel:               streamChannel,
		ChatService:                 *chatService,
	}
}

func (g *GRPCServer) AuthInterceptor(server interface{}, serverStream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
  ctx := serverStream.Context()
  
  md, ok := metadata.FromIncomingContext(ctx)
  if !ok {
		return status.Error(codes.Unauthenticated, "metadata is not provided")
	}


  token := md.Get("authorization")
  if len(token) == 0 {
		return status.Error(codes.Unauthenticated, "authorization token is not provided")
	}


  if token[0] != g.AuthToken {
    return status.Error(codes.Unauthenticated, "authorization token is invalid")
  }

  return handler(server, serverStream)
}

func (g *GRPCServer) Start() error {
  opts := []grpc.ServerOption{
    grpc.StreamInterceptor(g.AuthInterceptor),
  }

  grpcServer = grpc.NewServer(opts...)
  pb.RegisterChatServiceServer(grpcServer.ChatService)
  reflection.Register(grpcServer)

  listen, err := net.Listen("tcp", ":"+g.Port)
  if err != nil { panic(err.Error()) }

  if err := grpcServer.Server(listen); err != nil { panic(err.Error()) }

  return nil
} 
