package servidor

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type usuario struct {
	ID    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

//Criar usuário, Insere um usuário no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := ioutil.ReadAll(r.Body)

	tratarErrosNil(erro, "", w)

	var usuario usuario

	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		w.Write([]byte("Erro ao converter o usuário da requisição"))
		return
	}

	db, erro := banco.Conectar()
	tratarErrosNil(erro, "Erro ao conectar com o banco de dados", w)

	statment, erro := db.Prepare("insert into usuarios (nome, email) values (?, ?)")
	tratarErrosNil(erro, "Erro ao inserir no banco de dados", w)

	defer statment.Close()

	insersao, erro := statment.Exec(usuario.Nome, usuario.Email)
	tratarErrosNil(erro, "Erro ao executar o statement!", w)

	idInserido, erro := insersao.LastInsertId()
	tratarErrosNil(erro, "Erro ao obter o id inserido!", w)

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso! Id: %d", idInserido)))

}

// Busca todos usuários banco de dados
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {

	db, erro := banco.Conectar()
	tratarErrosNil(erro, "Erro ao obter conexão com o banco de dados", w)

	defer db.Close()

	linhas, erro := db.Query("select * from usuarios")
	tratarErrosNil(erro, "Erro ao buscar os usuários", w)

	defer linhas.Close()

	var usuarios []usuario

	for linhas.Next() {
		var usuario usuario

		if erro := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("Erro ao escanear o usuário"))
			return
		}
		usuarios = append(usuarios, usuario)
	}
	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(usuarios); erro != nil {
		w.Write([]byte("Erro ao converter usuarios para JSON"))
		return
	}
}

// Busca um usuário especifico
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	tratarErrosNil(erro, "Erro ao converter o parâmetro para inteiro", w)

	db, erro := banco.Conectar()
	tratarErrosNil(erro, "Erro ao obter conexão com o banco de dados", w)

	linha, erro := db.Query("select * from usuarios where id = ?", ID)
	tratarErrosNil(erro, "Erro ao buscar o usuario", w)

	var usuario usuario
	if linha.Next() {
		if erro := linha.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("Erro ao escanear o usuário"))
			return
		}
	}

	if erro := json.NewEncoder(w).Encode(usuario); erro != nil {
		w.Write([]byte("Erro ao converter o usuário para JSON"))
		return
	}
}

// altera os dados de um usuario no banco de dados
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {

	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	tratarErrosNil(erro, "Erro ao converter o parâmetro para inteiro", w)

	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	tratarErrosNil(erro, "Erro ao ler o corpo de requisicao", w)

	var usuario usuario

	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		w.Write([]byte("Erro ao converter usuário para struct"))
		return
	}

	db, erro := banco.Conectar()
	tratarErrosNil(erro, "Erro ao obter conexão com o banco de dados", w)

	defer db.Close()

	statement, erro := db.Prepare("update usuarios set nome = ?, email = ? where id = ?")
	tratarErrosNil(erro, "Erro ao criar o statement!", w)

	defer statement.Close()

	if _, erro := statement.Exec(usuario.Nome, usuario.Email, ID); erro != nil {
		w.Write([]byte("Erro ao atualizar o usuário!"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// DeletarUsuario remove um usuário do banco de dados
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter o parâmetro para inteiro"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco de dados!"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("delete from usuarios where id = ?")
	if erro != nil {
		w.Write([]byte("Erro ao criar o statement!"))
		return
	}
	defer statement.Close()

	if _, erro := statement.Exec(ID); erro != nil {
		w.Write([]byte("Erro ao deletar o usuário!"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func tratarErrosNil(erro error, msgErro string, w http.ResponseWriter) {
	if erro != nil {
		w.Write([]byte(msgErro))
		return
	}
}
