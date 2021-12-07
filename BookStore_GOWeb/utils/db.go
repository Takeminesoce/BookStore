package utils

//实用工具包
import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Db  *sql.DB
	err error
)

func init() {
	//func sql.Open(driverName string, dataSourceName string) (*sql.DB, error)
	//"用户名：密码@协议（网络地址：端口号）/数据库名称"
	Db, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/bookstore")
	if err != nil {
		panic(err.Error())
	}
}
