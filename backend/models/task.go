package models

import "time"

type Task struct {
	ID          uint64    `json:"id" gorm:"primary_key;auto_increment"`
	TaskName    string    `json:"task_name" gorm:"type:varchar(64);not null;comment:'任务名称'"`
	TaskGroup   string    `json:"task_group" gorm:"type:varchar(64);comment:'任务组名'"`
	InvokeTarget string   `json:"invoke_target" gorm:"type:varchar(500);not null;comment:'调用目标字符串'"`
	CronExpression string `json:"cron_expression" gorm:"type:varchar(255);comment:'cron执行表达式'"`
	MisfirePolicy string  `json:"misfire_policy" gorm:"type:varchar(32);default:'3';comment:'计划执行策略 1立即执行 2执行一次 3放弃执行'"`
	Concurrent  string    `json:"concurrent" gorm:"type:char(1);default:'1';comment:'是否并发执行 0允许 1禁止'"`
	Status      string    `json:"status" gorm:"type:char(1);default:'0';comment:'状态 0正常 1暂停'"`
	CreateBy    string    `json:"create_by" gorm:"type:varchar(64);comment:'创建者'"`
	CreateTime  time.Time `json:"create_time" gorm:"type:datetime;default:CURRENT_TIMESTAMP;comment:'创建时间'"`
	UpdateBy    string    `json:"update_by" gorm:"type:varchar(64);comment:'更新者'"`
	UpdateTime  time.Time `json:"update_time" gorm:"type:datetime;default:CURRENT_TIMESTAMP;comment:'更新时间'"`
	Remark      string    `json:"remark" gorm:"type:varchar(500);comment:'备注'"`
}

type TaskLog struct {
	ID          uint64    `json:"id" gorm:"primary_key;auto_increment"`
	TaskName    string    `json:"task_name" gorm:"type:varchar(64);not null;comment:'任务名称'"`
	TaskGroup   string    `json:"task_group" gorm:"type:varchar(64);comment:'任务组名'"`
	InvokeTarget string   `json:"invoke_target" gorm:"type:varchar(500);not null;comment:'调用目标字符串'"`
	TaskMessage string    `json:"task_message" gorm:"type:varchar(500);comment:'日志信息'"`
	Status      string    `json:"status" gorm:"type:char(1);default:'0';comment:'执行状态 0正常 1失败'"`
	ExceptionInfo string   `json:"exception_info" gorm:"type:varchar(2000);comment:'异常信息'"`
	CreateTime  time.Time `json:"create_time" gorm:"type:datetime;default:CURRENT_TIMESTAMP;comment:'创建时间'"`
	CostTime    int64     `json:"cost_time" gorm:"type:bigint;default:0;comment:'消耗时间'"`
}
