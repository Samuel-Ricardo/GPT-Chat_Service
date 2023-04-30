package server

import (
	"github.com/Samuel-Ricardo/GPT-Chat_Service/internal/domain/usecase/chatcompletionstream"
	"github.com/Samuel-Ricardo/GPT-Chat_Service/internal/infra/grpc/service"
)

type GRPCServer struct {
	ChatCompletionStreamUseCase chatcompletionstream.ChatCompletionUseCase
  ChatConfig                  chatcompletionstream.ChatCompletionConfigInputDTO
  ChatService                 service.ChatService
  Port                        string
  AuthToken                   string
  StreamChannel               chan chatcompletionstream.ChatCompletionOutputDTO
}
