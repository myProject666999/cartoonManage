package models

import "time"

type Menu struct {
	ID         uint64     `json:"id" gorm:"primary_key;auto_increment"`
	MenuName   string     `json:"menu_name" gorm:"type:varchar(64);not null;comment:'菜单名称'"`
	ParentID   uint64     `json:"parent_id" gorm:"default:0;comment:'父菜单ID'"`
	OrderNum   int        `json:"order_num" gorm:"type:int;default:0;comment:'显示顺序'"`
	Path       string     `json:"path" gorm:"type:varchar(255);comment:'路由地址'"`
	Component  string     `json:"component" gorm:"type:varchar(255);comment:'组件路径'"`
	Query      string     `json:"query" gorm:"type:varchar(255);comment:'路由参数'"`
	RouteName  string     `json:"route_name" gorm:"type:varchar(64);comment:'路由名称'"`
	IsFrame    string     `json:"is_frame" gorm:"type:char(1);default:'1';comment:'是否为外链 0是 1否'"`
	IsCache    string     `json:"is_cache" gorm:"type:char(1);default:'0';comment:'是否缓存 0缓存 1不缓存'"`
	MenuType   string     `json:"menu_type" gorm:"type:char(1);default:'M';comment:'菜单类型 M目录 C菜单 F按钮'"`
	Visible    string     `json:"visible" gorm:"type:char(1);default:'0';comment:'显示状态 0显示 1隐藏'"`
	Status     string     `json:"status" gorm:"type:char(1);default:'0';comment:'菜单状态 0正常 1停用'"`
	Perms      string     `json:"perms" gorm:"type:varchar(128);comment:'权限标识'"`
	Icon       string     `json:"icon" gorm:"type:varchar(128);default:'#';comment:'菜单图标'"`
	CreateBy   string     `json:"create_by" gorm:"type:varchar(64);comment:'创建者'"`
	CreateTime time.Time  `json:"create_time" gorm:"type:datetime;default:CURRENT_TIMESTAMP;comment:'创建时间'"`
	UpdateBy   string     `json:"update_by" gorm:"type:varchar(64);comment:'更新者'"`
	UpdateTime time.Time  `json:"update_time" gorm:"type:datetime;default:CURRENT_TIMESTAMP;comment:'更新时间'"`
	Remark     string     `json:"remark" gorm:"type:varchar(500);comment:'备注'"`
	Children   []Menu     `json:"children,omitempty" gorm:"-"`
}
