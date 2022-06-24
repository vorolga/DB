package main

import (
	app "TP_DB/internal/app"
	"log"
	"os"

	"github.com/valyala/fasthttp"
)

func main() {
	application, err := app.NewApp()
	if err != nil {
		os.Exit(1)
	}

	router := application.CreateRouter()
	log.Println("server running at 5001")
	panic(fasthttp.ListenAndServe(":5001", router.HandleRequest))
}
