package main

import (
	"net/http"

	"github.com/felipejzsouza/loja/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
