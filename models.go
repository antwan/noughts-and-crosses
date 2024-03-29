package main

import (
    "bytes"
    "fmt"
)

// BOARDROWS is the number of rows in the board
const BOARDROWS = 3

// BOARDCOLS is the number of rows in the board
const BOARDCOLS = 3

// Noughts is a structure holding a game of noughts and crosses
type Noughts struct {
    board [BOARDROWS][BOARDCOLS]int
    currentPlayer int
}

// Play records a move in a game of nought
func (n *Noughts) Play(row int, col int, player int) error {
    if row < 0 || BOARDROWS <= row || col < 0 || BOARDCOLS <= col {
        return fmt.Errorf("Invalid coordinate")
    }
    if player != 1 && player != 2 {
        return fmt.Errorf("Invalid player")
    }
    if n.board[row][col] != 0 {
        return fmt.Errorf("Cannot play this move, it's already filled")
    }
    n.board[row][col] = player
    return nil
}

// GetWinner returns the current game winner if any
func (n *Noughts) GetWinner() int {
    // Checks if winner exists on rows
    for r := 0; r < BOARDROWS; r++ {
        rowWinner := 0
        for c := 0; c < BOARDCOLS; c++ {
            if c == 0 && n.board[r][c] != 0 {
                rowWinner = n.board[r][c]
            } else if n.board[r][c] != rowWinner {
                rowWinner = 0
                break
            }
        }
        if rowWinner != 0 {
            fmt.Println("rowWinner")
            return rowWinner
        }
    }

    // Checks if winner exists on cols
    for c := 0; c < BOARDCOLS; c++ {
        colWinner := 0
        for r := 0; r < BOARDROWS; r++ {
            if r == 0 && n.board[r][c] != 0 {
                colWinner = n.board[r][c]
            } else if n.board[r][c] != colWinner {
                colWinner = 0
                break
            }
        }
        if colWinner != 0 {
            fmt.Println("colWinner")
            return colWinner
        }
    }

    // Checks if winner exists on diags
    if BOARDCOLS == BOARDROWS {
        diagWinner := 0
        for c := 0; c < BOARDCOLS; c++ {
            if c == 0 && n.board[c][c] != 0 {
                fmt.Printf("n.board[%v][%v]=%v", c, c, n.board[c][c])
                diagWinner = n.board[c][c]
            } else if n.board[c][c] != diagWinner {
                fmt.Printf("n.board[%v][%v]=%v != %v", c, c, n.board[c][c], diagWinner)
                diagWinner = 0
                break
            }
        }
        if diagWinner != 0 {
            fmt.Println("diagwinner")
            return diagWinner
        }
        diagWinner = 0
        for c := 0; c < BOARDCOLS; c++ {
            d := BOARDCOLS - 1 - c

            if c == 0 && n.board[d][c] != 0 {
                diagWinner = n.board[d][c]
            } else if n.board[d][c] != diagWinner {
                diagWinner = 0
                break
            }
        }
        if diagWinner != 0 {
            fmt.Println("diagwinner2")
            return diagWinner
        }
    }
    return 0
}

func (n *Noughts) String() string {
    w := new(bytes.Buffer)

    for r := 0; r < BOARDROWS; r++ {
        for c := 0; c < BOARDCOLS; c++ {
            fmt.Fprint(w, "+---")
        }
        fmt.Fprintln(w, "+")
        for c := 0; c < BOARDCOLS; c++ {
            fmt.Fprint(w, "|")
            if n.board[r][c] == 1 {
                fmt.Fprint(w, " O ")
            } else if n.board[r][c] == 2 {
                fmt.Fprint(w, " x ")
            } else {
                fmt.Fprint(w, "   ")
            }
        }
        fmt.Fprintln(w, "|")
    }
    for c := 0; c < BOARDCOLS; c++ {
        fmt.Fprint(w, "+---")
    }
    fmt.Fprintln(w, "+")

    return w.String()
}
