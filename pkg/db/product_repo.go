package db

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type ProductUpdateRequest struct {
	Name        string  `json:"name,omitempty"`
	Description string  `json:"description,omitempty"`
	Price       float64 `json:"price,omitempty"`
	Attribute   []byte
}

type ProductRepo interface {
	Create(*Product) (*Product, error)
	Get(id uuid.UUID) (*Product, error)
	Update(id uuid.UUID, req ProductUpdateRequest) (*Product, error)
	List(productType string) ([]Product, error)
	Delete(id uuid.UUID) error
}

type productRepo struct {
	Db *gorm.DB
}

func NewProductRepo(db *gorm.DB) ProductRepo {
	return &productRepo{
		Db: db,
	}
}

func (p *productRepo) Create(product *Product) (*Product, error) {
	product.CreatedAt = time.Now()
	result := p.Db.Create(product)
	if result.Error != nil {
		return nil, result.Error
	}
	return product, nil
}

func (p *productRepo) Get(id uuid.UUID) (*Product, error) {
	var product Product

	result := p.Db.First(&product, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &product, nil
}

func (p *productRepo) Update(id uuid.UUID, req ProductUpdateRequest) (*Product, error) {
	var product Product = Product{}

	if req.Description != "" {
		product.Description = req.Description
	}

	if req.Name != "" {
		product.Name = req.Name
	}

	if req.Price > 0 {
		product.Price = req.Price
	}

	if len(req.Attribute) > 0 {
		product.ProductAttribute = datatypes.JSON(req.Attribute)
	}

	product.UpdatedAt = time.Now()

	result := p.Db.Model(&product).Where("id = ?", id).Updates(product)

	if result.Error != nil {
		return nil, result.Error
	}

	return &product, nil
}

func (p *productRepo) List(productType string) ([]Product, error) {
	product := []Product{}

	var result *gorm.DB

	if productType == "" {
		result = p.Db.Find(&product)
	} else {
		result = p.Db.Model(&Product{}).Find(&product, "id = ?", productType)
	}

	if result.Error != nil {
		return product, result.Error
	}

	return product, nil

}

func (p *productRepo) Delete(id uuid.UUID) error {
	err := p.Db.Transaction(func(tx *gorm.DB) error {
		result := tx.Delete(&Product{}, id)
		if result.Error != nil {
			return result.Error
		}
		return nil
	})
	return err
}
