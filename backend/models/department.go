package models

import "time"

type Department struct {
	ID         uint64     `json:"id" gorm:"primary_key;auto_increment"`
	ParentID   uint64     `json:"parent_id" gorm:"default:0;comment:'父部门ID'"`
	DeptName   string     `json:"dept_name" gorm:"type:varchar(64);not null;comment:'部门名称'"`
	OrderNum   int        `json:"order_num" gorm:"type:int;default:0;comment:'显示顺序'"`
	Leader     string     `json:"leader" gorm:"type:varchar(64);comment:'负责人'"`
	Phone      string     `json:"phone" gorm:"type:varchar(32);comment:'联系电话'"`
	Email      string     `json:"email" gorm:"type:varchar(128);comment:'邮箱'"`
	Status     string     `json:"status" gorm:"type:char(1);default:'0';comment:'状态 0正常 1停用'"`
	CreateBy   string     `json:"create_by" gorm:"type:varchar(64);comment:'创建者'"`
	CreateTime time.Time  `json:"create_time" gorm:"type:datetime;default:CURRENT_TIMESTAMP;comment:'创建时间'"`
	UpdateBy   string     `json:"update_by" gorm:"type:varchar(64);comment:'更新者'"`
	UpdateTime time.Time  `json:"update_time" gorm:"type:datetime;default:CURRENT_TIMESTAMP;comment:'更新时间'"`
	Deleted    int        `json:"-" gorm:"type:tinyint;default:0;comment:'删除标志 0存在 1删除'"`
	Children   []Department `json:"children,omitempty" gorm:"-"`
}
