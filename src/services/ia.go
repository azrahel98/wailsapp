package services

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"vesgoapp/src/repositories"

	"github.com/sashabaranov/go-openai"
)

type IaService struct {
	repo repositories.IaRepository
	dash repositories.DashboardRepository
	pers repositories.PersonalRepository
}

func NewIaRepository(r repositories.IaRepository, d repositories.DashboardRepository, p repositories.PersonalRepository) *IaService {
	return &IaService{repo: r, dash: d, pers: p}
}

func (s *IaService) Consulta_IA(prompt string, chatHistory []openai.ChatCompletionMessage, isAdmin bool) (*string, error) {

	chatHistory = append(chatHistory, openai.ChatCompletionMessage{
		Role: openai.ChatMessageRoleAssistant,
		Content: `Eres un asistente de una aplicación de recursos humanos. Instrucciones:
		1. Transforma información técnica en respuestas amigables.
		2. siempre usa HTML  para estructurar las respuestas.
		3. Mantén un tono cordial y servicial.
		4. Usa Bootstrap para tablas (class="table table-striped").
		5. Cuando se te pase un mensaje de error , humaniza ese error y devuelve un texto corto`,
	}, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: prompt,
	})

	respuesta, err := s.repo.Crear_Query(prompt, chatHistory)
	if err != nil {
		return nil, fmt.Errorf("error al consultar la IA: %w", err)
	}
	fmt.Println(respuesta)

	contenido, toolCall, err := s.repo.ProcessarRespuesta(respuesta)
	if err != nil {
		return nil, fmt.Errorf("error al procesar respuesta: %w", err)
	}

	if contenido != nil {
		return contenido, nil
	}

	if toolCall != nil {
		if !s.usuarioConfirmo(chatHistory, toolCall.Function.Name) {
			mensajeConfirmacion := "¿Estás seguro de que quieres ejecutar la acción? Responde 'Sí' para confirmar."
			return &mensajeConfirmacion, nil
		}
		return s.EjecutarFuncion(toolCall, isAdmin)
	}

	return nil, fmt.Errorf("respuesta no válida de la IA")
}

func (s *IaService) usuarioConfirmo(chatHistory []openai.ChatCompletionMessage, functionName string) bool {
	for _, msg := range chatHistory {
		if msg.Role == openai.ChatMessageRoleUser && strings.Contains(strings.ToLower(msg.Content), "si") {
			return true
		}
	}
	return false
}

func (s *IaService) EjecutarFuncion(toolCall *openai.ToolCall, isAdmin bool) (*string, error) {
	var args map[string]string
	if err := json.Unmarshal([]byte(toolCall.Function.Arguments), &args); err != nil {
		return nil, err
	}

	switch toolCall.Function.Name {
	case "eliminar_usuario":
		if !isAdmin {
			return nil, fmt.Errorf("permiso denegado: solo los administradores pueden eliminar usuarios")
		}
		return s.manejarEliminacionUsuario(args)

	case "actualizar_usuario":
		if !isAdmin {
			return nil, fmt.Errorf("permiso denegado: los administradores no pueden actualizar usuarios")
		}
		return s.manejarActualizarUsuario(args)

	default:
		return nil, fmt.Errorf("función no soportada: %s", toolCall.Function.Name)
	}
}

func (s *IaService) manejarEliminacionUsuario(args map[string]string) (*string, error) {
	dni := args["dni"]
	fmt.Println(args)
	fmt.Printf("Eliminando usuario con DNI: %s\n", dni)
	respuesta := fmt.Sprintf("Usuario con DNI %s eliminado correctamente.", dni)
	return &respuesta, nil
}

func (s *IaService) manejarActualizarUsuario(args map[string]string) (*string, error) {
	dni := args["dni"]
	telf := args["telefono"]
	correo := args["correo"]
	direccion := args["direccion"]

	err := s.pers.EditByDni(context.Background(), &telf, nil, &direccion, &correo, nil, &dni)
	if err != nil {
		iares, err := s.Consulta_IA(fmt.Sprintf("se intento ejecutar la funcion de actualizar los datos de un usuario por su dni y tuvo este error %s", err), []openai.ChatCompletionMessage{}, true)
		if err != nil {
			res := fmt.Sprintf("Usuario con DNI %s actualizado correctamente.", dni)
			return &res, nil
		}

		return iares, nil
	}
	respuesta := "Usuarioactualizado correctamente"
	return &respuesta, nil
}
