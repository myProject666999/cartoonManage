package models

import "time"

type Role struct {
	ID         uint64    `json:"id" gorm:"primary_key;auto_increment"`
	RoleName   string    `json:"role_name" gorm:"type:varchar(64);not null;comment:'角色名称'"`
	RoleKey    string    `json:"role_key" gorm:"type:varchar(64);unique_index;not null;comment:'角色权限字符串'"`
	OrderNum   int       `json:"order_num" gorm:"type:int;default:0;comment:'显示顺序'"`
	Status     string    `json:"status" gorm:"type:char(1);default:'0';comment:'状态 0正常 1停用'"`
	MenuIDs    []uint64  `json:"menu_ids" gorm:"-"`
	Menus      []Menu    `json:"menus" gorm:"many2many:sys_role_menu;foreignkey:ID;association_foreignkey:ID"`
	CreateBy   string    `json:"create_by" gorm:"type:varchar(64);comment:'创建者'"`
	CreateTime time.Time `json:"create_time" gorm:"type:datetime;default:CURRENT_TIMESTAMP;comment:'创建时间'"`
	UpdateBy   string    `json:"update_by" gorm:"type:varchar(64);comment:'更新者'"`
	UpdateTime time.Time `json:"update_time" gorm:"type:datetime;default:CURRENT_TIMESTAMP;comment:'更新时间'"`
	Remark     string    `json:"remark" gorm:"type:varchar(500);comment:'备注'"`
	Deleted    int       `json:"-" gorm:"type:tinyint;default:0;comment:'删除标志 0存在 1删除'"`
}
