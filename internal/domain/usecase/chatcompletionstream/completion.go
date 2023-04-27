package chatcompletionstream

import (
	"errors"

	"github.com/Samuel-Ricardo/GPT-Chat_Service/internal/domain/entity"
	"github.com/Samuel-Ricardo/GPT-Chat_Service/internal/domain/gateway"
	"github.com/sashabaranov/go-openai"
)

type ChatCompletionConfigInputDTO struct {
	Model                string
	ModelMaxTokenx       int
	Temperature          float32
	TopP                 float32
	N                    int
	Stop                 []string
	MaxTokens            int
	PresencePenalty      float32
	FrequencyPenalty     float32
	InitialSystemMessage string
}

type ChatCompletionInputDTO struct {
	ChatID      string
	UserID      string
	UserMessage string
	Config      ChatCompletionConfigInputDTO
}

type ChatCompletionOutputDTO struct {
	ChatID  string
	UserID  string
	Content string
}

type ChatCompletionUseCase struct {
	ChatGateway gateway.ChatGateway
  OpenAIClient *openai.Client
  Stream chan ChatCompletionOutputDTO
}

func NewChatCompletionUseCase(chatGateway gateway.ChatGateway, openAiClient *openai.Client, stream chan ChatCompletionOutputDTO) *ChatCompletionUseCase {
  return &ChatCompletionUseCase{
    ChatGateway: chatGateway,
    OpenAIClient: openAiClient,
    Stream: stream,
  }
}

func createNewChat(input ChatCompletionInputDTO) (*entity.Chat, error) {
  model := entity.NewModel(input.Config.Model, input.Config.ModelMaxTokenx)

  chatConfig := &entity.ChatConfig{
    Temperature:      input.Config.Temperature,
		TopP:             input.Config.TopP,
		N:                input.Config.N,
		Stop:             input.Config.Stop,
		MaxTokens:        input.Config.MaxTokens,
		PresencePenalty:  input.Config.PresencePenalty,
		FrequencyPenalty: input.Config.FrequencyPenalty,
		Model:            model,
  }

  
  initialMessage, err := entity.NewMessage("system", input.Config.InitialSystemMessage, model) 

  if err != nil { 
    return nil, errors.New("error creating initial message: " + err.Error())
  }

  chat, err := entity.NewChat(input.UserID, initialMessage, chatConfig)

  if err != nil {
    return nil, errors.New("error creating new chat: " + err.Error())
  }

  return chat, nil
}
