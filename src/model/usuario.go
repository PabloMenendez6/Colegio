package model

type Usuario struct {
	Id        int `gorm:"primaryKey;autoIncrement"`
	Nombre    string
	Apellido  string
	Rol       string
	Direccion string
	Telefono  string
	Estatus   bool
}
