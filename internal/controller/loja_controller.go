package controller

import (
	"html/template"
	"loja-virtual/internal/exception"
	"loja-virtual/internal/model"
	"net/http"
	"strconv"
)

var temp = template.Must(template.ParseGlob("../../internal/templates/*.html"))


func Index(w http.ResponseWriter, r *http.Request)  {
	produtos := model.BuscarTodosProdutos()

	temp.ExecuteTemplate(w, "Index", produtos)
}

func New(w http.ResponseWriter, r *http.Request)  {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request)  {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertido, err := strconv.ParseFloat(preco, 64)
		exception.Error(err)

		quantidadeConvertida, err := strconv.Atoi(quantidade)
		exception.Error(err)

		newProduto := model.Produto{
			Nome: nome,
			Descricao: descricao,
			Preco: precoConvertido,
			Quantidade: quantidadeConvertida}

		newProduto.CriarNovoProduto()
	}

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Delete(w http.ResponseWriter, r *http.Request)  {
	idProduto := r.URL.Query().Get("id")
	idProdutoConvertido, err := strconv.Atoi(idProduto)
	exception.Error(err)

	model.DeletarProduto(idProdutoConvertido)

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Edit(w http.ResponseWriter, r *http.Request)  {
	idProduto := r.URL.Query().Get("id")
	idProdutoConvertido, err := strconv.Atoi(idProduto)
	exception.Error(err)

	produto := model.BuscarProduto(idProdutoConvertido)

	temp.ExecuteTemplate(w, "Edit", produto)
}

func Update(w http.ResponseWriter, r *http.Request)  {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		idConvertido, err := strconv.Atoi(id)
		exception.Error(err)

		precoConvertido, err := strconv.ParseFloat(preco, 64)
		exception.Error(err)

		quantidadeConvertida, err := strconv.Atoi(quantidade)
		exception.Error(err)

		produto := model.Produto{
			ID: idConvertido,
			Nome: nome,
			Descricao: descricao,
			Preco: precoConvertido,
			Quantidade: quantidadeConvertida}

		produto.AtualizarProduto()
	}

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
