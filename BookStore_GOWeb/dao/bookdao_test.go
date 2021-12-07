package dao

import (
	"fmt"
	"testing"

	"github.com/Takeminesoce/BookStore/model"
)

//测试前准备
func TestMain(m *testing.M) {
	fmt.Println("测试bookdao.go中的全部方法")
	m.Run()
}

func TestBookdao(t *testing.T) {
	t.Run("测试获取全部图书", testgetbooks)
	t.Run("测试添加一本图书", testaddbook)
	t.Run("测试删除一本图书", testdeletebook)
	t.Run("测试添加一本图书", testgetbookById)
	t.Run("测试删除一本图书", updatebook)
}
func testgetbooks(t *testing.T) {
	books, _ := GetBooks()
	for k, v := range books {
		fmt.Printf("第%d本图书的信息是：%v\n", k+1, v)
	}
}
func testaddbook(t *testing.T) {
	book := &model.Book{
		Title:   "三国演义",
		Author:  "罗贯中",
		Price:   88.88,
		Sales:   100,
		Stock:   100,
		ImgPath: "/static/img/default.jpg",
	}
	AddBook(book)
}
func testdeletebook(t *testing.T) {
	id := "1"
	fmt.Printf("从数据库中删除id为%s的图书", id)
	DeleteBook(id)
}
func testgetbookById(t *testing.T) {
	id := "2"
	fmt.Printf("从数据库中查询id为%s的图书", id)
	book, _ := GetBookByID(id)
	println("从数据库中查询到图书的信息是：", book)
}
func updatebook(t *testing.T) {
	book := &model.Book{
		Title:   "三国演义",
		Author:  "罗贯中",
		Price:   88.88,
		Sales:   100,
		Stock:   100,
		ImgPath: "/static/img/default.jpg",
	}
	fmt.Printf("从数据库中更新书名为%s的图书", book.Title)
	UpdateBook(book)
}
