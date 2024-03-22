package model

import (
	exeption "loja-virtual/internal/exception"
	"loja-virtual/internal/db"
)

// Produto is a struct that represents a product
type Produto struct {
	ID    int `json:"id"`
	Nome  string `json:"nome"`
	Preco float64 `json:"preco"`
	Descricao string `json:"descricao"`
	Quantidade int `json:"quantidade"`
}

func BuscarTodosProdutos() []Produto {
	 db := config.Connect()

	rows, err := db.Query("SELECT * FROM produto order by id asc")
	exeption.Error(err)

	produto := Produto{}
	produtos := []Produto{}

	for rows.Next() {
		err = rows.Scan(&produto.ID, &produto.Nome, &produto.Descricao, &produto.Preco, &produto.Quantidade)
		exeption.Error(err)
		produtos = append(produtos, produto)
		
	}

	defer db.Close()
	return produtos
}

func (p *Produto) CriarNovoProduto() {
	db := config.Connect()

	insereDadosNoBanco, err := db.Prepare("INSERT INTO produto (nome, descricao, preco, quantidade) VALUES ($1, $2, $3, $4)")
	exeption.Error(err)

	_, err = insereDadosNoBanco.Exec(p.Nome, p.Descricao, p.Preco, p.Quantidade)
	exeption.Error(err)

	defer db.Close()
}

func DeletarProduto(idProduto int) {
	db := config.Connect()

	deletarProduto, err := db.Prepare("DELETE FROM produto WHERE id = $1")
	exeption.Error(err)

	deletarProduto.Exec(idProduto)

	defer db.Close()
}

func BuscarProduto(idProduto int) Produto {
	db := config.Connect()

	linha, err := db.Query("SELECT * FROM produto WHERE id = $1", idProduto)
	exeption.Error(err)

	produto := Produto{}

	if linha.Next() {
		err = linha.Scan(&produto.ID, &produto.Nome, &produto.Descricao, &produto.Preco, &produto.Quantidade)
		exeption.Error(err)
	}

	defer db.Close()
	return produto
}

func (p *Produto) AtualizarProduto() {
	db := config.Connect()

	atualizarProduto, err := db.Prepare("UPDATE produto SET nome = $1, descricao = $2, preco = $3, quantidade = $4 WHERE id = $5")
	exeption.Error(err)

	atualizarProduto.Exec(p.Nome, p.Descricao, p.Preco, p.Quantidade, p.ID)

	defer db.Close()
}