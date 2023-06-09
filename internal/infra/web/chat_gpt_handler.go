package web

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Samuel-Ricardo/GPT-Chat_Service/internal/domain/usecase/chatcompletion"
)

type WebChatGPTHandler struct {
	CompletionUseCase chatcompletion.ChatcompletionUseCase
  Config chatcompletion.ChatCompletionConfigInputDTO
  AuthToken string
}

func NewWebChatGPTHandler(usecase chatcompletion.ChatcompletionUseCase, config chatcompletion.ChatCompletionConfigInputDTO, authToken string) *WebChatGPTHandler {
  return &WebChatGPTHandler{
    CompletionUseCase: usecase,
    Config: config,
    AuthToken: authToken,
  }
}


func (handler *WebChatGPTHandler) Handle(w http.ResponseWriter, r *http.Request) {
  
  if r.Method != http.MethodPost {
    w.WriteHeader(http.StatusMethodNotAllowed)
    return
  }

  if r.Header.Get("Authorization") != handler.AuthToken {
    w.WriteHeader(http.StatusUnauthorized)
    return
  }


  body, err := ioutil.ReadAll(r.Body)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  if !json.Valid(body) {
    http.Error(w, "Invalid JSON", http.StatusBadRequest)
    return
  }


  var dto chatcompletion.ChatCompletionInputDTO
  err = json.Unmarshal(body, &dto)
    
  if err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
		return
  }
  
  dto.Config = handler.Config

  result, err := handler.CompletionUseCase.Execute(r.Context(), dto)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
  
  w.WriteHeader(http.StatusOK)
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(result)
}
