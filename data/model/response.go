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
