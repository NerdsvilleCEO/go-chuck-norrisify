package controllers

import (
	"fmt"
	"net/http"
	"github.com/NerdsvilleCEO/go-chuck-norrisify/utils"
)

var AppState *App

type App struct {
	ServeMux *http.ServeMux
}

func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func InitializeServer() *App {
	AppState = &App{}
	AppState.ServeMux = http.NewServeMux()
	// TODO: Once caching implemented, up this a bit
	// Let's not wake up the engineers of these API's
	AppState.ServeMux.HandleFunc("/", utils.LimitNumClients(AppState.GetJokeHandler, 20))
	return AppState
}
