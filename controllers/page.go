package controllers

import (
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	// "encoding/json"

	"github.com/astaxie/beego"
)

type PageController struct {
	beego.Controller
}

type Usuarios struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	UserName string
	Nombre   string
	Apellido string
	Edad     int32
}

func (o *PageController) GetAllUsuarios() {
	session, err := mgo.Dial("localhost:27017")

	if err != nil {
		fmt.Println("error en la conexion")
	} else {
		fmt.Println("todo correcto C:")
	}

	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("prueba").C("Usuarios")

	// err = c.Insert(
	// 	&Usuarios{"Usuario4", "Pedro", "Picapiedra", 35},
	// 	&Usuarios{"Usuario5", "bambam", "marmol", 16})

	// if err != nil {
	// 	fmt.Println("error al insertar")
	// } else {
	// 	fmt.Println("insercion correcta")
	// }

	// result := Usuarios{}

	// err = c.Find(bson.M{"username": "Usuario5"}).One(&result)
	// if err != nil {
	// 	fmt.Println("error al consultar")
	// }

	// fmt.Println("Phone:", result)

	var results []Usuarios

	err = c.Find(nil).All(&results)

	if err != nil {
		fmt.Println("Error de en la consulta ALL")
	} else {
		fmt.Println(results)
	}

}

func (o *PageController) GetVer() {
	fmt.Println("andamos viendo jijiji")
}
