package main

import (
	"database/sql"
	"fmt"

	"github.com/Samuel-Ricardo/GPT-Chat_Service/configs"
	"github.com/Samuel-Ricardo/GPT-Chat_Service/internal/domain/usecase/chatcompletion"
	"github.com/Samuel-Ricardo/GPT-Chat_Service/internal/domain/usecase/chatcompletionstream"
	"github.com/Samuel-Ricardo/GPT-Chat_Service/internal/infra/repository"
	"github.com/Samuel-Ricardo/GPT-Chat_Service/internal/infra/web"
  "github.com/Samuel-Ricardo/GPT-Chat_Service/internal/infra/web/webserver"
	"github.com/sashabaranov/go-openai"
)

func main() {

  configs, err := configs.LoadConfig(".")
  if err != nil { panic(err) }

  connection, err := sql.Open(configs.DBDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&multiStatements=true", configs.DBUser, configs.DBPassword, configs.DBHost, configs.DBPort, configs.DBName)) 
  if err != nil { panic(err) }

  defer connection.Close()

  repo := repository.NewChatRepositoryMySQL(connection)
  client := openai.NewClient(configs.OpenAIApiKey)

  chatConfig := chatcompletion.ChatCompletionConfigInputDTO {
    Model:                configs.Model,
		ModelMaxTokens:       configs.ModelMaxTokens,
		Temperature:          float32(configs.Temperature),
		TopP:                 float32(configs.TopP),
		N:                    configs.N,
		Stop:                 configs.Stop,
		MaxTokens:            configs.MaxTokens,
		InitialSystemMessage: configs.InitialChatMessage,
  }

  chatConfigStrem := chatcompletionstream.ChatCompletionConfigInputDTO{
    Model:                configs.Model,
		ModelMaxTokens:       configs.ModelMaxTokens,
		Temperature:          float32(configs.Temperature),
		TopP:                 float32(configs.TopP),
		N:                    configs.N,
		Stop:                 configs.Stop,
		MaxTokens:            configs.MaxTokens,
		InitialSystemMessage: configs.InitialChatMessage,
  }

  usecase := chatcompletion.NewChatCompletionUseCase(repo,client)

  streamChannel := make(chan chatcompletionstream.ChatCompletionOutputDTO)
  usecaseStream := chatcompletionstream.NewChatCompletionUseCase(repo,client,streamChannel)

  //fmt.Println("Starting gRPC server on port "+ configs.GRPCServerPort)

  webServer := webserver.NewWebServer(":" + configs.WebServerPort)
  webServerChatHandler := web.NewWebChatGPTHandler(*usecase, chatConfig, configs.AuthToken)

  webServer.AddHandler("/chat", webServerChatHandler.Handle)

  fmt.Println("Server Runing on Port "+ configs.WebServerPort)
  webServer.Start()
}
