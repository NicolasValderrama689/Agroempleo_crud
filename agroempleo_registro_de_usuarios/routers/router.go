// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/sena_2824182/Agroempleo_cud/agroempleo_registro_de_usuarios/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/informacionlaboral",
			beego.NSInclude(
				&controllers.InformacionlaboralController{},
			),
		),

		beego.NSNamespace("/informacion_academica_buscador",
			beego.NSInclude(
				&controllers.InformacionAcademicaBuscadorController{},
			),
		),

		beego.NSNamespace("/postulaciones",
			beego.NSInclude(
				&controllers.PostulacionesController{},
			),
		),

		beego.NSNamespace("/Learning",
			beego.NSInclude(
				&controllers.LearningController{},
			),
		),

		beego.NSNamespace("/rol",
			beego.NSInclude(
				&controllers.RolController{},
			),
		),

		beego.NSNamespace("/Usuarios",
			beego.NSInclude(
				&controllers.UsuariosController{},
			),
		),

		beego.NSNamespace("/Tipo_de_empleo",
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

		beego.NSNamespace("/contraseñas",
			beego.NSInclude(
				&controllers.ContraseñasController{},
			),
		),

		beego.NSNamespace("/Identificacion",
			beego.NSInclude(
				&controllers.IdentificacionController{},
			),
		),

		beego.NSNamespace("/tipo_documento",
			beego.NSInclude(
				&controllers.TipoDocumentoController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
