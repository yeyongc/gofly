package dto

import "gogofly/model"

// 登录DTO
type UserLoginDTO struct {
	Name     string `form:"name" json:"name" xml:"name" binding:"required" message:"用户名填写错误" required_err:"用户名不能为空"`
	Password string `form:"password" json:"password" xml:"password" binding:"required" message:"密码填写错误" required_err:"密码不能为空"`
}

// 添加用户DTO
type UserAddDTO struct {
	Name     string `json:"name" form:"name" xml:"name" binding:"required" message:"username not null"`
	RealName string `json:"real_name" form:"real_name" xml:"real_name"`
	Avtar    string `json:"avtar" form:"avtar" xml:"avtar"`
	Mobile   string `json:"mobile" form:"mobile" xml:"mobile"`
	Email    string `json:"email" form:"email" xml:"email"`
	Password string `json:"password,omitempty" form:"password" xml:"password" binding:"required" message:"username not null"`
	ID       uint
}

// 用户列表DTO
type UserListDTO struct {
	Pagination
}

// 更新userDTO
type UserUpdatetDTO struct {
	Name     string `json:"name" form:"name" xml:"name" `
	RealName string `json:"real_name" form:"real_name" xml:"real_name"`
	Mobile   string `json:"mobile" form:"mobile" xml:"mobile"`
	Email    string `json:"email" form:"email" xml:"email"`
	Password string `json:"password,omitempty" form:"password" xml:"password"`
	ID       uint   `json:"user_id" form:"user_id" xml:"user_id" uri:"id"`
}

func (u *UserAddDTO) ConvertToModel(user *model.User) {
	user.Name = u.Name
	user.Avtar = u.Avtar
	user.Password = u.Password
	user.Mobile = u.Mobile
	user.Email = u.Email
	user.RealName = u.RealName
}

func (u *UserUpdatetDTO) ConvertToModel(user *model.User) {
	user.ID = u.ID
	user.Name = u.Name
	user.RealName = u.RealName
	user.Email = u.Email
	user.Password = u.Password
	user.Mobile = u.Mobile
}
