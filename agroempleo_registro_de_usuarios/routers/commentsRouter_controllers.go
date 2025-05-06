package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_crud/agroempleo_registro_de_usuarios/controllers:ContraseñasController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_crud/agroempleo_registro_de_usuarios/controllers:ContraseñasController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_crud/agroempleo_registro_de_usuarios/controllers:ContraseñasController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_crud/agroempleo_registro_de_usuarios/controllers:ContraseñasController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_crud/agroempleo_registro_de_usuarios/controllers:ContraseñasController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_crud/agroempleo_registro_de_usuarios/controllers:ContraseñasController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_crud/agroempleo_registro_de_usuarios/controllers:ContraseñasController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_crud/agroempleo_registro_de_usuarios/controllers:ContraseñasController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_crud/agroempleo_registro_de_usuarios/controllers:ContraseñasController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_crud/agroempleo_registro_de_usuarios/controllers:ContraseñasController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_crud/agroempleo_registro_de_usuarios/controllers:RolController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_crud/agroempleo_registro_de_usuarios/controllers:RolController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_crud/agroempleo_registro_de_usuarios/controllers:RolController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_crud/agroempleo_registro_de_usuarios/controllers:RolController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_crud/agroempleo_registro_de_usuarios/controllers:RolController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_crud/agroempleo_registro_de_usuarios/controllers:RolController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_crud/agroempleo_registro_de_usuarios/controllers:RolController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_crud/agroempleo_registro_de_usuarios/controllers:RolController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_crud/agroempleo_registro_de_usuarios/controllers:RolController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_crud/agroempleo_registro_de_usuarios/controllers:RolController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_crud/agroempleo_registro_de_usuarios/controllers:TipoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_crud/agroempleo_registro_de_usuarios/controllers:TipoDocumentoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_crud/agroempleo_registro_de_usuarios/controllers:TipoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_crud/agroempleo_registro_de_usuarios/controllers:TipoDocumentoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_crud/agroempleo_registro_de_usuarios/controllers:TipoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_crud/agroempleo_registro_de_usuarios/controllers:TipoDocumentoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_crud/agroempleo_registro_de_usuarios/controllers:TipoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_crud/agroempleo_registro_de_usuarios/controllers:TipoDocumentoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_crud/agroempleo_registro_de_usuarios/controllers:TipoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_crud/agroempleo_registro_de_usuarios/controllers:TipoDocumentoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_crud/agroempleo_registro_de_usuarios/controllers:UsuariosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_crud/agroempleo_registro_de_usuarios/controllers:UsuariosController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_crud/agroempleo_registro_de_usuarios/controllers:UsuariosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_crud/agroempleo_registro_de_usuarios/controllers:UsuariosController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_crud/agroempleo_registro_de_usuarios/controllers:UsuariosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_crud/agroempleo_registro_de_usuarios/controllers:UsuariosController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_crud/agroempleo_registro_de_usuarios/controllers:UsuariosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_crud/agroempleo_registro_de_usuarios/controllers:UsuariosController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_crud/agroempleo_registro_de_usuarios/controllers:UsuariosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_crud/agroempleo_registro_de_usuarios/controllers:UsuariosController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
