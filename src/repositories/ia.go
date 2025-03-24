package repositories

import (
	"context"
	"fmt"
	"vesgoapp/src/models"

	"github.com/jmoiron/sqlx"
	"github.com/sashabaranov/go-openai"
)

type IaRepository interface {
	Crear_Query(prompt string, chats []openai.ChatCompletionMessage) (*openai.ChatCompletionResponse, error)
	ProcessarRespuesta(respuesta *openai.ChatCompletionResponse) (*string, *openai.ToolCall, error)
}

type iaRepository struct {
	db *sqlx.DB
}

func (i *iaRepository) Crear_Query(prompt string, chats []openai.ChatCompletionMessage) (*openai.ChatCompletionResponse, error) {
	client := openai.NewClient(models.CHATGPKEY)

	historial := append(chats, openai.ChatCompletionMessage{Role: "system", Content: " "})

	resp, err := client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model:       openai.GPT4oMini,
		Temperature: 0,
		Messages:    historial,
		Tools: []openai.Tool{
			{
				Type: openai.ToolTypeFunction,
				Function: &openai.FunctionDefinition{
					Name:        "eliminar_usuario",
					Description: "Eliminar un usuario de la base de datos usando su DNI",
					Parameters: map[string]any{
						"type": "object",
						"properties": map[string]any{
							"dni": map[string]any{
								"type":        "string",
								"description": "El DNI del usuario a eliminar",
								"pattern":     "^[0-9]{8}$",
							},
						},
						"required": []string{"dni"},
					},
				},
			},
			{
				Type: openai.ToolTypeFunction,
				Function: &openai.FunctionDefinition{
					Name:        "actualizar_usuario",
					Description: "Actualizar la información de un usuario en la base de datos usando su DNI",
					Parameters: map[string]any{
						"type": "object",
						"properties": map[string]any{
							"dni": map[string]any{
								"type":        "string",
								"description": "El DNI del usuario a actualizar",
								"pattern":     "^[0-9]{8}$",
							},
							"direccion": map[string]any{
								"type":        "string",
								"description": "La nueva dirección del usuario (opcional)",
							},
							"correo": map[string]any{
								"type":        "string",
								"description": "El nuevo correo electrónico del usuario (opcional)",
								"format":      "email",
							},
							"fecha_de_nacimiento": map[string]any{
								"type":        "string",
								"description": "La nueva fecha de nacimiento del usuario en formato YYYY-MM-DD (opcional)",
								"format":      "date",
							},
							"telefono": map[string]any{
								"type":        "string",
								"description": "El nuevo número de teléfono del usuario (opcional)",
								"pattern":     "^[0-9]+$",
							},
						},
						"required": []string{"dni", "correo"},
					},
				},
			},
		},
	})

	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (i *iaRepository) ProcessarRespuesta(respuesta *openai.ChatCompletionResponse) (*string, *openai.ToolCall, error) {
	if len(respuesta.Choices) == 0 || respuesta.Choices[0].Message.Content != "" {
		return &respuesta.Choices[0].Message.Content, nil, nil
	}

	if len(respuesta.Choices[0].Message.ToolCalls) > 0 {
		toolCall := respuesta.Choices[0].Message.ToolCalls[0]
		return nil, &toolCall, nil
	}

	return nil, nil, fmt.Errorf("respuesta no válida de la IA")
}

func CreateRepoIa(db *sqlx.DB) IaRepository {
	return &iaRepository{db}
}
