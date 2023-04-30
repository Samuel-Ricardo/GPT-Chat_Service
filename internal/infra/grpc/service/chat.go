package service

import (
	"github.com/Samuel-Ricardo/GPT-Chat_Service/internal/domain/usecase/chatcompletionstream"
	"github.com/Samuel-Ricardo/GPT-Chat_Service/internal/infra/grpc/pb"
)

type ChatService struct {
	pb.UnimplementedChatServiceServer
  ChatCompletionStreamUseCase chatcompletionstream.ChatCompletionUseCase
  ChatConfigStream            chatcompletionstream.ChatCompletionConfigInputDTO
  StreamChannel               chan chatcompletionstream.ChatCompletionOutputDTO
}
