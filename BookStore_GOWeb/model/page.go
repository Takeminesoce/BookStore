package model

//Page 结构
type Page struct {
	Books       []*Book
	PageNow     int64 //当前页面
	PageSize    int64 //每页显示总数
	TotalPage   int64 //总页数，计算可得
	TotalRecord int64 //总记录数，查询数据库可得
	MinPrice    string
	MaxPrice    string
	IsLogin     bool
	Username    string
}

//IsHasPrev 判断是否有上一页
func (p *Page) IsHasPrev() bool {
	return p.PageNow > 1
}

//IsHasNext 判断是否有下一页
func (p *Page) IsHasNext() bool {
	return p.PageNow < p.TotalPage
}

//GetPrevPageNo 获取上一页
func (p *Page) GetPrevPageNo() int64 {
	if p.IsHasPrev() {
		return p.PageNow - 1
	}
	return 1
}

//GetNextPageNo 获取下一页
func (p *Page) GetNextPageNo() int64 {
	if p.IsHasNext() {
		return p.PageNow + 1
	}
	return p.TotalPage

}
