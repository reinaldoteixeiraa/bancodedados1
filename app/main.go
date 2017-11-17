package main

import (
	"log"
	"net/http"
	"html/template"
//	_ "github.com/go-sql-driver/mysql"
//	sql "database/sql"
	_"fmt"
<<<<<<< HEAD
<<<<<<< HEAD
	"debug/dwarf"
	"time"
=======
	_"io/ioutil"
	
>>>>>>> 8e46668d2b5b852bdc83c61468a017e23175852d
=======
>>>>>>> 6b1c07ab90c76d6fad1b442721db217e9f08b1d6
)
var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))//erro aqui
}

func main() {
	//Rotas
	http.HandleFunc("/", index)
	http.HandleFunc("/aluno",aluno)
	http.HandleFunc("/professor",professor)
	http.HandleFunc("/coordenacao",coordenacao)
	http.HandleFunc("/coordenacao/add_aluno",addAluno)
	http.HandleFunc("/coordenacao/add_professor",addProfessor)
	http.HandleFunc("/coordenacao/add_coordenacao",addCoordenador)
	http.HandleFunc("/coordenacao/add_disciplinas",addDisciplinas)
	http.HandleFunc("/coordenacao/add_atuser",addAtuser)
	http.HandleFunc("/coordenacao/addvisualizar",addVizualizar)
	http.HandleFunc("/coordenacao/add_monitores",addMonitores)
	http.HandleFunc("/coordenacao/add_colegios",addColegios)
	//Host
	http.ListenAndServe(":9500",nil)
}

type usuario struct{
	cpf int   	  'json: "cpf"'
	rg  int	   	  'json: "rg"'
	nome string	  'json: "nome"'
	email string	  'json: "email"'
	telefone
	end
}

type end struct {
	logradouro  string	'json: "logradouro"'
	numero 	    int		'json: "numero"'
	bairro      string	'json: "bairro"'
	cidade	    string	'json: "cidade"'
	cep 	    int		'json: "cep"'
}

type telefone struct {
	usuario		'json: "usuario"'
	numero 		int 'json: "numero"'
}

type aluno struct {
	matricula  int		'json: "matricula"'
	usuario			'json: "usuario"'
	nasc 	   time.Time	'json: "nasc"'
	le 	   bool		'json: "le"'
	responsavel  string	'json: "responsavel"'
}

type disciplinas struct {
	professor		'json: "professor"'
	nome 	string		'json: "nome"'
}

type professor struct {
	idprofessor 	int	'json: "idprofessor"'
	usuario			'json: "usuario"'
	disciplina		'json: "disciplina"'
}

type coordenador struct {
	idcoordenador 	int  	'json: "idcoordenador"'
	usuario			'json: "usuario"'
}

type monitores struct{
	cpf int   	  'json: "cpf"'
	rg  int	   	  'json: "rg"'
	nome string	  'json: "nome"'
	email string	  'json: "email"'
	telefone
	end
}
type colegios struct {
	professor		'json: "professor"'
	nome 	string		'json: "nome"'
}

func index(w http.ResponseWriter, req *http.Request){
	err := tpl.ExecuteTemplate(w, "index.gohtml",nil)
	HandleError(w,err)
}
func aluno(w http.ResponseWriter, req *http.Request){
	err := tpl.ExecuteTemplate(w, "aluno.gohtml",nil)
	HandleError(w,err)
}
func professor(w http.ResponseWriter, req *http.Request){
	err := tpl.ExecuteTemplate(w, "professor.gohtml",nil)
	HandleError(w,err)
}
func coordenacao(w http.ResponseWriter, req *http.Request){
	err := tpl.ExecuteTemplate(w,"coordenacao.gohtml",nil)
	HandleError(w,err)
}
func addAluno(w http.ResponseWriter, req *http.Request){
	err := tpl.ExecuteTemplate(w, "add_aluno.gohtml",nil)
	HandleError(w,err)
}
func addProfessor(w http.ResponseWriter, req *http.Request){
	err := tpl.ExecuteTemplate(w, "add_professor.gohtml",nil)
	HandleError(w,err)
}
func addCoordenador(w http.ResponseWriter, req *http.Request){
	err := tpl.ExecuteTemplate(w, "add_coordenador.gohtml",nil)
	HandleError(w,err)
}
func addDisciplinas(w http.ResponseWriter, req *http.Request){
	err := tpl.ExecuteTemplate(w, "add_disciplinas.gohtml",nil)
	HandleError(w,err)
}
func addAtuser(w http.ResponseWriter, req *http.Request){
	err := tpl.ExecuteTemplate(w, "add_atuser.gohtml",nil)
	HandleError(w,err)
}
func addVizualizar(w http.ResponseWriter, req *http.Request){
	err := tpl.ExecuteTemplate(w, "visualizar.gohtml",nil)
	HandleError(w,err)
}
func addMonitores(w http.ResponseWriter, req *http.Request){
	err := tpl.ExecuteTemplate(w, "add_monitores.gohtml",nil)
	HandleError(w,err)
}
func addColegios(w http.ResponseWriter, req *http.Request){
	err := tpl.ExecuteTemplate(w, "add_colegios.gohtml",nil)
	HandleError(w,err)
}

func HandleError(w http.ResponseWriter,err error){
	if err!= nil{
		http.Error(w,err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
