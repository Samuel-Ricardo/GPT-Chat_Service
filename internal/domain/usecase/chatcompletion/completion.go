package chatcompletion

import (
	"errors"

	"github.com/Samuel-Ricardo/GPT-Chat_Service/internal/domain/entity"
	"github.com/Samuel-Ricardo/GPT-Chat_Service/internal/domain/gateway"
	"github.com/sashabaranov/go-openai"
)

type ChatCompletionConfigInputDTO struct {
	Model                string
	ModelMaxTokens       int
	Temperature          float32  // 0.0 to 1.0
	TopP                 float32  // 0.0 to 1.0 - to a low value, like 0.1, the model will be very conservative in its word choices, and will tend to generate relatively predictable prompts
	N                    int      // number of messages to generate
	Stop                 []string // list of tokens to stop on
	MaxTokens            int      // number of tokens to generate
	PresencePenalty      float32  // -2.0 to 2.0 - Number between -2.0 and 2.0. Positive values penalize new tokens based on whether they appear in the text so far, increasing the model's likelihood to talk about new topics.
	FrequencyPenalty     float32  // -2.0 to 2.0 - Number between -2.0 and 2.0. Positive values penalize new tokens based on their existing frequency in the text so far, increasing the model's likelihood to talk about new topics.
	InitialSystemMessage string
}

type ChatCompletionInputDTO struct {
	ChatID      string                       `json:"chat_id,omitempty"`
	UserID      string                       `json:"user_id"`
	UserMessage string                       `json:"user_message"`
	Config      ChatCompletionConfigInputDTO `json:"Config"`
}

type ChatCompletionOutputDTO struct {
	ChatID  string `json:"chat_id"`
	UserID  string `json:"user_id"`
	Content string `json:"content"`
}

type ChatcompletionUseCase struct {
	ChatGateway gateway.ChatGateway
  OpenAIClient *openai.Client
}

func NewChatCompletionUseCase (chatGateway gateway.ChatGateway, openAIClient *openai.Client) *ChatcompletionUseCase {
  return &ChatcompletionUseCase {
    ChatGateway: chatGateway,
    OpenAIClient: openAIClient,
  }
}


func createNewChat(input ChatCompletionInputDTO) (*entity.Chat, error) {
  model := entity.NewModel(input.Config.Model, input.Config.ModelMaxTokens)
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

  initalmessage, err := entity.NewMessage("system", input.Config.InitialSystemMessage, model)
  if err != nil {
		return nil, errors.New("error creating initial message: " + err.Error())
	}

  chat, err := entity.NewChat(input.UserID, initalmessage, chatConfig)
  if err != nil {
		return nil, errors.New("error creating new chat: " + err.Error())
	}

  return chat, nil;
}

