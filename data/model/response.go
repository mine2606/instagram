package model

//RUser struct
type RUser struct {
	ID       int
	Email    string
	Nombre   string
	Usuario  string
	Password string
}

//RLogin struct
type RLogin struct {
	Usuario  string
	Password string
}

//RRegistro struct
type RRegistro struct {
	Registro bool
}

//RFoto struct
type RFoto struct {
	ID    string
	Texto string
	URL   string
}
