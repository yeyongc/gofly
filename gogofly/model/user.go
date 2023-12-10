package model

import (
	"gogofly/utils"
	"time"

	"gorm.io/gorm"
)

// entity：对应数据库sys_user表
type User struct {
	ID        uint           `grom:"primarykey" json:"id" uri:"id" xml:"id" form:"id"`
	CreatedAt time.Time      `json:"created_time"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `grom:"index" json:"deleted_at" xml:"deleted_at" from:"deleted_at"`
	Name      string         `json:"name" gorm:"size:64; not null"`
	RealName  string         `json:"real_name" gorm:"size:128"`
	Avtar     string         `json:"avtar" gorm:"size:255"`
	Mobile    string         `json:"mobile" gorm:"size:11"`
	Email     string         `json:"email" gorm:"size:128"`
	Password  string         `json:"password" gorm:"size:128;not null"`
}

// 对密码进行加密
func (m *User) EncryptPassword() error {
	hash, err := utils.Encrypt(m.Password)
	if err == nil {
		m.Password = string(hash)
	}
	return err
}

// 实现gorm生命周期钩子: 在创建之前执行，无需手动调用
func (m *User) BeforeCreate(orm *gorm.DB) error {
	return m.EncryptPassword()
}

// 用户登录信息
type LoginUser struct {
	ID   uint   `json:"name"`
	Name string `json:"id"`
}
