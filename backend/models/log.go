package models

import "time"

type OperLog struct {
	ID             uint64    `json:"id" gorm:"primary_key;auto_increment"`
	Title          string    `json:"title" gorm:"type:varchar(64);comment:'模块标题'"`
	BusinessType   int       `json:"business_type" gorm:"type:int;default:0;comment:'业务类型 0其它 1新增 2修改 3删除'"`
	Method         string    `json:"method" gorm:"type:varchar(255);comment:'方法名称'"`
	RequestMethod  string    `json:"request_method" gorm:"type:varchar(16);comment:'请求方式'"`
	OperatorType   int       `json:"operator_type" gorm:"type:int;default:0;comment:'操作类别 0其它 1后台用户 2手机端用户'"`
	OperName       string    `json:"oper_name" gorm:"type:varchar(64);comment:'操作人员'"`
	DeptName       string    `json:"dept_name" gorm:"type:varchar(64);comment:'部门名称'"`
	OperURL        string    `json:"oper_url" gorm:"type:varchar(500);comment:'请求URL'"`
	OperIP         string    `json:"oper_ip" gorm:"type:varchar(128);comment:'主机地址'"`
	OperLocation   string    `json:"oper_location" gorm:"type:varchar(255);comment:'操作地点'"`
	OperParam      string    `json:"oper_param" gorm:"type:text;comment:'请求参数'"`
	JSONResult     string    `json:"json_result" gorm:"type:text;comment:'返回参数'"`
	Status         int       `json:"status" gorm:"type:int;default:0;comment:'操作状态 0正常 1异常'"`
	ErrorMsg       string    `json:"error_msg" gorm:"type:varchar(2000);comment:'错误消息'"`
	OperTime       time.Time `json:"oper_time" gorm:"type:datetime;default:CURRENT_TIMESTAMP;comment:'操作时间'"`
	CostTime       int64     `json:"cost_time" gorm:"type:bigint;default:0;comment:'消耗时间'"`
}

type LoginLog struct {
	ID          uint64    `json:"id" gorm:"primary_key;auto_increment"`
	LoginName   string    `json:"login_name" gorm:"type:varchar(64);comment:'登录账号'"`
	IPAddress   string    `json:"ip_address" gorm:"type:varchar(128);comment:'登录IP地址'"`
	LoginLocation string  `json:"login_location" gorm:"type:varchar(255);comment:'登录地点'"`
	Browser     string    `json:"browser" gorm:"type:varchar(64);comment:'浏览器类型'"`
	OS          string    `json:"os" gorm:"type:varchar(64);comment:'操作系统'"`
	Status      string    `json:"status" gorm:"type:char(1);default:'0';comment:'登录状态 0成功 1失败'"`
	Msg         string    `json:"msg" gorm:"type:varchar(255);comment:'提示消息'"`
	LoginTime   time.Time `json:"login_time" gorm:"type:datetime;default:CURRENT_TIMESTAMP;comment:'访问时间'"`
}
