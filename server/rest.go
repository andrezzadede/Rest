package main 

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"io/ioutil"

  )
	type Produto struct {
		Id         uint64 `json:"id"`
		Descricao  string `json:"descricao"`
		Quantidade int `json:"quantidade"`
		Valor      float64 `json:"valor"`
	}

func main(){
read() // preenchendo o vetor de produtos	

fmt.Println("Conectando o servidor...")

r := mux.NewRouter().StrictSlash(true)

//retornando lista de produtos por GET

r.HandleFunc("/produtos", getProdutos).Methods("GET")


//Retornando produto pelo ID

r.HandleFunc("/produtos/{id}", getProduto).Methods("GET")

//Criando um novo produto usando POST
r.HandleFunc("/produtos", postProduto).Methods("POST")

// Atualizando o produto pelo Id

r.HandleFunc("/produtos/{id}", putProduto).Methods("PUT")

// Deletando o produto pelo id

r.HandleFunc("/produtos/{id}", deleteProduto).Methods("DELETE")


log.Fatal(http.ListenAndServe(":3000", r))
}

var produtos []Produto
// funçao que retorna a lista de produtos para o navegador
func getProdutos (w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(produtos)
}



// função que retorna o produto pelo ID para o navegador

func getProduto(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	param := mux.Vars(r) // capturando a requisiçao

	//converter string para vint64

	var base = 10 // base decimal
	var tam = 32 // tamanho em bits

	IdProd, _ := strconv.ParseUint(param["id"], base, tam)

	// Percorrendo o vetor de produtos 

	for _, produto := range produtos{

		if produto.Id == IdProd{
			json.NewEncoder(w).Encode(produto)
			return 
		}

	}

	json.NewEncoder(w).Encode(&Produto{})
}

// Função pra criar um novo produto "post"

func postProduto (w http.ResponseWriter, r *http.Request){
	var produto Produto

	body, _:= ioutil.ReadAll(r.Body) // lê todos os requisitos que o usuario 

	json.Unmarshal(body, &produto) // adicionando o novo elemento a estrutura

	produtos = append(produtos, produto)

	insert(produto)

}

// função para atualizar o produto 

func putProduto(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type", "application/json")

	param := mux.Vars(r)

	var base = 10
	var tam = 32


	IdProd, _ := strconv.ParseUint(param["id"], base, tam)

	for index, produto := range produtos {

		if produto.Id == IdProd {

			// o produto encontra é substituido pelo novo produto

			produtos = append(produtos[:index], produtos[index + 1:]...)

			var novoProduto Produto

			_= json.NewDecoder(r.Body).Decode(&novoProduto)
			produtos = append(produtos, novoProduto)
			json.NewEncoder(w).Encode(novoProduto)


			update(novoProduto)


			return
		}
	}
	json.NewEncoder(w).Encode(produtos)
}

func deleteProduto(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	param := mux.Vars(r)

	var base = 10
	var tam = 32


	IdProd, _ := strconv.ParseUint(param["id"], base, tam)

	for index, produto := range produtos {

		if produto.Id == IdProd {

			// o produto encontra é substituido pelo novo produto

			produtos = append(produtos[:index], produtos[index + 1:]...)

			delete(IdProd)

			return
		}
	}
	json.NewEncoder(w).Encode(produtos)
}




/* Função que retorna os produtos para o banco */
func read(){
	fmt.Println("Olá bem vindo ao meu banco! Vamos as compras")

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/crud_rest")

	sql := "SELECT * from tb_produto"

	rs, err := db.Query(sql)

	if err != nil {
		panic(err.Error())
	}

	for rs.Next(){
		var prod Produto 

		err = rs.Scan(&prod.Id, &prod.Descricao, &prod.Quantidade, &prod.Valor)
	


		if err != nil {
			panic(err.Error())
		}

		produtos = append(produtos, prod)
		
	}

	fmt.Println(produtos)
}

// funcao para inserir um produto

func insert (prod Produto){
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/crud_rest")

	sql := "INSERT INTO tb_produto (id, descricao, quantidade, valor) VALUES (?, ?, ?, ?)"


	stmt, err := db.Prepare(sql)

	if err != nil {
		panic(err.Error())
	}

	_, er := stmt.Exec(prod.Id, prod.Descricao, prod.Quantidade, prod.Valor)

	if er !=nil{
		panic(er.Error())
	}
	//defer rs.close()

	fmt.Println("Insert OK")
}

func update (prod Produto){
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/crud_rest")

	var sql string = "UPDATE tb_produto SET descricao = ?, quantidade= ?, valor= ? WHERE id= ?"


	stmt, err := db.Prepare(sql)

	if err != nil {
		panic(err.Error())
	}

	_, er := stmt.Exec(prod.Descricao, prod.Quantidade, prod.Valor, prod.Id)

	if er !=nil{
		panic(er.Error())
	}

	
	fmt.Println("update okkk")
}

// DELETANDO PRODUTO pelo id SEMPRE

func delete(id uint64){ // uint64 nao aceita numeros negativos
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/crud_rest")

	var sql string = "DELETE FROM tb_produto WHERE id= ?"


	stmt, err := db.Prepare(sql)

	if err != nil {
		panic(err.Error())
	}

	_, er := stmt.Exec(id)

	if er !=nil{
		panic(er.Error())
	}

	
	fmt.Println("DELETOU")
}