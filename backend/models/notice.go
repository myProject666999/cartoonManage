package models

import "time"

type Notice struct {
	ID          uint64    `json:"id" gorm:"primary_key;auto_increment"`
	NoticeTitle string    `json:"notice_title" gorm:"type:varchar(255);not null;comment:'公告标题'"`
	NoticeType  string    `json:"notice_type" gorm:"type:char(1);not null;comment:'公告类型 1通知 2公告'"`
	NoticeContent string  `json:"notice_content" gorm:"type:text;comment:'公告内容'"`
	Status      string    `json:"status" gorm:"type:char(1);default:'0';comment:'状态 0正常 1关闭'"`
	CreateBy    string    `json:"create_by" gorm:"type:varchar(64);comment:'创建者'"`
	CreateTime  time.Time `json:"create_time" gorm:"type:datetime;default:CURRENT_TIMESTAMP;comment:'创建时间'"`
	UpdateBy    string    `json:"update_by" gorm:"type:varchar(64);comment:'更新者'"`
	UpdateTime  time.Time `json:"update_time" gorm:"type:datetime;default:CURRENT_TIMESTAMP;comment:'更新时间'"`
	Remark      string    `json:"remark" gorm:"type:varchar(500);comment:'备注'"`
}
