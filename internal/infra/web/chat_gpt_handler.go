package web

import "github.com/Samuel-Ricardo/GPT-Chat_Service/internal/domain/usecase/chatcompletion"

type WebChatGPTHandler struct {
	CompletionUseCase chatcompletion.ChatcompletionUseCase
  Config chatcompletion.ChatCompletionConfigInputDTO
  AuthToken string
}


