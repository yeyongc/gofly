package dto

import "gogofly/model"

//==============================================================================================================================
//	用户todo Model对应的DTO（数据传输对象）：
//		ID: 用户todo DTO id，[uint]，待做事项id
//   Title: 用户todo DTO title，[string]，待做事项title
// HasDone: 用户todo DTO hasdone，[bool]，待做事项hasdone
//  Detail: 用户todo DTO detail，[string]，待做事项detail
// UserID: 用户todo DTO UserID，[uint]，代做事项的所有者ID
//==============================================================================================================================
type UserTodoDTO struct {
	ID      uint   `json:"todo_id" form:"todo_id"`
	Title   string `json:"todo_title" form:"todo_title"`
	HasDone bool   `json:"todo_hasdone" form:"todo_hasdone"`
	Detail  string `json:"todo_detail" form:"todo_detail"`
	UserID  uint   `json:"owner_id" form:"owner_id"`
}

func (u *UserTodoDTO) ConvertToModel(userTodo *model.UserTodo) {
	userTodo.Detail = u.Detail
	userTodo.HasDone = u.HasDone
	userTodo.Title = u.Title
	userTodo.UserID = u.UserID
}
