package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/Samuel-Ricardo/GPT-Chat_Service/internal/domain/entity"
	"github.com/Samuel-Ricardo/GPT-Chat_Service/internal/infra/db"
)


type ChatRepositoryMySQL struct {
	DB *sql.DB
  Queries *db.Queries
}


func NewChatRepositoryMySQL(database *sql.DB) *ChatRepositoryMySQL {
  return &ChatRepositoryMySQL{
    DB: database,
    Queries: db.New(database),
  }
}

func (r *ChatRepositoryMySQL) CreateChat(ctx context.Context, chat *entity.Chat) error {
  err := r.Queries.CreateChat(
    ctx,
    db.CreateChatParams{
      ID:               chat.ID,
			UserID:           chat.UserID,
			InitialMessageID: chat.InitialSystemMessage.Content,
			Status:           chat.Status,
			TokenUsage:       int32(chat.TokenUsage),
			Model:            chat.Config.Model.Name,
			ModelMaxTokens:   int32(chat.Config.Model.MaxTokens),
			Temperature:      float64(chat.Config.Temperature),
			TopP:             float64(chat.Config.TopP),
			N:                int32(chat.Config.N),
			Stop:             chat.Config.Stop[0],
			MaxTokens:        int32(chat.Config.MaxTokens),
			PresencePenalty:  float64(chat.Config.PresencePenalty),
			FrequencyPenalty: float64(chat.Config.FrequencyPenalty),
			CreatedAt:        time.Now(),
			UpdatedAt:        time.Now(),
		},
  )

  if err != nil { return err }

  err = r.Queries.AddMessage(
    ctx,
    db.AddMessageParams{
			ID:        chat.InitialSystemMessage.ID,
			ChatID:    chat.ID,
			Content:   chat.InitialSystemMessage.Content,
			Role:      chat.InitialSystemMessage.Role,
			Tokens:    int32(chat.InitialSystemMessage.Tokens),
			CreatedAt: chat.InitialSystemMessage.CreatedAt,
		},
  )

  if err != nil {return err}

  return nil
}
