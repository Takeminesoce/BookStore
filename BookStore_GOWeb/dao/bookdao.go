package dao

import (
	"strconv"

	"github.com/Takeminesoce/BookStore/model"
	"github.com/Takeminesoce/BookStore/utils"
)

func GetBooks() ([]*model.Book, error) {
	//SQL获取语句
	sqlstring := "select id,title,author,price,sales,stock,img_path from books"
	//执行 Query返回一个队列数据
	rows, err := utils.Db.Query(sqlstring, sqlstring)
	if err != nil {
		return nil, err
	}
	var books []*model.Book
	for rows.Next() {
		var book *model.Book
		//赋值
		rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		books = append(books, book)
	}
	return books, nil
}

//AddBook 向数据库中添加一本图书
func AddBook(b *model.Book) error {
	//写sql语句
	slqStr := "insert into books(title,author,price,sales,stock,img_path) values(?,?,?,?,?,?)"
	//执行 Exec仅仅执行 不返还任何数据
	_, err := utils.Db.Exec(slqStr, b.Title, b.Author, b.Price, b.Sales, b.Stock, b.ImgPath)
	if err != nil {
		return err
	}
	return nil
}

//DeleteBook 根据图书的id从数据库中删除一本图书
func DeleteBook(bookID string) error {
	//写sql语句
	sqlStr := "delete from books where id = ?"
	//执行
	_, err := utils.Db.Exec(sqlStr, bookID)
	if err != nil {
		return err
	}
	return nil
}

//GetBookByID 根据图书的id从数据库中查询出一本图书
func GetBookByID(bookID string) (*model.Book, error) {
	//写sql语句
	sqlStr := "select id,title,author,price,sales,stock,img_path from books where id = ?"
	//执行 QueryRow最多返回一行数据
	row := utils.Db.QueryRow(sqlStr, bookID)
	//创建Book
	book := &model.Book{}
	//为book中的字段赋值，If more than one row matches the query,
	// Scan uses the first row and discards the rest. If no row matches the query, Scan returns ErrNoRows.
	row.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
	return book, nil
}

//UpdateBook 根据图书的id更新图书信息
func UpdateBook(b *model.Book) error {
	//写sql语句
	sqlStr := "update books set title=?,author=?,price=?,sales=?,stock=? where id=?"
	//执行
	_, err := utils.Db.Exec(sqlStr, b.Title, b.Author, b.Price, b.Sales, b.Stock, b.ID)
	if err != nil {
		return err
	}
	return nil
}

//GetPageBooks 获取带分页的图书信息
func GetPageBooks(pageNow string) (*model.Page, error) {
	//将页码转换为int64类型
	iPageNow, _ := strconv.ParseInt(pageNow, 10, 64)
	//获取数据库中图书的总记录数
	sqlStr := "select count(*) from books"
	//设置一个变量接收总记录数
	var totalRecord int64
	//执行
	row := utils.Db.QueryRow(sqlStr)
	row.Scan(&totalRecord)
	//设置每页只显示4条记录
	var pageSize int64 = 4
	//变量 总页数
	var totalPageNow int64
	if totalRecord%pageSize == 0 {
		totalPageNow = totalRecord / pageSize
	} else {
		totalPageNow = totalRecord/pageSize + 1
	}

	//获取当前页中的图书
	sqlStr2 := "select id,title,author,price,sales,stock,img_path from books limit ?,?"
	//执行 Query 返回一个队列
	rows, err := utils.Db.Query(sqlStr2, (iPageNow-1)*pageSize, pageSize)
	if err != nil {
		return nil, err
	}
	var books []*model.Book
	//Next检测下一元素是否存在 返回布尔类型 常与Scan关联使用 一条条完全复制到结果中
	for rows.Next() {
		book := &model.Book{}
		rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		//将book添加到books中
		books = append(books, book)
	}
	//创建page
	page := &model.Page{
		Books:       books,
		PageNow:     iPageNow,
		PageSize:    pageSize,
		TotalPage:   totalPageNow,
		TotalRecord: totalRecord,
	}
	return page, nil
}

//GetPageBooksByPrice 获取带分页和价格范围的图书信息
func GetPageBooksByPrice(pageNow string, minPrice string, maxPrice string) (*model.Page, error) {
	//将页码转换为int64类型
	iPageNow, _ := strconv.ParseInt(pageNow, 10, 64)
	//获取数据库中图书的总记录数
	sqlStr := "select count(*) from books where price between ? and ?"
	//设置一个变量接收总记录数
	var totalRecord int64
	//执行
	row := utils.Db.QueryRow(sqlStr, minPrice, maxPrice)
	row.Scan(&totalRecord)
	//设置每页只显示4条记录
	var pageSize int64 = 4
	//设置一个变量接收总页数
	var totalPageNow int64
	if totalRecord%pageSize == 0 {
		totalPageNow = totalRecord / pageSize
	} else {
		totalPageNow = totalRecord/pageSize + 1
	}
	//获取当前页中的图书 limit 4,4， 从第4行后再算4行
	sqlStr2 := "select id,title,author,price,sales,stock,img_path from books where price between ? and ? limit ?,?"
	//执行
	rows, err := utils.Db.Query(sqlStr2, minPrice, maxPrice, (iPageNow-1)*pageSize, pageSize)
	if err != nil {
		return nil, err
	}
	var books []*model.Book
	for rows.Next() {
		book := &model.Book{}
		rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		//将book添加到books中
		books = append(books, book)
	}
	//创建page
	page := &model.Page{
		Books:       books,
		PageNow:     iPageNow,
		PageSize:    pageSize,
		TotalPage:   totalPageNow,
		TotalRecord: totalRecord,
	}
	return page, nil
}
