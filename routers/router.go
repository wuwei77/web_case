package routers

import (
	"newsWeb/controllers"
	"github.com/astaxie/beego"
    "github.com/astaxie/beego/context"
)

func init() {
    //路由过滤器
    beego.InsertFilter("/Article/*",beego.BeforeExec,filters)

    beego.Router("/register",&controllers.UserController{},"get:ShowRegister;post:HandleReg")
    //登录业务
    beego.Router("/login",&controllers.UserController{},"get:ShowLogin;post:HandleLogin")
    //首页展示
    beego.Router("/Article/index",&controllers.ArticleController{},"get:ShowIndex")
    //添加分类业务
    beego.Router("/Article/AddArticleType",&controllers.ArticleController{},"get:ShowAddType;post:HandleAddType")
}

func filters(ctx *context.Context){
    //检查是否登录
    userName := ctx.Input.Session("userName")
    if userName == nil{
        ctx.Redirect(302,"/login")
        return
    }
}
