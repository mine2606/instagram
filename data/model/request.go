package model

import "time"

//Foto struct
type Foto struct {
	Texto string
	URL   string
}

//Comentario struct
type Comentario struct {
	Texto string
}

//User struct
type User struct {
	Nombre   string
	Usuario  string
	Email    string
	Password string
}

//Filtro struct
type Filtro struct {
	Fecha time.Time
}

//Login struct
type Login struct {
	Usuario  string
	Password string
}
