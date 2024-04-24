package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello devuelve un saludo para una persona especifica
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Nombre vacio")
	}
	// Devuelve un saludo que incluye un nombre en un mensaje
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// Hellos revibe varios nombres y devuelve un mapa de nombres
// con un mensaje de saludo
func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		messages[name] = message
	}

	return messages, nil
}

// randomFormat devuelve uno de varios formatos de mensajes de salida
// se selecciona de forma aleatoria
func randomFormat() string {
	formats := []string{
		"¡Hola, %v! ¡Bienvenido!",
		"¡Que bueno verte, %v!",
		"¡Saludos, %v! ¡Encantado de conocerte!",
	}
	return formats[rand.Intn(len(formats))]
}
