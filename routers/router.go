// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"pruebaRest/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// ns := beego.NewNamespace("/v1",
	// 	beego.NSNamespace("/object",
	// 		beego.NSInclude(
	// 			&controllers.ObjectController{},
	// 		),
	// 	),
	// 	beego.NSNamespace("/user",
	// 		beego.NSInclude(
	// 			&controllers.UserController{},
	// 		),
	// 	),
	// 	beego.NSNamespace("/usuarios",
	// 		beego.NSInclude(
	// 			&controllers.PageController{},
	// 		),
	// 	),
	// 	beego.NSNamespace("/pelicula",
	// 		beego.NSInclude(
	// 			&controllers.PeliculasController{},
	// 		),
	// 	),
	// )
	// beego.AddNamespace(ns)
	beego.Router("/", &controllers.PageController{}, "get:GetVer")
	beego.Router("/usuarios", &controllers.PageController{}, "get:GetAllUsuarios")
	// beego.Router("/addpeli", &controllers.PeliculasController{}, "post:CreatePelicula")
	beego.Router("/pelicula", &controllers.PeliculasController{}, "get:GetAll")
	beego.Router("/pelicula/:id", &controllers.PeliculasController{}, "get:GetOne")
	beego.Router("/pelicula/create", &controllers.PeliculasController{}, "post:Create")
	beego.Router("/pelicula/delete/:id", &controllers.PeliculasController{}, "delete:Delete")
	beego.Router("/pelicula/update/:id", &controllers.PeliculasController{}, "put:Update")

}
