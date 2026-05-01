package models

import (
	"cartoonManage/utils"
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	ID         uint64     `json:"id" gorm:"primary_key;auto_increment"`
	Username   string     `json:"username" gorm:"type:varchar(64);unique_index;not null;comment:'用户名'"`
	Password   string     `json:"-" gorm:"type:varchar(255);not null;comment:'密码'"`
	Nickname   string     `json:"nickname" gorm:"type:varchar(64);comment:'昵称'"`
	Email      string     `json:"email" gorm:"type:varchar(128);comment:'邮箱'"`
	Phone      string     `json:"phone" gorm:"type:varchar(32);comment:'手机号'"`
	Gender     string     `json:"gender" gorm:"type:char(1);default:'0';comment:'性别 0男 1女 2未知'"`
	Avatar     string     `json:"avatar" gorm:"type:varchar(255);comment:'头像'"`
	Status     string     `json:"status" gorm:"type:char(1);default:'0';comment:'状态 0正常 1停用'"`
	DeptID     uint64     `json:"dept_id" gorm:"comment:'部门ID'"`
	PostIDs    []uint64   `json:"post_ids" gorm:"-"`
	RoleIDs    []uint64   `json:"role_ids" gorm:"-"`
	Roles      []Role     `json:"roles" gorm:"many2many:sys_user_role;foreignkey:ID;association_foreignkey:ID"`
	Posts      []Post     `json:"posts" gorm:"many2many:sys_user_post;foreignkey:ID;association_foreignkey:ID"`
	CreateBy   string     `json:"create_by" gorm:"type:varchar(64);comment:'创建者'"`
	CreateTime time.Time  `json:"create_time" gorm:"type:datetime;default:CURRENT_TIMESTAMP;comment:'创建时间'"`
	UpdateBy   string     `json:"update_by" gorm:"type:varchar(64);comment:'更新者'"`
	UpdateTime time.Time  `json:"update_time" gorm:"type:datetime;default:CURRENT_TIMESTAMP;comment:'更新时间'"`
	Remark     string     `json:"remark" gorm:"type:varchar(500);comment:'备注'"`
	Deleted    int        `json:"-" gorm:"type:tinyint;default:0;comment:'删除标志 0存在 1删除'"`
}

func (u *User) BeforeCreate(scope *gorm.Scope) error {
	if u.Password != "" {
		hashedPassword, err := utils.HashPassword(u.Password)
		if err != nil {
			return err
		}
		scope.SetColumn("Password", hashedPassword)
	}
	return nil
}

func (u *User) BeforeUpdate(scope *gorm.Scope) error {
	if u.Password != "" {
		// 检查密码是否已经是加密后的（长度大于等于60通常表示已经加密）
		if len(u.Password) < 60 {
			hashedPassword, err := utils.HashPassword(u.Password)
			if err != nil {
				return err
			}
			scope.SetColumn("Password", hashedPassword)
		}
	}
	return nil
}
