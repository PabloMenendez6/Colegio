package model

type Usuario struct {
	Id        int
	Nombre    string
	Apellido  string
	Rol       string
	Direccion string
	Telefono  string
	Estatus   bool
}

type Usuarios struct {
	Usuarios []Usuario
}

func (u *Usuarios) AgregarUsuario(us Usuario) {
	u.Usuarios = append(u.Usuarios, us)
}

// aca creamos una funcion que recibe un puntero de una lista de usuarios, le colocamos el nombre y el atributo que recibe esta funcion, para luego devolver un usuario y bool
func (u *Usuarios) BuscarUsuario(index int) (*Usuario, bool) {
	// hacemos un for recorrer la lista de usuarios
	for i := range u.Usuarios {
		// hacemos una comparacion que si el id del usuario recorrido es igual al que nosotros estamos mandando significa que lo encontro y que existe obviamente
		if u.Usuarios[i].Id == index {
			// aca ya devolvemos al usuario encontrado, y un true que significa que si lo encontramos digamos
			return &u.Usuarios[i], true
		}
	}
	// aca no devolvemos nada y un false porque eso significa que no se encontro nada
	return nil, false
}

func (u *Usuarios) ActualizarUsuario(usuario Usuario) bool {
	for i, v := range u.Usuarios {
		if v.Id == usuario.Id {
			u.Usuarios[i].Nombre = usuario.Nombre
			u.Usuarios[i].Apellido = usuario.Apellido
			u.Usuarios[i].Rol = usuario.Rol
			u.Usuarios[i].Direccion = usuario.Direccion
			u.Usuarios[i].Telefono = usuario.Telefono
			u.Usuarios[i].Estatus = usuario.Estatus
			return true

		}
	}
	return false
}

func (u *Usuarios) EliminarUsuario(id int) bool {
	for i, v := range u.Usuarios {
		if v.Id == id {
			u.Usuarios[i].Estatus = false
			return true
		}
	}
	return false
}
