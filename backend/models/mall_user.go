package models

import "time"

type MallUser struct {
	ID          uint64    `json:"id" gorm:"primary_key;auto_increment"`
	Username    string    `json:"username" gorm:"type:varchar(64);unique_index;not null;comment:'用户名'"`
	Password    string    `json:"-" gorm:"type:varchar(255);not null;comment:'密码'"`
	Nickname    string    `json:"nickname" gorm:"type:varchar(64);comment:'昵称'"`
	Avatar      string    `json:"avatar" gorm:"type:varchar(255);comment:'头像'"`
	Email       string    `json:"email" gorm:"type:varchar(128);comment:'邮箱'"`
	Phone       string    `json:"phone" gorm:"type:varchar(32);comment:'手机号'"`
	Gender      string    `json:"gender" gorm:"type:char(1);default:'0';comment:'性别 0男 1女 2未知'"`
	Birthday    time.Time `json:"birthday" gorm:"type:date;comment:'生日'"`
	Status      string    `json:"status" gorm:"type:char(1);default:'0';comment:'状态 0正常 1禁用'"`
	LastLoginIP string    `json:"last_login_ip" gorm:"type:varchar(128);comment:'最后登录IP'"`
	LastLoginTime time.Time `json:"last_login_time" gorm:"type:datetime;comment:'最后登录时间'"`
	CreateTime  time.Time `json:"create_time" gorm:"type:datetime;default:CURRENT_TIMESTAMP;comment:'创建时间'"`
	UpdateTime  time.Time `json:"update_time" gorm:"type:datetime;default:CURRENT_TIMESTAMP;comment:'更新时间'"`
	Deleted     int       `json:"-" gorm:"type:tinyint;default:0;comment:'删除标志 0存在 1删除'"`
}
