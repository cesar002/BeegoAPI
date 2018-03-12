package controllers

import (
	"fmt"
	"pruebaRest/models"

	"github.com/astaxie/beego"
)

type PeliculasController struct {
	beego.Controller
}

func (p *PeliculasController) Create() {

	pel := models.Peliculas{}

	if err := p.ParseForm(&pel); err != nil {
		fmt.Println("error")
	}

	models.AddPelicula(pel)
}

func (p *PeliculasController) GetAll() {
	pelis := models.GetAllPelis()
	p.Data["json"] = pelis
	p.ServeJSON()
}

func (p *PeliculasController) GetOne() {
	id := p.GetString(":id")

	peli := models.GetOnePeli(id)

	p.Data["json"] = peli
	p.ServeJSON()

}

func (p *PeliculasController) Delete() {
	id := p.GetString(":id")

	models.DeletePeli(id)

}

func (p *PeliculasController) Update() {
	id := p.GetString(":id")

	pel := models.Peliculas{}

	if err := p.ParseForm(&pel); err != nil {
		fmt.Println("error")
	} else {
		models.UpdatePeli(id, pel)
	}
}
