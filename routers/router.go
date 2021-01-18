package routers

import (
	"XinFeiWebPortal-API/controllers"
	"github.com/astaxie/beego"
)

func init() {
	// Web
	//beego.Router("/", &controllers.WebController{}, "GET:Home")

	// Admin
	beego.Router("/index.php/*", &controllers.AdminController{}, "GET:RedirectHome")

	beego.Router("admin/token", &controllers.AdminController{}, "GET:Token") // 验证token
	beego.Router("admin/welcome", &controllers.AdminController{}, "GET:Welcome")
	beego.Router("manage/request_login", &controllers.AdminController{}, "POST:AdminLogin")
	beego.Router("admin/exit", &controllers.AdminController{}, "GET:ExitLogin")
	// 成员
	beego.Router("admin/user/home", &controllers.UserController{}, "GET:UserHome")                     // 用户列表
	beego.Router("admin/user/add", &controllers.UserController{}, "POST:UserAdd")                      // 单个添加用户
	beego.Router("admin/user/activate", &controllers.UserController{}, "GET:UserActivate")             // 单个用户的激活或者禁用
	beego.Router("admin/user/delete", &controllers.UserController{}, "POST:UserDelete")                // 单个用户的删除
	beego.Router("admin/user/alter/password", &controllers.UserController{}, "POST:UserAlterPassword") // 修改用户密码
	beego.Router("admin/user/alter/user", &controllers.UserController{}, "POST:UserAlterInfo")         // 修改用户信息
	beego.Router("admin/user/query/user", &controllers.UserController{}, "POST:UserQuery")             // 条件查询
	// 文章
	beego.Router("admin/article/home", &controllers.ArticleController{}, "GET:ArticleHome")          //文章列表
	beego.Router("admin/article/query", &controllers.ArticleController{}, "GET:ArticleQueryAllDate") // 查询使用文章
	beego.Router("admin/article/alter", &controllers.ArticleController{}, "POST:ArticleAlter")       //修改文章
	beego.Router("admin/article/delete", &controllers.ArticleController{}, "POST:ArticleDeleteOne")  //删除一篇文章
	beego.Router("admin/article/add", &controllers.ArticleController{}, "POST:ArticleAddOne")        //增加一篇文章
	beego.Router("admin/article/:id", &controllers.ArticleController{}, "GET:ArticleQueryOne")       //根据文章id查询文章详情

	//uploading
	beego.Router("admin/uploading/image", &controllers.UploadingController{}, "POST:UploadingImage")      //上传图片
	beego.Router("admin/uploading/selfie", &controllers.UploadingController{}, "POST:UploadingSelfie")    //上传自拍照
	beego.Router("admin/uploading/honor", &controllers.UploadingController{}, "POST:UploadingHonerImage") //上传荣誉图
	beego.Router("admin/uploading/excel", &controllers.UploadingController{}, "POST:UploadingMailModel")  //上传群发模板

	// 获奖
	beego.Router("admin/honor/add", &controllers.HonorController{}, "POST:InsertHonor")     //插入荣誉信息
	beego.Router("admin/honor/home", &controllers.HonorController{}, "GET:GetHonorAllList") //获取荣誉列表信息
	beego.Router("admin/honor/cl", &controllers.HonorController{}, "GET:GetHonorList")      //根据分类获取荣誉列表信息
	beego.Router("admin/honor/delete", &controllers.HonorController{}, "POST:DeleteHonor")  //删除荣誉列表信息

	// 邮件
	beego.Router("admin/email/send/one", &controllers.MailController{}, "POST:SendEmaiOne")  //发送一封邮件
	beego.Router("admin/email/send/emails", &controllers.MailController{}, "POST:SendEmais") //发送一封邮件

	//API
	beego.Router("admin/getsystemstatus", &controllers.AdminController{}, "GET:GetSystemStatus") //返回JSON数据
	beego.Router("admin/reboot_system", &controllers.AdminController{}, "GET:RebootSystem")      //重启系统

}
