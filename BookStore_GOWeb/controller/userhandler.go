package controller

import (
	"html/template"
	"net/http"

	"github.com/Takeminesoce/BookStore/dao"
	"github.com/Takeminesoce/BookStore/model"
	"github.com/Takeminesoce/BookStore/utils"
)

//Login 处理用户登录的函数
func Login(w http.ResponseWriter, r *http.Request) {
	//判断是否重复登录
	flag, _ := dao.IsLogin(r)
	if flag {
		//已登录 去首页
		GetPageBooksByPrice(w, r)
	} else {
		//获取用户名和密码
		username := r.PostFormValue("username")
		password := r.PostFormValue("password")
		//调用userdao中验证用户名和密码的方法
		user, _ := dao.CheckByUsernameAndPassword(username, password)
		if user.ID > 0 {

			//fmt.Println(user)
			//用户名和密码正确
			//生成UUID作为Session的id
			uuid := utils.CreateUUID()
			//创建一个Session
			sess := &model.Session{
				SessionID: uuid,
				UserName:  user.UserName,
				UserID:    user.ID,
			}
			//将Session保存到数据库中
			dao.AddSession(sess)
			//创建一个Cookie，让它与Session相关联
			cookie := http.Cookie{
				Name:     "user",
				Value:    uuid,
				HttpOnly: true,
			}
			//将cookie发送给浏览器
			http.SetCookie(w, &cookie)
			t := template.Must(template.ParseFiles("views/pages/user/login_success.html"))
			t.Execute(w, user)
		} else {
			//用户名或密码不正确
			t := template.Must(template.ParseFiles("views/pages/user/login.html"))
			t.Execute(w, "用户名或密码不正确！")
		}
	}

}

//Logout //处理用户注销的函数
func Logout(w http.ResponseWriter, r *http.Request) {
	//获取Cookie
	cookie, _ := r.Cookie("user")
	if cookie != nil {
		//获取cookie的value值
		cookieValue := cookie.Value
		//删除数据库中与之对应的Session
		dao.DeleteSession(cookieValue)
		//设置cookie失效
		cookie.MaxAge = -1
		//将修改之后的cookie发送给浏览器
		http.SetCookie(w, cookie)
	}
	//去首页
	GetPageBooksByPrice(w, r)
}

//Regist 处理用户注册的函数
func Regist(w http.ResponseWriter, r *http.Request) {
	//获取用户名和密码
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	email := r.PostFormValue("email")
	//调用userdao中验证用户名和密码的方法
	user, _ := dao.CheckByUsername(username)
	if user.ID > 0 {
		//用户名已经存在
		//fmt.Println(user)
		t := template.Must(template.ParseFiles("views/pages/user/regist.html"))
		t.Execute(w, "用户名已经存在!")
	} else {
		//用户名不存在,保存到数据库
		dao.SaveUser(username, password, email)
		t := template.Must(template.ParseFiles("views/pages/user/regist_success.html"))
		//对应欢迎“用户名”页面
		t.Execute(w, username)
	}
}

//通过Ajax发送请求验证用户名
func CheckUserName(w http.ResponseWriter, r *http.Request) {
	//获取用户输入的用户名
	username := r.PostFormValue("username")
	//调用userdao中验证用户名和密码的方法
	user, _ := dao.CheckByUsername(username)
	if user.ID > 0 {
		//用户名已存在
		w.Write([]byte("用户名已存在！"))
	} else {
		//用户名可用
		w.Write([]byte("<font style='color:green'>用户名可用！</font>"))
	}
}
