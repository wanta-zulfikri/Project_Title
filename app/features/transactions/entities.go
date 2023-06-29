package transactions

import (

	"github.com/midtrans/midtrans-go"
	"gorm.io/gorm"
)

type (
	ReqCharge struct {
		PaymentType string
		Invoice     string
		Total       int
		ItemDetails *[]midtrans.ItemDetails 
		CustomerDetails *midtrans.CustomerDetails
	} 

	Carts struct { 
		UserID        uint           `gorm:"not null"` 
		BookID        uint           `gorm:"not null"` 
		DeletedAt     gorm.DeletedAt `gorm:"not null"`	
	}

	BillingBook struct {
		ID         uint  `gorm:"primaryKey;not null;autoIncrement"`
		Name       string 
		Book       string
		Email      string 
		DeletedAt  gorm.DeletedAt `gorm:"index"` 
		Date       string   `gorm:"type:timestamp;not null"` 
		Total      int 
	} 
    Transaction struct {
		Invoice      string  `gorm:"primaryKey;not null;type:varchar(20)" json:"invoice,omitempty"`
		UserID       uint    `gorm:"not null"` 
		BookID       uint    `gorm:"not null"`
		Total        int     `gorm:"not null"` 
		PaymentCode  string  `gorm:"not null"` 
		PaymentMethod string `gorm:"not null"` 
		Status        string `gorm:"not null"`
		TransactionItems []TransactionItems   
	}
	TransactionItems struct {
		TransactionInvoice   string 
		ItemName             string 
		ItemPrice            string 
	}

	ReqCheckout struct {
		BookID       int     `json:"book_id" validate:"required"` 
		Type         string  `json:"type" validate:"required"` 
		PaymentMethod string `json:"payment_method" validate:"required"`
	} 
	ResTransaction struct {
		Invoice     string     `json:"invoice"`
		PaymentMethod  string  `json:"payment_method"`
		Total       int        `json:"total"`
		PaymentCode string     `json:"payment_code"` 
		Expiredate  string     `json:"expire_date"`
	} 
	ResGetAllTransaction struct {
		BookTitle      string    `json:"book_title"` 
		BookImage      string    `json:"book_image"` 
		BookId         string    `json:"book_id"`
	} 
	ResDetailRegisCart struct {
		ItemName    string    `json:"item_name"` 
		ItemPrice   string    `json:"item_price"` 
		Type        string    `json:"type"` 
		Total       string    `json:"total"`
	}
	ResDetailPayment struct {
		Name string    `json:"name"` 
		Price  int     `json:"price"`
	} 
	ResDetailHerRegisCart struct {
		OneTime []ResDetailPayment  `json:"one_time"` 
		Interval []ResDetailPayment `json:"interval"` 
		Type     string             `json:"type"`
		Total    string             `json:"total"`
	} 
	ResDetailTransaction struct {
		Invoice       string   `json:"invoice"` 
		PaymentMethod string   `json:"payment_method"` 
		Total         string   `json:"total"` 
		PaymentCode   string   `json:"payment_code"`
		Expire        string   `json:"expire"`
	}
) 
