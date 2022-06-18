package modelos

import "github.com/felipejzsouza/loja/db"

type Produto struct {
	Id         int
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

		produto.Id = id
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
	db := db.ConectaComBancoDeDados()

	insereDadosNoBanco, err := db.Prepare("insert into produtos (nome, descricao, preco, quantidade) values ($1, $2, $3, $4);")
	if err != nil {
		panic(err.Error())
	}
	insereDadosNoBanco.Exec(nome, descricao, preco, quantidade)
	defer db.Close()
}

func DeletarProduto(id string) {
	db := db.ConectaComBancoDeDados()
	deletarProduto, err := db.Prepare("delete from produtos where id = $1;")
	if err != nil {
		panic(err.Error())
	}
	deletarProduto.Exec(id)
	defer db.Close()
}

func BuscarProduto(idProduto string) Produto {
	db := db.ConectaComBancoDeDados()
	selecionarProduto, err := db.Query("select * from produtos where id = $1", idProduto)
	if err != nil {
		panic(err.Error())
	}

	var id, quantidade int
	var nome, descricao string
	var preco float64
	var produto = Produto{}

	for selecionarProduto.Next() {
		err = selecionarProduto.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		produto.Id = id
		produto.Nome = nome
		produto.Descricao = descricao
		produto.Preco = preco
		produto.Quantidade = quantidade
	}

	defer db.Close()

	return produto
}

func AtualizarProduto(id int, nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaComBancoDeDados()

	atualizaProduto, err := db.Prepare("update produtos set nome = $1, descricao = $2, preco = $3,  quantidade = $4 where id = $5;")
	if err != nil {
		panic(err.Error())
	}
	atualizaProduto.Exec(nome, descricao, preco, quantidade, id)
	defer db.Close()
}
