package main

import (
	"github.com/NerdsvilleCEO/go-chuck-norrisify/controllers"
	"github.com/NerdsvilleCEO/go-chuck-norrisify/utils"

	"log"
	"net/http"
)

func main() {
	c := utils.LoadConfig()
	app := controllers.InitializeServer()
	log.Fatal(http.ListenAndServe(":"+c.Port, app.ServeMux))
}
