package models

import (
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Peliculas struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	Titulo   string
	Fecha    int32
	Director string
}

func GetAllPelis() []Peliculas {
	session, err := mgo.Dial("localhost:27017")

	if err != nil {
		fmt.Println("error en la conexion")
	} else {
		fmt.Println("todo correcto C:")
	}

	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("prueba").C("Peliculas")

	var results []Peliculas

	err = c.Find(nil).All(&results)

	if err != nil {
		fmt.Println("Error de en la consulta ALL")
	}

	return results
}

func GetOnePeli(id string) Peliculas {
	session, err := mgo.Dial("localhost:27017")

	if err != nil {
		fmt.Println("error en la conexion")
	} else {
		fmt.Println("todo correcto C:")
	}

	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("prueba").C("Peliculas")

	result := Peliculas{}
	// ObjID := "ObjectId: '" + id + "'"

	if err = c.FindId(bson.ObjectIdHex(id)).One(&result); err != nil {
		fmt.Println("Error en la consulta")
	}

	return result
}

func AddPelicula(p Peliculas) {
	session, err := mgo.Dial("localhost:27017")

	if err != nil {
		panic(err.Error)
	}

	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("prueba").C("Peliculas")

	err = c.Insert(p)

	if err != nil {
		fmt.Println("Error al insertar")
	} else {
		fmt.Println("Correcto")
	}

}

func UpdatePeli(id string, pel Peliculas) {
	session, err := mgo.Dial("localhost:27017")

	if err != nil {
		panic(err.Error)
	}

	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("prueba").C("Peliculas")

	c.UpsertId(bson.ObjectIdHex(id), pel)
}

func DeletePeli(id string) {
	session, err := mgo.Dial("localhost:27017")

	if err != nil {
		panic(err.Error)
	}

	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("prueba").C("Peliculas")

	if err = c.RemoveId(bson.ObjectIdHex(id)); err != nil {
		fmt.Println("Eliminacion correcta")
	}
}
