package dao

import (
	"fmt"
	"testing"
)

func TestUser(t *testing.T) {
	fmt.Println("测试登录注册查询函数")
	t.Run("验证插入数据", testSave)
	t.Run("验证用户名和密码正确性", testLogin)
	t.Run("验证用户名唯一性 防止注册时重复", testRegist)

}
func testLogin(t *testing.T) {
	user, _ := CheckByUsernameAndPassword("admin", "0000000")
	fmt.Println("获取用户信息是：", user)
}
func testRegist(t *testing.T) {
	user, _ := CheckByUsername("admin")
	fmt.Println("获取用户信息是：", user)
}
func testSave(t *testing.T) {
	SaveUser("admin", "123456", "admin@github.com")
}
