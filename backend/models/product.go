package models

import "time"

type ProductCategory struct {
	ID         uint64            `json:"id" gorm:"primary_key;auto_increment"`
	ParentID   uint64            `json:"parent_id" gorm:"default:0;comment:'父分类ID'"`
	CategoryName string          `json:"category_name" gorm:"type:varchar(128);not null;comment:'分类名称'"`
	Level      int               `json:"level" gorm:"type:int;default:1;comment:'分类层级'"`
	OrderNum   int               `json:"order_num" gorm:"type:int;default:0;comment:'显示顺序'"`
	Icon       string            `json:"icon" gorm:"type:varchar(255);comment:'分类图标'"`
	Status     string            `json:"status" gorm:"type:char(1);default:'0';comment:'状态 0正常 1停用'"`
	CreateTime time.Time         `json:"create_time" gorm:"type:datetime;default:CURRENT_TIMESTAMP;comment:'创建时间'"`
	UpdateTime time.Time         `json:"update_time" gorm:"type:datetime;default:CURRENT_TIMESTAMP;comment:'更新时间'"`
	Children   []ProductCategory `json:"children,omitempty" gorm:"-"`
}

type Product struct {
	ID             uint64    `json:"id" gorm:"primary_key;auto_increment"`
	ProductName    string    `json:"product_name" gorm:"type:varchar(255);not null;comment:'商品名称'"`
	ProductCode    string    `json:"product_code" gorm:"type:varchar(64);unique_index;not null;comment:'商品编码'"`
	CategoryID     uint64    `json:"category_id" gorm:"comment:'分类ID'"`
	CategoryName   string    `json:"category_name" gorm:"type:varchar(128);comment:'分类名称'"`
	Brand          string    `json:"brand" gorm:"type:varchar(128);comment:'品牌'"`
	Subtitle       string    `json:"subtitle" gorm:"type:varchar(500);comment:'商品副标题'"`
	MainImage      string    `json:"main_image" gorm:"type:varchar(255);comment:'主图'"`
	SubImages      string    `json:"sub_images" gorm:"type:text;comment:'副图，JSON数组'"`
	Detail         string    `json:"detail" gorm:"type:text;comment:'商品详情'"`
	OriginalPrice  float64   `json:"original_price" gorm:"type:decimal(10,2);comment:'原价'"`
	CurrentPrice   float64   `json:"current_price" gorm:"type:decimal(10,2);comment:'现价'"`
	Stock          int       `json:"stock" gorm:"type:int;default:0;comment:'库存'"`
	Unit           string    `json:"unit" gorm:"type:varchar(32);comment:'单位'"`
	Weight         float64   `json:"weight" gorm:"type:decimal(10,2);comment:'重量(kg)'"`
	Status         string    `json:"status" gorm:"type:char(1);default:'0';comment:'状态 0上架 1下架'"`
	IsHot          string    `json:"is_hot" gorm:"type:char(1);default:'0';comment:'是否热销 0否 1是'"`
	IsNew          string    `json:"is_new" gorm:"type:char(1);default:'0';comment:'是否新品 0否 1是'"`
	IsRecommend    string    `json:"is_recommend" gorm:"type:char(1);default:'0';comment:'是否推荐 0否 1是'"`
	Sort           int       `json:"sort" gorm:"type:int;default:0;comment:'排序'"`
	Sales          int       `json:"sales" gorm:"type:int;default:0;comment:'销量'"`
	CreateTime     time.Time `json:"create_time" gorm:"type:datetime;default:CURRENT_TIMESTAMP;comment:'创建时间'"`
	UpdateTime     time.Time `json:"update_time" gorm:"type:datetime;default:CURRENT_TIMESTAMP;comment:'更新时间'"`
	Deleted        int       `json:"-" gorm:"type:tinyint;default:0;comment:'删除标志 0存在 1删除'"`
}

type ProductAttribute struct {
	ID          uint64    `json:"id" gorm:"primary_key;auto_increment"`
	ProductID   uint64    `json:"product_id" gorm:"index;comment:'商品ID'"`
	AttrName    string    `json:"attr_name" gorm:"type:varchar(128);comment:'属性名称'"`
	AttrValue   string    `json:"attr_value" gorm:"type:varchar(500);comment:'属性值'"`
	AttrType    string    `json:"attr_type" gorm:"type:char(1);default:'0';comment:'属性类型 0规格 1参数'"`
	Sort        int       `json:"sort" gorm:"type:int;default:0;comment:'排序'"`
	CreateTime  time.Time `json:"create_time" gorm:"type:datetime;default:CURRENT_TIMESTAMP;comment:'创建时间'"`
	UpdateTime  time.Time `json:"update_time" gorm:"type:datetime;default:CURRENT_TIMESTAMP;comment:'更新时间'"`
}
