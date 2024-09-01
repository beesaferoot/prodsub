package db

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type SubscriptionUpdateRequest struct {
	PlanName string  `json:"plan_name,omitempty"`
	Duration int64   `json:"duration,omitempty"`
	Price    float64 `json:"price,omitempty"`
}

type SubscriptionRepo interface {
	Create(*Subscription) (*Subscription, error)
	Get(id uuid.UUID) (*Subscription, error)
	Update(id uuid.UUID, req SubscriptionUpdateRequest) (*Subscription, error)
	Delete(id uuid.UUID) error
	List(productId uuid.UUID) ([]Subscription, error)
}

type subscriptionRepo struct {
	Db *gorm.DB
}

func NewSubscriptionRepo(db *gorm.DB) SubscriptionRepo {
	return &subscriptionRepo{
		Db: db,
	}
}

func (s *subscriptionRepo) Create(subscription *Subscription) (*Subscription, error) {
	result := s.Db.Create(subscription)
	if result.Error != nil {
		return nil, result.Error
	}
	return subscription, nil
}

func (s *subscriptionRepo) Get(id uuid.UUID) (*Subscription, error) {
	var subscription Subscription

	result := s.Db.First(&subscription, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &subscription, nil
}

func (s *subscriptionRepo) Update(id uuid.UUID, req SubscriptionUpdateRequest) (*Subscription, error) {
	var subscription Subscription = Subscription{}

	if req.PlanName != "" {
		subscription.PlanName = req.PlanName
	}

	if req.Duration > 0 {
		subscription.Duration = req.Duration
	}

	if req.Price > 0 {
		subscription.Price = req.Price
	}

	result := s.Db.Model(&subscription).Where("id = ?", id).Updates(subscription)

	if result.Error != nil {
		return nil, result.Error
	}

	subscription.Id = id

	return &subscription, nil
}

func (s *subscriptionRepo) List(productId uuid.UUID) ([]Subscription, error) {
	subscription := []Subscription{}

	result := s.Db.Model(&Subscription{}).Find(&subscription, "product_id = ?", productId)

	if result.Error != nil {
		return subscription, result.Error
	}

	return subscription, nil
}

func (s *subscriptionRepo) Delete(id uuid.UUID) error {
	err := s.Db.Transaction(func(tx *gorm.DB) error {
		result := tx.Delete(&Subscription{}, id)
		if result.Error != nil {
			return result.Error
		}
		return nil
	})
	return err
}
