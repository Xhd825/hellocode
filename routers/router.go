package routers

import (
	"FirstHelloBeego/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//第一个参数是请求的路径，请求的路径必须保持唯一
	//第二个参数是
    beego.Router("/index", &controllers.MainController{})//处理请求的
    //加了login变成了404  因为

}
