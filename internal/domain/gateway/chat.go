package gateway

import (
	"context"

	"github.com/Samuel-Ricardo/GPT-Chat_Service/internal/domain/entity"
)

type ChatGateway interface {
	CreateChat(ctx context.Context, chat *entity.Chat) error
  FindChatByID(ctx context.Context, chatID string) (*entity.Chat, error)
  SaveChat(ctx context.Context, chat *entity.Chat) error
}
