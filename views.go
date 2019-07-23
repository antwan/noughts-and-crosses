package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

var (
    game *Noughts
    currentPlayer int
)

func initGame() {
    game = new(Noughts)
    currentPlayer = 1
}

// Reset resets the game
func Reset(w http.ResponseWriter, r *http.Request) {
    initGame()
    fmt.Fprintln(w, "Game reintialized")
}

// Status displays the current game status
func Status(w http.ResponseWriter, r *http.Request) {
    if game == nil {
        initGame()
    }
    fmt.Fprintln(w, game)

    fmt.Fprintf(w, "Next player is: Player %v\n", currentPlayer)

    winner := game.GetWinner()
    if winner != 0 {
        fmt.Fprintf(w, "** PLAYER %v WINS THE GAME **\n", winner)
    } else {
        fmt.Fprintf(w, "No winner currently\n")
    }
}

// PlayData is the structure expected to be passed when playing
type PlayData struct {
    Row int `json:"row"`
    Col int `json:"col"`
}

// Play allows current player to play a move and displays game
func Play(w http.ResponseWriter, r *http.Request) {
    if game == nil {
        initGame()
    }

    // Retrieves and parse data
    var data PlayData
    decoder := json.NewDecoder(r.Body)
    decoder.DisallowUnknownFields()

    if err := decoder.Decode(&data); err != nil {
        w.WriteHeader(http.StatusBadRequest)
        fmt.Fprintln(w, err)
        return
    }

    // Records the move if possible
    if err := game.Play(data.Row - 1, data.Col - 1, currentPlayer); err != nil {
        w.WriteHeader(http.StatusBadRequest)
        fmt.Fprintln(w, err)
        return
    }

    // Displays game and changes player
    fmt.Fprintln(w, game)
    fmt.Fprintf(w, "Move recorded at position %vx%v for player %v\n", data.Row, data.Col, currentPlayer)
    currentPlayer = 3 - currentPlayer

    winner := game.GetWinner()
    if winner != 0 {
        fmt.Fprintf(w, "** PLAYER %v WINS THE GAME **\n", winner)
    } else {
        fmt.Fprintf(w, "No winner currently\n")
    }

}
