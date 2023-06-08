package service

import (
	"fmt"

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

func (chat *ChatService) ChatStream(req *pb.ChatRequest, stream pb.ChatService_ChatStreamServer) error {
  chatConfig := chatcompletionstream.ChatCompletionConfigInputDTO{
    Model:                chat.ChatConfigStream.Model,
		ModelMaxTokens:       chat.ChatConfigStream.ModelMaxTokens,
		Temperature:          chat.ChatConfigStream.Temperature,
		TopP:                 chat.ChatConfigStream.TopP,
		N:                    chat.ChatConfigStream.N,
		Stop:                 chat.ChatConfigStream.Stop,
		MaxTokens:            chat.ChatConfigStream.MaxTokens,
		InitialSystemMessage: chat.ChatConfigStream.InitialSystemMessage,
  }

  input := chatcompletionstream.ChatCompletionInputDTO{
    UserMessage:  req.GetUserMessage(),
    UserID:       req.GetUserId(),
    ChatID:       req.GetChatId(),
    Config:       chatConfig,
  }

  ctx := stream.Context()
  go func () {
    for msg := range chat.StreamChannel {
 
      fmt.Println(msg)

      stream.Send(&pb.ChatResponse{
        ChatId : msg.ChatID,
        UserId : msg.UserID,
        Content: msg.Content,
      })
    }
  }()

  
  _,err := chat.ChatCompletionStreamUseCase.Execute(ctx, input)
  if err != nil { return err }

  return nil
}
