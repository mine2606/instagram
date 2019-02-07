package dataclient

import (
	"database/sql"
	"fmt"
	"instagram/data/model"

	_ "github.com/go-sql-driver/mysql" ///El driver se registra en database/sql en su función Init(). Es usado internamente por éste
)

//InsertarUser test
func InsertarUser(objeto *model.User) bool {

	db, err := sql.Open("mysql", "ubuntu:ubuntu@tcp(localhost:3306)/Instagram")
	if err != nil {
		panic(err.Error())
	}

	resp := false

	//Comprobamos si el usuario y el email existen
	comando := "SELECT ID FROM User WHERE (Usuario = '" + objeto.Usuario + "' OR Email = '" + objeto.Email + "') LIMIT 1"
	fmt.Println(comando)
	query, err := db.Query("SELECT ID FROM User WHERE (Usuario = ? OR Email = ?) LIMIT 1", objeto.Usuario, objeto.Email)

	var resultado string

	//Si existe no muestra la respuesta
	for query.Next() {
		err = query.Scan(&resultado)
		if err != nil {
			panic(err.Error())
		}
	}

	//Si está vacío y los datos no coinciden, los guarda en la base de datos y nos muestra la respuesta
	if resultado == "" {
		fmt.Println("nombre: ", objeto.Nombre)
		defer db.Close()
		insert, err := db.Query("INSERT INTO User(Nombre, Usuario, Password, Email) VALUES (?, ?, ?, ?)", objeto.Nombre, objeto.Usuario, objeto.Password, objeto.Email)

		if err != nil {
			panic(err.Error())
		}
		insert.Close()
		resp = true
	}

	return resp

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

//SubirFoto test
func SubirFoto(url string, texto string, id int) {
	db, err := sql.Open("mysql", "ubuntu:ubuntu@tcp(localhost:3306)/Instagram")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	insert, err := db.Query("INSERT INTO Foto(URL, Texto, User_ID) VALUES (?, ?, ?)", url, texto, id)
	if err != nil {
		panic(err.Error())
	}
	insert.Close()
}

//ConsultaID test
func ConsultaID(usuario string) int {
	db, err := sql.Open("mysql", "ubuntu:ubuntu@tcp(localhost:3306)/Instagram")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	comando := "SELECT ID FROM User WHERE (Usuario = '" + usuario + "')"
	fmt.Println(comando)
	query, err := db.Query("SELECT ID FROM User WHERE (Usuario = '" + usuario + "')")
	if err != nil {
		panic(err.Error())
	}
	defer query.Close()
	var resultado int
	for query.Next() {
		err := query.Scan(&resultado)
		if err != nil {
			panic(err.Error())
		}
	}
	return resultado
}

//MostrarFoto test
func MostrarFoto() []model.RFoto {
	db, err := sql.Open("mysql", "ubuntu:ubuntu@tcp(localhost:3306)/Instagram")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	comando := "SELECT ID, Texto, Url FROM Foto"
	fmt.Println(comando)

	query, err := db.Query("SELECT ID, Texto, Url FROM Foto")

	if err != nil {
		panic(err.Error())
	}

	resultado := make([]model.RFoto, 0)

	for query.Next() {
		var foto = model.RFoto{}
		err = query.Scan(&foto.ID, &foto.Texto, &foto.URL)
		if err != nil {
			panic(err.Error())
		}
		resultado = append(resultado, foto)
	}
	return resultado
}
