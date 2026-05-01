package models

import "time"

type File struct {
	ID         uint64    `json:"id" gorm:"primary_key;auto_increment"`
	FileName   string    `json:"file_name" gorm:"type:varchar(255);not null;comment:'文件名称'"`
	FilePath   string    `json:"file_path" gorm:"type:varchar(500);not null;comment:'文件路径'"`
	FileSize   int64     `json:"file_size" gorm:"type:bigint;default:0;comment:'文件大小'"`
	FileType   string    `json:"file_type" gorm:"type:varchar(64);comment:'文件类型'"`
	MimeType   string    `json:"mime_type" gorm:"type:varchar(128);comment:'MIME类型'"`
	CreateBy   string    `json:"create_by" gorm:"type:varchar(64);comment:'创建者'"`
	CreateTime time.Time `json:"create_time" gorm:"type:datetime;default:CURRENT_TIMESTAMP;comment:'创建时间'"`
	UpdateBy   string    `json:"update_by" gorm:"type:varchar(64);comment:'更新者'"`
	UpdateTime time.Time `json:"update_time" gorm:"type:datetime;default:CURRENT_TIMESTAMP;comment:'更新时间'"`
}
