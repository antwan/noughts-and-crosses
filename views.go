package main

import (
    //"encoding/json"
    "fmt"
    "net/http"
    //"github.com/gorilla/mux"
)

var (
    game *Noughts
    currentPlayer int
)

// Start displays a welcome message
func Start(w http.ResponseWriter, r *http.Request) {
    game = new(Noughts)
    currentPlayer = 1

    fmt.Fprintln(w, "Game intialized")
}

// Status displays the current game status
func Status(w http.ResponseWriter, r *http.Request) {
    if game == nil {
        fmt.Fprintln(w, "Game is not started")
        return
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

// Play allows current player to play a move
func Play(w http.ResponseWriter, r *http.Request) {

    if err := game.Play(1, 1, 1); err != nil {
        fmt.Fprintln(w, err)
    } else {
        fmt.Fprintln(w, "Move recorded at position 1x1")
    }

}
