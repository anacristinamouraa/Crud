package main

import (
	"database/sql"  // Pacote Database SQL para realizar Query
	"log"           // Mostra mensagens no console
	"net/http"      // Gerencia URLs e Servidor Web
	"text/template" // Gerencia templates

	_ "github.com/go-sql-driver/mysql" // Driver Mysql para Go
)

//Struct utilizada para exibir dados no template
type Livors struct {
	Id    int
	titulo  string
	categoria string
	autor string
	sinopse string

}

// Função dbConn, abre a conexão com o banco de dados
func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "crudgo"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

// A variável tmpl renderiza todos os templates da pasta 'tmpl' independente da extensão
var tmpl = template.Must(template.ParseGlob("tmpl/*"))

// Função usada para renderizar o arquivo Index
func Index(w http.ResponseWriter, r *http.Request) {
	// Abre a conexão com o banco de dados utilizando a função dbConn()
	db := dbConn()
	// Realiza a consulta com banco de dados e trata erros
	selDB, err := db.Query("SELECT * FROM livros ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}

	// Monta a struct para ser utilizada no template
	n := Names{}

	// Monta um array para guardar os valores da struct
	res := []Names{}

	// Realiza a estrutura de repetição pegando todos os valores do banco
	for selDB.Next() {
		// Armazena os valores em variáveis
		var id int
		var name, email string

		// Faz o Scan do SELECT
		err = selDB.Scan(&id, &titulo, &categoria, &autor, &sinopse)
		if err != nil {
			panic(err.Error())
		}

		// Envia os resultados para a struct
		n.Id = id
		n.Name = name
		n.Email = email

		// Junta a Struct com Array
		res = append(res, n)
	}

	// Abre a página Index e exibe todos os registrados na tela
	tmpl.ExecuteTemplate(w, "Index", res)

	// Fecha a conexão
	defer db.Close()
}

// Função Show exibe apenas um resultado
func Show(w http.ResponseWriter, r *http.Request) {
	db := dbConn()

	// Pega o ID do parametro da URL
	nId := r.URL.Query().Get("id")

	// Usa o ID para fazer a consulta e tratar erros
	selDB, err := db.Query("SELECT * FROM livros WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}

	// Monta a strcut para ser utilizada no template
	l := Livros{}

	// Realiza a estrutura de repetição pegando todos os valores do banco
	for selDB.Next() {
		// Armazena os valores em variaveis
		var id int
		var titulo, categoria, autor, sinopse string

		// Faz o Scan do SELECT
		err = selDB.Scan(&id, &titulo, &categoria, &autor, &sinopse)
		if err != nil {
			panic(err.Error())
		}

		// Envia os resultados para a struct
		n.Id = id
		n.Titulo = titulo
		n.Categoria = categoria
		n.Autor = autor 
		n.Sinopse = sinopse

	}

	// Mostra o template
	tmpl.ExecuteTemplate(w, "Show", n)

	// Fecha a conexão
	defer db.Close()

}

// Função New apenas exibe o formulário para inserir novos dados
func New(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}

// Função Edit, edita os dados
func Edit(w http.ResponseWriter, r *http.Request) {
	// Abre a conexão com banco de dados
	db := dbConn()

	// Pega o ID do parametro da URL
	nId := r.URL.Query().Get("id")

	selDB, err := db.Query("SELECT * FROM livros WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}

	// Monta a struct para ser utilizada no template
	l := Livros{}

	// Realiza a estrutura de repetição pegando todos os valores do banco
	for selDB.Next() {
		//Armazena os valores em variaveis
		var id int
		var name, email string

		// Faz o Scan do SELECT
		err = selDB.Scan(&id, &name, &email)
		if err != nil {
			panic(err.Error())
		}

		// Envia os resultados para a struct
		n.Id = id
		n.titulo = titulo
		n.categoria = categoria
		n.autor = autor
		n.sinopse = sinopse
	}

	// Mostra o template com formulário preenchido para edição
	tmpl.ExecuteTemplate(w, "Edit", n)

	// Fecha a conexão com o banco de dados
	defer db.Close()
}

// Função Insert, insere valores no banco de dados
func Insert(w http.ResponseWriter, r *http.Request) {

	//Abre a conexão com banco de dados usando a função: dbConn()
	db := dbConn()

	// Verifica o METHOD do fomrulário passado
	if r.Method == "POST" {

		// Pega os campos do formulário
		titulo := r.FormValue("titulo")
		categoria := r.FormValue("categoria")
		autor := r.FormValue("autor")
		sinopse := r.FormValue("sinopse")

		// Prepara a SQL e verifica errors
		insForm, err := db.Prepare("INSERT INTO livros(titulo, categoria, autor, sinopse) VALUES(?,?)")
		if err != nil {
			panic(err.Error())
		}

		// Insere valores do formulario com a SQL tratada e verifica errors
		insForm.Exec(titulo, categoria, autor, sinopse)

		// Exibe um log com os valores digitados no formulário
		log.Println("INSERT: Titulo: " + Titulo + " | Categoria: " + categoria |  Autor: " + autor | Sinopse: " + sinopse)
	}

	// Encerra a conexão do dbConn()
	defer db.Close()

	//Retorna a HOME
	http.Redirect(w, r, "/", 301)
}

// Função Update, atualiza valores no banco de dados
func Update(w http.ResponseWriter, r *http.Request) {

	// Abre a conexão com o banco de dados usando a função: dbConn()
	db := dbConn()

	// Verifica o METHOD do formulário passado
	if r.Method == "POST" {

		// Pega os campos do formulário
		titulo := r.FormValue("titulo")
		categoria := r.FormValue("categoria")
		autor := r.FormValue("autor")
		sinopse := r.FormValue("sinopse")
		id := r.FormValue("uid")

		// Prepara a SQL e verifica errors
		insForm, err := db.Prepare("UPDATE names SET name=?, email=? WHERE id=?")
		if err != nil {
			panic(err.Error())
		}

		// Insere valores do formulário com a SQL tratada e verifica erros
		insForm.Exec(titulo, categoria, autor, sinopse)

		// Exibe um log com os valores digitados no formulario
		log.Println("UPDATE: : "Titulo: " + Titulo + " | Autor: " + autor |  Categoria: " + categoria | Sinopse: " + sinopse)
	}

	// Encerra a conexão do dbConn()
	defer db.Close()

	// Retorna a HOME
	http.Redirect(w, r, "/", 301)
}

// Função Delete, deleta valores no banco de dados
func Delete(w http.ResponseWriter, r *http.Request) {

	// Abre conexão com banco de dados usando a função: dbConn()
	db := dbConn()

	nId := r.URL.Query().Get("id")

	// Prepara a SQL e verifica errors
	delForm, err := db.Prepare("DELETE FROM livros WHERE id=?")
	if err != nil {
		panic(err.Error())
	}

	// Insere valores do form com a SQL tratada e verifica errors
	delForm.Exec(nId)

	// Exibe um log com os valores digitados no form
	log.Println("DELETE")

	// Encerra a conexão do dbConn()
	defer db.Close()

	// Retorna a HOME
	http.Redirect(w, r, "/", 301)
}

func main() {

	// Exibe mensagem que o servidor foi iniciado
	log.Println("Server started on: http://localhost:9000")

	// Gerencia as URLs
	http.HandleFunc("/", Index)
	http.HandleFunc("/show", Show)
	http.HandleFunc("/new", New)
	http.HandleFunc("/edit", Edit)

	// Ações
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/delete", Delete)

	// Inicia o servidor na porta 9000
	http.ListenAndServe(":9000", nil)
}
