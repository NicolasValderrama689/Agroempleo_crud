package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_crud/agroempleo_postulaciones/controllers:PostulacionesController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_crud/agroempleo_postulaciones/controllers:PostulacionesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_crud/agroempleo_postulaciones/controllers:PostulacionesController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_crud/agroempleo_postulaciones/controllers:PostulacionesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_crud/agroempleo_postulaciones/controllers:PostulacionesController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_crud/agroempleo_postulaciones/controllers:PostulacionesController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_crud/agroempleo_postulaciones/controllers:PostulacionesController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_crud/agroempleo_postulaciones/controllers:PostulacionesController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_crud/agroempleo_postulaciones/controllers:PostulacionesController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_crud/agroempleo_postulaciones/controllers:PostulacionesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
