package main

import (
	"fmt"
	"log"

	"main.go/config"
	"main.go/src/model"
)

func main() {
	db, err := config.DbConnection()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Conexion exitosa")

	err = db.AutoMigrate(&model.Usuario{})
	if err != nil {
		log.Fatal("Error al crear la tabla:", err)
	}

	fmt.Println("Tabla usuarios creada correctamente")
}
