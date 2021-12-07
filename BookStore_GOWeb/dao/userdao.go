package dao

import (
	"github.com/Takeminesoce/BookStore/model"
	"github.com/Takeminesoce/BookStore/utils"
)

//依据用户名和密码从数据库中查询数据
func CheckByUsernameAndPassword(username string, password string) (*model.Users, error) {
	//数据库查询语句
	sqlstring := "SELECT id,username,password,email from users where username = ? and password = ?"
	//执行
	row := utils.Db.QueryRow(sqlstring, username, password)
	user := &model.Users{}
	row.Scan(&user.ID, &user.UserName, &user.PassWord, &user.Email)
	return user, nil
}

//依据用户名从数据库中查询数据，解决用户名是否重复，已存在的问题
func CheckByUsername(username string) (*model.Users, error) {
	//数据库查询语句
	sqlstring := "SELECT id,username,password,email from users where username = ? "
	//执行
	row := utils.Db.QueryRow(sqlstring, username)
	user := &model.Users{}
	row.Scan(&user.ID, &user.UserName, &user.PassWord, &user.Email)
	return user, nil
}

//SaveUser 向数据库中插入用户信息 注册
func SaveUser(username, password, email string) error {
	//Sql
	sqlstring := "insert into users(username,password,email) values(?,?,?)"
	//执行
	_, err := utils.Db.Exec(sqlstring, username, password, email)
	if err != nil {
		return err
	}
	return nil
}
