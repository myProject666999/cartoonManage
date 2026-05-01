package models

import "time"

type Order struct {
	ID              uint64    `json:"id" gorm:"primary_key;auto_increment"`
	OrderNo         string    `json:"order_no" gorm:"type:varchar(64);unique_index;not null;comment:'订单号'"`
	MallUserID      uint64    `json:"mall_user_id" gorm:"index;comment:'商城用户ID'"`
	UserName        string    `json:"user_name" gorm:"type:varchar(64);comment:'用户名'"`
	TotalAmount     float64   `json:"total_amount" gorm:"type:decimal(10,2);comment:'订单总金额'"`
	PayAmount       float64   `json:"pay_amount" gorm:"type:decimal(10,2);comment:'实付金额'"`
	FreightAmount   float64   `json:"freight_amount" gorm:"type:decimal(10,2);default:0;comment:'运费金额'"`
	DiscountAmount  float64   `json:"discount_amount" gorm:"type:decimal(10,2);default:0;comment:'优惠金额'"`
	PayType         string    `json:"pay_type" gorm:"type:char(1);comment:'支付方式 1微信 2支付宝 3其他'"`
	PayTime         time.Time `json:"pay_time" gorm:"type:datetime;comment:'支付时间'"`
	OrderStatus     string    `json:"order_status" gorm:"type:char(1);default:'0';comment:'订单状态 0待付款 1待发货 2已发货 3已收货 4已完成 5已取消'"`
	ReceiverName    string    `json:"receiver_name" gorm:"type:varchar(64);comment:'收货人姓名'"`
	ReceiverPhone   string    `json:"receiver_phone" gorm:"type:varchar(32);comment:'收货人电话'"`
	ReceiverProvince string   `json:"receiver_province" gorm:"type:varchar(64);comment:'收货人省份'"`
	ReceiverCity    string    `json:"receiver_city" gorm:"type:varchar(64);comment:'收货人城市'"`
	ReceiverDistrict string   `json:"receiver_district" gorm:"type:varchar(64);comment:'收货人区县'"`
	ReceiverAddress string    `json:"receiver_address" gorm:"type:varchar(255);comment:'收货人详细地址'"`
	DeliveryTime    time.Time `json:"delivery_time" gorm:"type:datetime;comment:'发货时间'"`
	ReceiveTime     time.Time `json:"receive_time" gorm:"type:datetime;comment:'收货时间'"`
	OrderRemark     string    `json:"order_remark" gorm:"type:varchar(500);comment:'订单备注'"`
	CreateTime      time.Time `json:"create_time" gorm:"type:datetime;default:CURRENT_TIMESTAMP;comment:'创建时间'"`
	UpdateTime      time.Time `json:"update_time" gorm:"type:datetime;default:CURRENT_TIMESTAMP;comment:'更新时间'"`
	Deleted         int       `json:"-" gorm:"type:tinyint;default:0;comment:'删除标志 0存在 1删除'"`
	OrderItems      []OrderItem `json:"order_items" gorm:"-"`
}

type OrderItem struct {
	ID          uint64    `json:"id" gorm:"primary_key;auto_increment"`
	OrderID     uint64    `json:"order_id" gorm:"index;comment:'订单ID'"`
	OrderNo     string    `json:"order_no" gorm:"type:varchar(64);index;comment:'订单号'"`
	ProductID   uint64    `json:"product_id" gorm:"index;comment:'商品ID'"`
	ProductName string    `json:"product_name" gorm:"type:varchar(255);comment:'商品名称'"`
	ProductCode string    `json:"product_code" gorm:"type:varchar(64);comment:'商品编码'"`
	ProductImage string   `json:"product_image" gorm:"type:varchar(255);comment:'商品图片'"`
	ProductPrice float64  `json:"product_price" gorm:"type:decimal(10,2);comment:'商品单价'"`
	Quantity    int       `json:"quantity" gorm:"type:int;comment:'购买数量'"`
	TotalPrice  float64   `json:"total_price" gorm:"type:decimal(10,2);comment:'商品总价'"`
	SkuID       uint64    `json:"sku_id" gorm:"comment:'SKU ID'"`
	SkuSpecs    string    `json:"sku_specs" gorm:"type:varchar(500);comment:'SKU规格 JSON格式'"`
	CreateTime  time.Time `json:"create_time" gorm:"type:datetime;default:CURRENT_TIMESTAMP;comment:'创建时间'"`
	UpdateTime  time.Time `json:"update_time" gorm:"type:datetime;default:CURRENT_TIMESTAMP;comment:'更新时间'"`
}
