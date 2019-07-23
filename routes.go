package main

import (
    "net/http"
    "github.com/gorilla/mux"
)

// Route is a structure holding a route for the mux router
type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}


// NewRouter returns the route
func NewRouter() *mux.Router {
    router := mux.NewRouter().StrictSlash(true)
    for _, route := range routes {
        handler := Logger(route.HandlerFunc, route.Name)
        router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(handler)
    }
    return router
}

// Routes holds the list of routes for the mux router
type Routes []Route

var routes = Routes{
    Route{
        "Show status",
        "GET",
        "/",
        Status,
    },
    Route{
        "Reset game",
        "GET",
        "/reset",
        Reset,
    },
    Route{
        "Play a move",
        "POST",
        "/",
        Play,
    },
}
