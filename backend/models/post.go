package models

import "time"

type Post struct {
	ID         uint64    `json:"id" gorm:"primary_key;auto_increment"`
	PostCode   string    `json:"post_code" gorm:"type:varchar(64);unique_index;not null;comment:'岗位编码'"`
	PostName   string    `json:"post_name" gorm:"type:varchar(64);not null;comment:'岗位名称'"`
	OrderNum   int       `json:"order_num" gorm:"type:int;default:0;comment:'显示顺序'"`
	Status     string    `json:"status" gorm:"type:char(1);default:'0';comment:'状态 0正常 1停用'"`
	CreateBy   string    `json:"create_by" gorm:"type:varchar(64);comment:'创建者'"`
	CreateTime time.Time `json:"create_time" gorm:"type:datetime;default:CURRENT_TIMESTAMP;comment:'创建时间'"`
	UpdateBy   string    `json:"update_by" gorm:"type:varchar(64);comment:'更新者'"`
	UpdateTime time.Time `json:"update_time" gorm:"type:datetime;default:CURRENT_TIMESTAMP;comment:'更新时间'"`
	Remark     string    `json:"remark" gorm:"type:varchar(500);comment:'备注'"`
	Deleted    int       `json:"-" gorm:"type:tinyint;default:0;comment:'删除标志 0存在 1删除'"`
}
