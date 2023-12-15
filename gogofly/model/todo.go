package model

//==============================================================================================================================
//	用户todo，对应sys_user_todo表：
//		    ID: todo id, [uint], 待做事项的id
//		 Title: todo title, [string], 待做事项的标题
//     HasDone：todo hasdone, [bool], 代做事项的状态（是否完成）
//		Detail: todo detail, [string], 待做事项的细节
//      UserID: 用户todo DTO UserID，[uint]，代做事项的所有者ID
//==============================================================================================================================
type UserTodo struct {
	ID      uint   `gorm:"primarykey" json:"todo_id" form:"todo_id"`
	Title   string `gorm:"size:128" json:"todo_title" form:"todo_title"`
	HasDone bool   `gorm:"bool" json:"todo_has_done" form:"todo_has_done"`
	Detail  string `gorm:"size:1024" json:"todo_detail" form:"todo_detail"`
	UserID  uint   `gorm:"not null" json:"owner_id" form:"owner_id"`
}
