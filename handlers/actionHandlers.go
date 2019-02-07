package handlers

import (
	"encoding/json"
	"fmt"
	client "instagram/data/dataclient"
	"instagram/data/model"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/securecookie"
	"golang.org/x/crypto/bcrypt"
)

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

func getUserName(request *http.Request) (usuario string) {
	if cookie, err := request.Cookie("session"); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil {
			usuario = cookieValue["name"]
		}
	}
	return usuario
}

func setSession(usuario string, response http.ResponseWriter) {
	value := map[string]string{
		"name": usuario,
	}
	if encoded, err := cookieHandler.Encode("session", value); err == nil {
		cookie := &http.Cookie{
			Name:  "session",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(response, cookie)
	}
}

func clearSession(response http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(response, cookie)
}

//Login Función para acceder a la página
func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Incoming request from " + r.URL.EscapedPath())
	if r.URL.Path != PathLogin {
		http.NotFound(w, r)
		return
	}
	if r.Method != http.MethodPost {
		http.NotFound(w, r)
		return
	}

	defer r.Body.Close()
	bytes, e := ioutil.ReadAll(r.Body)

	respuesta := false
	if e == nil {
		// datos que recibe del cliente
		var user model.Login
		enTexto := string(bytes)
		fmt.Println("En texto: " + enTexto)
		_ = json.Unmarshal(bytes, &user)

		fmt.Println(user.Usuario)

		if user.Usuario == "" || user.Password == "" {
			fmt.Fprintln(w, "La petición está vacía")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Contraseña de la base de datos
		password := client.Login(&user)

		// Comprueba que las dos contraseñas sean iguales
		if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(user.Password)); err != nil {
			fmt.Printf("No has podido inicar sesión")
		} else {
			respuesta = true
			setSession(user.Usuario, w)
			fmt.Println("Inicio de sesión realizado")
			getUserName(r)
		}

	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, respuesta)
	}

	fmt.Fprintln(w, respuesta)
}

//Logout Función para cerrar sesion
func Logout(response http.ResponseWriter, request *http.Request) {
	clearSession(response)
	http.Redirect(response, request, "/", 302)
}

//Uploader funcion para subir una foto y guardarla en carpeta
/*func Uploader(w http.ResponseWriter, r *http.Request) {

	r.ParseMultipartForm(2000)

	file, fileInfo, err := r.FormFile("archivo")

	f, err := os.OpenFile("./files/"+fileInfo.Filename, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		http.NotFound(w, r)
		return
	}

	defer f.Close()

	io.Copy(f, file)
	fmt.Fprintf(w, fileInfo.Filename)

}*/

//Uploader Función sube archivos
func Uploader(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Incoming request from " + r.URL.EscapedPath())
	if r.URL.Path != PathUploader {
		http.NotFound(w, r)
		return
	}
	if r.Method != http.MethodPost {
		http.NotFound(w, r)
		return
	}
	r.ParseMultipartForm(2000)
	//Coger el archivo y meterlo en una variable
	file, fileInto, err := r.FormFile("archivo")
	//Coger el texto del formulario y merterlo en una variable
	texto := r.FormValue("texto")

	usuario := getUserName(r)

	fmt.Println(texto, "Nombre Usuario: ", usuario)

	f, err := os.OpenFile("./files/"+fileInto.Filename, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		log.Fatal(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)
	//La linea de abajo que esta comentada me manda a la página donde está el nombre del archivo
	//fmt.Fprintf(w, fileInto.Filename)
	//Esta linea de aqui abajo me manda a la pagina principal donde están todas las fotos
	http.Redirect(w, r, "/", 301)
	//Datos de la base de datos
	id := client.ConsultaID(usuario)
	fmt.Println(id)
	//Subir foto a la base de datos
	go client.SubirFoto(fileInto.Filename, texto, id)
}

//RegistroUsuario Función que inserta un registro en la base de datos local
func RegistroUsuario(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Incoming request from " + r.URL.EscapedPath())
	if r.URL.Path != PathRegistroUsuario {
		http.NotFound(w, r)
		return
	}
	if r.Method != http.MethodPost {
		http.NotFound(w, r)
		return
	}

	defer r.Body.Close()
	bytes, e := ioutil.ReadAll(r.Body)
	if e == nil {
		var user model.User
		enTexto := string(bytes)
		fmt.Println("En texto: " + enTexto)
		_ = json.Unmarshal(bytes, &user)
		fmt.Println(user.Nombre)
		if user.Nombre == "" || user.Usuario == "" || user.Password == "" || user.Email == "" {
			fmt.Fprintln(w, "La petición está vacía")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			fmt.Println(err)
		}
		hashComoCadena := string(hash)
		user.Password = hashComoCadena

		resp := client.InsertarUser(&user)

		fmt.Fprint(w, resp)

	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, e)
	}

}

//ListarFoto Función que devuelve las peticiones de la base de datos dado un filtro
func ListarFoto(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Incoming request from " + r.URL.EscapedPath())
	if r.URL.Path != PathListarFoto {
		http.NotFound(w, r)
		return
	}

	if r.Method != http.MethodPost {
		http.NotFound(w, r)
		return
	}

	lista := client.MostrarFoto()
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	respuesta, _ := json.Marshal(&lista)
	fmt.Fprint(w, string(respuesta))
}
