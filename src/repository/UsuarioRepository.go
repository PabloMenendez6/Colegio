package repository

import (
	"gorm.io/gorm"
	"main.go/src/model"
)

type UsuarioRepository struct {
	DB *gorm.DB
}

func (r *UsuarioRepository) AgregarUsuario(usuario *model.Usuario) error {
	result := r.DB.Create(usuario)

	return result.Error
}

func (r *UsuarioRepository) BuscarUsuario(id int) (*model.Usuario, error) {
	var usuario model.Usuario
	result := r.DB.First(&usuario, id)

	return &usuario, result.Error
}

func (r *UsuarioRepository) EliminarUsuario(id int) error {
	result := r.DB.Model(&model.Usuario{}).
		Where("id = ?", id).
		Update("estatus", false)

	return result.Error
}

func (r *UsuarioRepository) ActualizarUsuario(id int, usuario *model.Usuario) error {
	result := r.DB.Model(&model.Usuario{}).
		Where("id = ?", id).
		Updates(usuario)

	return result.Error
}
