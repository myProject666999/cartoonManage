package models

import "time"

type Config struct {
	ID          uint64    `json:"id" gorm:"primary_key;auto_increment"`
	ConfigName  string    `json:"config_name" gorm:"type:varchar(128);comment:'参数名称'"`
	ConfigKey   string    `json:"config_key" gorm:"type:varchar(128);unique_index;not null;comment:'参数键名'"`
	ConfigValue string    `json:"config_value" gorm:"type:varchar(500);comment:'参数键值'"`
	ConfigType  string    `json:"config_type" gorm:"type:char(1);default:'N';comment:'系统内置 Y是 N否'"`
	CreateBy    string    `json:"create_by" gorm:"type:varchar(64);comment:'创建者'"`
	CreateTime  time.Time `json:"create_time" gorm:"type:datetime;default:CURRENT_TIMESTAMP;comment:'创建时间'"`
	UpdateBy    string    `json:"update_by" gorm:"type:varchar(64);comment:'更新者'"`
	UpdateTime  time.Time `json:"update_time" gorm:"type:datetime;default:CURRENT_TIMESTAMP;comment:'更新时间'"`
	Remark      string    `json:"remark" gorm:"type:varchar(500);comment:'备注'"`
}
