package models

import (
	"time"

	"gorm.io/gorm"
)

type Address struct {
	gorm.Model
	ID int `gorm:"primaryKey" json:"id"`

	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`

	Addressline1 string `json:"address_line_1"`
	Addressline2 string `json:"address_line_2"`
	Landmark     string `json:"landmark"`
	City         string `json:"city"`
	Country      string `json:"country"`
	Pincode      string `json:"pincode"`

	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type CartItem struct {
	gorm.Model
	ID        int `gorm:"primaryKey" json:"id"`
	ProductID int
	Product   Product `json:"product"`
	Quantity  int     `json:"quantity"`
	CartID    int     `json:"cart_id"`
}

type Cart struct {
	gorm.Model
	ID        int        `gorm:"primaryKey" json:"id"`
	UserID    int        `json:"user_id"`
	Total     int        `json:"total"`
	CartItems []CartItem `json:"cart_items"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

type Category struct {
	gorm.Model
	ID       int    `gorm:"primaryKey" json:"id"`
	Name     string `json:"name"`
	Products []Product
}

// type OrderStatus int

// const (
// 	Pending OrderStatus = iota
// 	Processing
// 	Shipped
// 	Delivered
// 	Cancelled
// )

type OrderStatus string

const (
	Pending    OrderStatus = "Pending"
	Processing OrderStatus = "Processing"
	Shipped    OrderStatus = "Shipped"
	Delivered  OrderStatus = "Delivered"
	Cancelled  OrderStatus = "Cancelled"
	Paid       OrderStatus = "Paid"
)

type OrderItem struct {
	gorm.Model
	ID int `gorm:"primaryKey" json:"id"`

	Product   Product `gorm:"foreignKey:ProductID;" json:"product"`
	ProductID int     `json:"product_id"`
	Quantity  int     `json:"quantity"`
	OrderID   int     `json:"order_id"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// type ProductOrder struct {
// 	ProductID int `json:"product_id"`
// 	OrderID   int `json:"order_id"`
// }

type Order struct {
	gorm.Model
	ID     int  `gorm:"primaryKey" json:"id"`
	UserID int  `json:"user_id"`
	User   User `gorm:"foreignKey:UserID" json:"user"`

	OrderItems []OrderItem `json:"order_items"`
	Address    Address     `gorm:"foreignKey:AddressID"`
	AddressID  int         `json:"address_id"`
	SubTotal   int         `json:"sub_total"`
	Tax        int         `json:"tax"`
	Status     string      `json:"status"`
	Total      int         `json:"total"`
	Products   []Product   `gorm:"many2many:product_orders" json:"products"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
}

type PaymentStatus int

const (
	Success PaymentStatus = iota
	Failed
)

type Payment struct {
	gorm.Model
	ID                int           `gorm:"primaryKey" json:"id"`
	UserID            int           `json:"user_id"`
	OrderID           int           `json:"order_id"`
	RazorpayPaymentID string        `json:"razorpay_payment_id"`
	RazorpaySignature string        `json:"razorpay_signature"`
	RazorpayOrderID   string        `json:"razorpay_order_id"`
	Status            PaymentStatus `json:"status"`
}

type Product struct {
	gorm.Model
	ID          int       `gorm:"primaryKey" json:"id"`
	Name        string    `json:"name"`
	Link        string    `json:"link"`
	Price       int       `json:"price"`
	SalePrice   int       `json:"sale_price"`
	CategoryID  int       `json:"category_id"`
	Category    Category  `json:"category"`
	Tax         int       `json:"tax"`
	Image       string    `json:"image"`
	Description string    `json:"description"`
	Orders      []Order   `gorm:"many2many:product_orders" json:"orders"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}

type Shop struct {
	gorm.Model
	ID      int `gorm:"primaryKey"`
	Name    string
	Address string
}

type User struct {
	gorm.Model
	ID        int     `gorm:"primaryKey" json:"id"`
	Name      string  `json:"name"`
	Username  string  `json:"username"`
	Password  string  `json:"-"`
	RoleID    int     `json:"role_id"`
	FirstName string  `json:"last_name"`
	LastName  string  `json:"firs_name"`
	Phone     string  `json:"phone"`
	Email     string  `json:"email"`
	Gender    string  `json:"gender"`
	City      string  `json:"city"`
	Country   string  `json:"country"`
	Orders    []Order `json:"orders"`
}

type Role struct {
	gorm.Model
	ID   int    `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}
