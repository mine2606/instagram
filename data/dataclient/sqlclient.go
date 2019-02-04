package dataclient

import (
	"database/sql"
	"fmt"
	"instagram/data/model"

	_ "github.com/go-sql-driver/mysql" ///El driver se registra en database/sql en su función Init(). Es usado internamente por éste
)

//InsertarUser test
func InsertarUser(objeto *model.User) {

	db, err := sql.Open("mysql", "ubuntu:ubuntu@tcp(localhost:3306)/Instagram")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("nombre: ", objeto.Nombre)
	defer db.Close()
	insert, err := db.Query("INSERT INTO User(Nombre, Usuario, Password, Email) VALUES (?, ?, ?, ?)", objeto.Nombre, objeto.Usuario, objeto.Password, objeto.Email)

	if err != nil {
		panic(err.Error())
	}
	insert.Close()
}

//Login función para iniciar sesion
func Login(objeto *model.Login) string {
	db, err := sql.Open("mysql", "ubuntu:ubuntu@tcp(localhost:3306)/Instagram")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	comando := "SELECT Password FROM User WHERE (Usuario = '" + objeto.Usuario + "')"
	fmt.Println(comando)
	query, err := db.Query("SELECT Password FROM User WHERE (Usuario = '" + objeto.Usuario + "')")
	if err != nil {
		panic(err.Error())
	}
	defer query.Close()

	var resultado string

	for query.Next() {
		err = query.Scan(&resultado)
		if err != nil {
			panic(err.Error())
		}
	}
	return resultado
}
