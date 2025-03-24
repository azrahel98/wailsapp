package models

import (
	"github.com/sashabaranov/go-openai"
)

type FuncionHandler func(toolCall *openai.ToolCall, chats []openai.ChatCompletionMessage) (*string, error)

type FuncionIa struct {
	Nombre  string
	Handler FuncionHandler
}
