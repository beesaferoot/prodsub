package db

import (
	"time"

	"gorm.io/datatypes"

	uuid "github.com/satori/go.uuid"
)

type Product struct {
	Id               uuid.UUID   `gorm:"primary_key;type:uuid;column:id;default:uuid_generate_v4()"`
	Name             string      `gorm:"not null"`
	Type             ProductType `gorm:"not null"`
	Description      string      `gorm:"not null"`
	Price            float64
	CreatedAt        time.Time
	UpdatedAt        time.Time
	ProductAttribute datatypes.JSON `gorm:"not null"`
}

type ProductType uint8

const (
	PhysicalProduct ProductType = iota
	DigitalProduct
	SubscriptionProduct
)

type Subscription struct {
	Id        uuid.UUID `gorm:"primary_key;type:uuid;column:id;default:uuid_generate_v4()"`
	ProductId uuid.UUID `gorm:"not null;type:uuid"`
	Product   Product   `gorm:"foreignKey:ProductId;references:Id;constraint:OnDelete:CASCADE"`
	PlanName  string    `gorm:"not null"`
	Duration  int64     `gorm:"not null"`
	Price     float64   `gorm:"not null"`
}
