// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/sena_2824182/Agroempleo_crud/agroempleo_registro_de_empleo/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/Tipodeempleo",
			beego.NSInclude(
				&controllers.TipoDeEmpleoController{},
			),
		),

		beego.NSNamespace("/Vacantes",
			beego.NSInclude(
				&controllers.VacantesController{},
			),
		),

		beego.NSNamespace("/Ciudad",
			beego.NSInclude(
				&controllers.CiudadController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
