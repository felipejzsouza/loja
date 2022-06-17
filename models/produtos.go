package modelos

import "github.com/felipejzsouza/loja/db"

type Produto struct {
	id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaProdutos() []Produto {
	db := db.ConectaComBancoDeDados()

	selectDeTodosOsProdutos, err := db.Query("select * from produtos")
	if err != nil {
		panic(err.Error())
	}

	produto := Produto{}

	produtos := []Produto{}

	for selectDeTodosOsProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectDeTodosOsProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		produto.id = id
		produto.Nome = nome
		produto.Descricao = descricao
		produto.Preco = preco
		produto.Quantidade = quantidade

		produtos = append(produtos, produto)
	}

	defer db.Close()
	return produtos
}

func CriarNovoProduto(nome, descricao string, preco float64, quantidade int) {
	dba := db.ConectaComBancoDeDados()

	insereDadosNoBanco, err := dba.Prepare("insert into produtos (nome, descricao, preco, quantidade) values ($1, $2, $3, $4);")
	if err != nil {
		panic(err.Error())
	}

	insereDadosNoBanco.Exec(nome, descricao, preco, quantidade)

	defer dba.Close()
}
