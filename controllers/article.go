package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"path"
	"github.com/weilaihui/fdfs_client"

)

type ArticleController struct {
	beego.Controller
}

//展示首页
func(this*ArticleController)ShowIndex(){
	this.Layout = "layout.html"
	this.TplName = "index.html"
}

//展示添加分类
func(this*ArticleController)ShowAddType(){
	this.Layout = "layout.html"
	this.TplName = "addType.html"
}

//上传文件函数
func UploadFile(this*ArticleController,fileImage string)string{
	file,head ,err:= this.GetFile(fileImage)

	if err != nil {
		this.Data["errmsg"] = "获取文件失败，请重新添加"
		this.TplName = "add.html"
		return ""
	}

	//上传文件一般需要校验
	//1.文件大小
	if head.Size > 10000000{
		this.Data["errmsg"] = "图片太大，请重新选择"
		this.TplName = "add.html"
		return ""
	}
	//2.校验文件格式
	ext := path.Ext(head.Filename)
	fmt.Println("当前文章格式为:",ext)
	if ext != ".jpg" && ext != ".png" && ext != ".jpeg"{
		this.Data["errmsg"] = "文件格式错误，请重新选择"
		this.TplName = "add.html"
		return ""
	}
	//3.上传
	//fastDFS客户端(参数:客户端配置文件)
	fdfsClient,err := fdfs_client.NewFdfsClient("/etc/fdfs/client.conf")
	if err != nil {
		fmt.Println("获取fdfs客户端错误: ",err)
		return ""
	}
	//file--二进制文件流,我们用字节上传
	fileBuffer := make([]byte,head.Size)
	_,err = file.Read(fileBuffer)
	if err != nil {
		fmt.Println("写入文件错误,",err)
		return ""
	}

	//上传到storage
	uploadResponse,_ :=fdfsClient.UploadByBuffer(fileBuffer,ext[1:])
	fmt.Println(uploadResponse)



	return ""
}

//处理添加类型业务
func(this*ArticleController)HandleAddType(){
	//获取数据
	typeName := this.GetString("typeName")
	logoPath :=UploadFile(this,"uploadlogo")
	typeImagePath :=UploadFile(this,"uploadTypeImage")
	//校验数据
	if typeName == "" || logoPath == "" || typeImagePath == "" {
		fmt.Println("获取内容为空")
		return
	}

	//处理数据

	//返回数据
}