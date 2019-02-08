package handlers

import "net/http"

//PathInicio Ruta raíz
const PathInicio string = "/"

//PathJSFiles Ruta a la carpeta de scripts de javascript
const PathJSFiles string = "/js/"

//PathCSSFiles Ruta a la carpeta de scripts de css
const PathCSSFiles string = "/css/"

//PathRegistroUsuario Ruta para el registro
const PathRegistroUsuario string = "/registro"

//PathRegistroFile Ruta para el fichero registro.html
const PathRegistroFile string = "/registroFile"

//PathPerfil Ruta para el registro
const PathPerfil string = "/perfil"

//PathEnvioPeticion Ruta de envío de peticiones
const PathEnvioPeticion string = "/envio"

//PathUploader Ruta de envío de una foto
const PathUploader string = "/uploader"

//PathFoto Ruta de envío de una foto
const PathFoto string = "/foto"

//PathLogin Ruta de envío de inicio de sesion
const PathLogin string = "/login"

//PathLogout Ruta de envío de inicio de sesion
const PathLogout string = "/logout"

//PathListarFoto Ruta de envío de una foto
const PathListarFoto string = "/listarFoto"

//PathNombreUsuario Ruta de envío de una foto
const PathNombreUsuario string = "/nombreUsuario"

//ManejadorHTTP encapsula como tipo la función de manejo de peticiones HTTP, para que sea posible almacenar sus referencias en un diccionario
type ManejadorHTTP = func(w http.ResponseWriter, r *http.Request)

//Manejadores Lista es el diccionario general de las peticiones que son manejadas por nuestro servidor
var Manejadores map[string]ManejadorHTTP

func init() {
	Manejadores = make(map[string]ManejadorHTTP)
	Manejadores[PathInicio] = IndexFile
	Manejadores[PathJSFiles] = JsFile
	Manejadores[PathCSSFiles] = CSSFile
	Manejadores[PathFoto] = FotoFile
	Manejadores[PathUploader] = Uploader
	Manejadores[PathRegistroUsuario] = RegistroUsuario
	Manejadores[PathRegistroFile] = RegistroFile
	Manejadores[PathPerfil] = Perfil
	Manejadores[PathLogin] = Login
	Manejadores[PathLogout] = Logout
	Manejadores[PathListarFoto] = ListarFoto
	Manejadores[PathNombreUsuario] = NombreUsuario
}
