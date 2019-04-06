package main

import (
	"log"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"github.com/rof20004/vuttr/handlers"
)

func main() {
	router := fasthttprouter.New()
	router.ServeFiles("/doc/*filepath", "./public")

	handlers.InitializeToolsHandlers(router)

	log.Println("Servidor iniciado na porta 3000")
	log.Fatal(fasthttp.ListenAndServe(":3000", router.Handler))
}
