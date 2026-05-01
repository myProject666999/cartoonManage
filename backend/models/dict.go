package models

import "time"

type DictType struct {
	ID         uint64    `json:"id" gorm:"primary_key;auto_increment"`
	DictName   string    `json:"dict_name" gorm:"type:varchar(64);comment:'字典名称'"`
	DictType   string    `json:"dict_type" gorm:"type:varchar(64);unique_index;not null;comment:'字典类型'"`
	Status     string    `json:"status" gorm:"type:char(1);default:'0';comment:'状态 0正常 1停用'"`
	CreateBy   string    `json:"create_by" gorm:"type:varchar(64);comment:'创建者'"`
	CreateTime time.Time `json:"create_time" gorm:"type:datetime;default:CURRENT_TIMESTAMP;comment:'创建时间'"`
	UpdateBy   string    `json:"update_by" gorm:"type:varchar(64);comment:'更新者'"`
	UpdateTime time.Time `json:"update_time" gorm:"type:datetime;default:CURRENT_TIMESTAMP;comment:'更新时间'"`
	Remark     string    `json:"remark" gorm:"type:varchar(500);comment:'备注'"`
}

type DictData struct {
	ID         uint64    `json:"id" gorm:"primary_key;auto_increment"`
	DictType   string    `json:"dict_type" gorm:"type:varchar(64);index;not null;comment:'字典类型'"`
	DictLabel  string    `json:"dict_label" gorm:"type:varchar(128);not null;comment:'字典标签'"`
	DictValue  string    `json:"dict_value" gorm:"type:varchar(255);not null;comment:'字典键值'"`
	DictSort   int       `json:"dict_sort" gorm:"type:int;default:0;comment:'字典排序'"`
	CSSClass   string    `json:"css_class" gorm:"type:varchar(128);comment:'样式属性'"`
	ListClass  string    `json:"list_class" gorm:"type:varchar(128);comment:'表格回显样式'"`
	IsDefault  string    `json:"is_default" gorm:"type:char(1);default:'N';comment:'是否默认 Y是 N否'"`
	Status     string    `json:"status" gorm:"type:char(1);default:'0';comment:'状态 0正常 1停用'"`
	CreateBy   string    `json:"create_by" gorm:"type:varchar(64);comment:'创建者'"`
	CreateTime time.Time `json:"create_time" gorm:"type:datetime;default:CURRENT_TIMESTAMP;comment:'创建时间'"`
	UpdateBy   string    `json:"update_by" gorm:"type:varchar(64);comment:'更新者'"`
	UpdateTime time.Time `json:"update_time" gorm:"type:datetime;default:CURRENT_TIMESTAMP;comment:'更新时间'"`
	Remark     string    `json:"remark" gorm:"type:varchar(500);comment:'备注'"`
}
