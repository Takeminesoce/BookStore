package main

import (
	"net/http"

	"github.com/Takeminesoce/BookStore/controller"
)

func main() {
	//处理静态资源，如css和js文件
	//func http.StripPrefix(prefix string, h http.Handler) http.Handler
	//将请求去除给定前缀 然后用传入handler处理
	//func http.FileServer(root http.FileSystem) http.Handler
	//将给定目录下的静态文件转换为handler
	//func http.Handle(pattern string, handler http.Handler)
	//调用多路复用器中给定路径的处理器 处理
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static/"))))
	//视同静态资源 无需后台处理 直接访问页面
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("views/pages/"))))
	http.HandleFunc("/main", controller.GetPageBooksByPrice)
	//去登录
	http.HandleFunc("/login", controller.Login)
	//去注册
	http.HandleFunc("/regist", controller.Regist)
	//去注销
	http.HandleFunc("/logout", controller.Logout)
	//通过Ajex请求验证用户名是否存在
	//Ajex 一种新方法 不需要重新加载整个页面 与服务器交换数据更新部分网页内容
	http.HandleFunc("/checkUserName", controller.CheckUserName)
	//管理员系统
	//获取所有图书
	//http.HandleFunc("/getBooks", controller.GetBooks)
	//添加图书
	//http.HandleFunc("/addBook", controller.AddBook)
	//删除图书
	http.HandleFunc("/deleteBook", controller.DeleteBook)
	//去更新图书的页面
	http.HandleFunc("/toUpdateBookPage", controller.ToUpdateBookPage)
	//更新或添加图书
	http.HandleFunc("/updateOraddBook", controller.UpdateOrAddBook)
	//带分页的图书
	http.HandleFunc("/getPageBooks", controller.GetPageBooks)
	http.HandleFunc("/getPageBooksByPrice", controller.GetPageBooksByPrice)
	//添加图书到购物车中
	http.HandleFunc("/addBook2Cart", controller.AddBook2Cart)
	//获取购物车信息
	http.HandleFunc("/getCartInfo", controller.GetCartInfo)
	//清空购物车
	http.HandleFunc("/deleteCart", controller.DeleteCart)
	//删除购物项
	http.HandleFunc("/deleteCartItem", controller.DeleteCartItem)
	//更新购物项
	http.HandleFunc("/updateCartItem", controller.UpdateCartItem)
	//去结账
	http.HandleFunc("/checkout", controller.Checkout)
	//获取所有订单
	http.HandleFunc("/getOrders", controller.GetOrders)
	//获取订单详情，即订单所对应的所有的订单项
	http.HandleFunc("/getOrderInfo", controller.GetOrderInfo)
	//获取我的订单
	http.HandleFunc("/getMyOrder", controller.GetMyOrders)
	//发货
	http.HandleFunc("/sendOrder", controller.SendOrder)
	//确认收货
	http.HandleFunc("/takeOrder", controller.TakeOrder)
	//去购物车
	//http.HandleFunc("/pages/cart/cart.html", controller.GetCartInfo)
	http.ListenAndServe(":10086", nil)

}
