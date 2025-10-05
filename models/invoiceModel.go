package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ถ้าเป็น * เช่น *string คือการที่เป็น nil ได้เเละ ถ้าเป้น string คือเป้น nil ไม่ได้เเต่เป็น "" (ค่าว่างได้)
type Invoice struct {
	ID               primitive.ObjectID `bson:"_id" json:"_id"`
	Invoice_id       string             `bson:"invoice_id" json:"invoice_id"`
	Order_id         string             `bson:"order_id" json:"order_id"`
	Payment_method   *string            `bson:"payment_method" json:"payment_method" validate:"eq=CARD|eq=CASH|eq="`
	Payment_status   *string            `bson:"payment_status" json:"payment_status" validate:"required, eq=PENDING|eq=PAID"`
	Payment_due_date time.Time          `bson:"payment_due_date" json:"payment_due_date"`
	Created_at       time.Time          `bson:"created_at" json:"created_at"`
	Updated_at       time.Time          `bson:"updated_at" json:"updated_at"`
}
