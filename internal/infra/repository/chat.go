package repository

import (
	"database/sql"

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
