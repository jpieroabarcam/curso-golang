package handlers

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"rpsweb/rps"
	"strconv"
)

type Player struct {
	Name string
}

var player Player

const (
	templateDir  = "templates/"
	templateBase = templateDir + "base.html"
)

func Index(w http.ResponseWriter, r *http.Request) {
	restartValue()
	renderTemplate(w, "index.html", nil)

	//tpl, err := template.ParseFiles("templates/index.html", "templates/base.html")
	// if err != nil {
	// 	http.Error(w, "Error al analizar plantillas", http.StatusInternalServerError)
	// }

	// data := struct {
	// 	Title   string
	// 	Message string
	// }{
	// 	Title:   "Pagina de Inicio",
	// 	Message: "¡Bienvenido a Piedra, Papel o Tijera!",
	// }

	//fmt.Fprintln(w, "Pagina de inicio")
}

func NewGame(w http.ResponseWriter, r *http.Request) {
	restartValue()
	renderTemplate(w, "new-game.html", nil)
}

func Game(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error parsing form", http.StatusBadRequest)
		}
		player.Name = r.Form.Get("name")
	}

	if player.Name == "" {
		http.Redirect(w, r, "/new", http.StatusFound)
	}

	renderTemplate(w, "game.html", player)
}

func Play(w http.ResponseWriter, r *http.Request) {
	playerChoise, _ := strconv.Atoi(r.URL.Query().Get("c"))
	result := rps.PlayRound(playerChoise)

	out, err := json.MarshalIndent(result, "", "    ")

	if err != nil {
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)

}

func About(w http.ResponseWriter, r *http.Request) {
	restartValue()
	renderTemplate(w, "about.html", nil)
}

func renderTemplate(w http.ResponseWriter, page string, data any) {
	tpl := template.Must(template.ParseFiles(templateBase, templateDir+page))

	erro := tpl.ExecuteTemplate(w, "base", data)
	if erro != nil {
		http.Error(w, "Error al renderizar plantillas", http.StatusInternalServerError)
		log.Println(erro)
		return
	}
}

// Reiniciar valores
func restartValue() {
	player.Name = ""
	rps.ComputerScore = 0
	rps.PlayerScore = 0
}
