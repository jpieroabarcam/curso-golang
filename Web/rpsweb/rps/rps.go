package rps

import (
	"math/rand"
	"strconv"
)

const (
	ROCK     = 0 // Piedra. Vence a las tijeras. (tijeras +1) % 3 = 0
	PAPER    = 1 // Papel. Vence a la piedra. (piedra +1) % 3 = 1
	SCISSORS = 2 // Tijeras. Vence al papel. (papel +1) % 3 = 2
)

// Estructura para dar resultado de cada ruta
type Round struct {
	Message           string `json:"message"`
	ComputerChoise    string `json:"computer_choise"`
	RoundResult       string `json:"round_result"`
	ComputerChoiseInt int    `json:"computer_choise_int"`
	ComputerScore     string `json:"computer_score"`
	PlayerScore       string `json:"player_score"`
}

// Mensajes para ganar
var winMessages = []string{
	"¡Bien Hecho!",
	"¡Buen Trabajo",
	"Deberias comprar un boleto de loteria",
}

// Mensajes perdedor
var loseMessages = []string{
	"¡Que lastima!",
	"¡Intentalo de nuevo!",
	"Hoy simplemente no est tu dia",
}

// Mensajes empate
var drawMessages = []string{
	"Las grnades mentes piensan igual",
	"Oh oh. Intentalo de nuevo",
	"Nadie gana, pero puedes intentarlo de nuevo",
}

// variables para el puntaje
var ComputerScore, PlayerScore int

func PlayRound(playerValue int) Round {
	computerValue := rand.Intn(3)

	var computerChoice, roundResult string
	var computerChoiceInt int

	// Mensajes dependiendo de lo que eligio la computadora
	switch computerValue {
	case ROCK:
		computerChoiceInt = ROCK
		computerChoice = "La computadora eligio PIEDRA"
	case PAPER:
		computerChoiceInt = PAPER
		computerChoice = "La computadora eligio PAPEL"
	case SCISSORS:
		computerChoiceInt = SCISSORS
		computerChoice = "La computadora eligio TIJERAS"
	}

	//generar un numero aleatorio de 0-2, que usamos para elegir un mensaje aleatorio
	messageInt := rand.Intn(3)

	// declarar una var para contener el mensaje
	var message string
	if playerValue == computerValue {
		roundResult = "Es un empate"
		// seleccionar mensaje de drawMessages
		message = drawMessages[messageInt]
	} else if playerValue == (computerValue+1)%3 {
		PlayerScore++
		roundResult = "¡El jugador gana!"
		//seleccionar mensaje de winMessages
		message = winMessages[messageInt]
	} else {
		ComputerScore++
		roundResult = "¡La computadora gana!"
		//seleccionar mensaje de loseMessages
		message = loseMessages[messageInt]
	}

	//Retornar resultado
	return Round{
		Message:           message,
		ComputerChoise:    computerChoice,
		RoundResult:       roundResult,
		ComputerChoiseInt: computerChoiceInt,
		ComputerScore:     strconv.Itoa(ComputerScore),
		PlayerScore:       strconv.Itoa(PlayerScore),
	}

}
