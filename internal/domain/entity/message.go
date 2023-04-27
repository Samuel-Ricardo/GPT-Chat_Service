package entity

import (
	"time"

	"github.com/google/uuid"
	tiktoken_go "github.com/j178/tiktoken-go"
)
type Message struct {
	ID        string
	Role      string
	Content   string
	Tokens    int
	Model     *Model
	CreatedAt time.Time
}


func NewMessage(role, content string, model *Model) (*Message, error) {
  totalTokens := tiktoken_go.CountTokens(model.GetModelName(), content)
  msg := &Message {
    ID: uuid.New().String(),
    Role: role,
    Content: content,
    Tokens: totalTokens,
    Model: model,
    CreatedAt: time.Now(),
  }

  if err := msg.Validate(); err != nil { return nil,err }

  return msg, nil
}
