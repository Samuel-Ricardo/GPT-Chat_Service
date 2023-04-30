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

func NewChatService(chatCompletionStreamUseCase chatcompletionstream.ChatCompletionUseCase, chatConfigStream chatcompletionstream.ChatCompletionConfigInputDTO, streamChannel chan chatcompletionstream.ChatCompletionOutputDTO) *ChatService {
	return &ChatService{
		ChatCompletionStreamUseCase: chatCompletionStreamUseCase,
		ChatConfigStream:            chatConfigStream,
		StreamChannel:               streamChannel,
	}
}
