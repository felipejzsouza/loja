package controllers

import (
	"log"
	"net/http"
	"strconv"
	"text/template"

	modelos "github.com/felipejzsouza/loja/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := modelos.BuscaProdutos()
	temp.ExecuteTemplate(w, "Index", produtos)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertido, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro ao converter preco:", preco)
		}

		quantidadeConvertida, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro ao converter qunatidade:", quantidade)
		}
		modelos.CriarNovoProduto(nome, descricao, precoConvertido, quantidadeConvertida)
	}

	http.Redirect(w, r, "/", 301)
}
