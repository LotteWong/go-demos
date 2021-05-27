package models

type User struct {
	Id       string `json:"id" form:"id" uri:"id" example:"1"`         // 用户标识
	Username string `json:"username" form:"username" example:"admin"`  // 登录名称
	Password string `json:"password" form:"password" example:"123456"` // 登录密码
}

type Users struct {
	TotalCount int    `json:"totalcount" form:"totalcount" example:"0"` // 共计条数
	Items      []User `json:"items" form:"items" example:""`            // 用户列表
}
