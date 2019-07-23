# Go Noughts & Crosses

A simple Noughts & Crosses written in go.
This is a server game intented to be played via HTTP calls.

## Installation

* Download package locally
* Install 3rd- party libraries `go get`
* Run package using `go run .`

## Playing

Browse provided URL to play game

### GET `http://localhost:8080/`

Displays status and board

### POST `http://localhost:8080/`

Plays a move for the current player. Valid `application/json` is expected at the current format:
```
{
    "col": x
    "row": y
}
```
Where `x` is the column and `y` is the row.

### GET `http://localhost:8080/reset/`

Resets game


## Design notes

* This use `mux` router to direct requests to appropriate handlers.
* Basic validatation is done when sending data via POST.
* Data is retained in memory only, game is reset on each launch.
* Game does not stop when won, but it displays a message.
* Tests are not provided due to time constraint but some functions could be good candidates for unit-tests, in particular model methods `GetWinner` and `Play`.



