package routers

import (

//	"github.com/astaxie/beego/context"

	"github.com/astaxie/beego"
	"microsvc01/controllers"
)

func Router()  {
	namespace:= beego.NewNamespace("/api/svc")
	namespace.Router("/",&controllers.MicroSVC01Controller{},"Get:Get")
	beego.AddNamespace(namespace)
}