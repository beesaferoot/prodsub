package service_test

import (
	"context"
	"errors"
	"testing"

	"github.com/prodsub/mocks"
	pb "github.com/prodsub/pb/gen"
	"github.com/prodsub/pkg/db"
	"github.com/prodsub/pkg/service"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/mock"
	"github.com/tj/assert"
)

func TestService_CreateProduct(t *testing.T) {

	t.Run("CreateProductSuccess", func(t *testing.T) {
		productIdString := "d28aa4b9-ec42-4018-b221-0a6024d7aa57"
		productId, _ := uuid.FromString(productIdString)

		productRepo := &mocks.ProductRepo{}
		subscrRepo := &mocks.SubscriptionRepo{}

		productRepo.On("Create", mock.Anything).Return(&db.Product{
			Name:        "Shoe",
			Description: "Custom shoe",
			Price:       200.0,
			Id:          productId,
		}, nil).Once()

		s := service.NewProductSubscriptionService(productRepo, subscrRepo)

		res, err := s.CreateProduct(context.TODO(), &pb.ProductCreateRequest{
			Name:        "Shoe",
			Description: "Custom shoe",
			Price:       200.0,
			ProductType: pb.ProductType_DIGITAL,
			ProductAttribute: &pb.ProductAttribute{
				FileSize:     88,
				DownloadLink: "http://download",
			},
		})

		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, productId.String(), res.Product.Id)
	})

	t.Run("CreateProductFailed", func(t *testing.T) {
		productRepo := &mocks.ProductRepo{}
		subscrRepo := &mocks.SubscriptionRepo{}

		productRepo.On("Create", mock.Anything).Return(&db.Product{
			Name:        "Shoe",
			Description: "Custom shoe",
			Price:       200.0,
		}, errors.New("Failed to create product record")).Once()

		s := service.NewProductSubscriptionService(productRepo, subscrRepo)
		res, err := s.CreateProduct(context.TODO(), &pb.ProductCreateRequest{
			Name:        "Shoe",
			Description: "Custom shoe",
			Price:       200.0,
			ProductType: pb.ProductType_PHYSICAL,
			ProductAttribute: &pb.ProductAttribute{
				Weight:     8.0,
				Dimensions: "14x8",
			},
		})
		assert.Error(t, err)
		assert.Nil(t, res)
		productRepo.AssertExpectations(t)
	})
}

func TestService_GetProduct(t *testing.T) {

	t.Run("GetProductSuccess", func(t *testing.T) {
		productRepo := &mocks.ProductRepo{}
		subscrRepo := &mocks.SubscriptionRepo{}

		productIdString := "d28aa4b9-ec42-4018-b221-0a6024d7aa57"
		productId, _ := uuid.FromString(productIdString)
		productRepo.On("Get", mock.Anything).Return(&db.Product{
			Name:        "Shoe",
			Description: "Custom shoe",
			Price:       200.0,
			Id:          productId,
		}, nil).Once()

		s := service.NewProductSubscriptionService(productRepo, subscrRepo)
		res, err := s.GetProduct(context.TODO(), &pb.ProductGetRequest{
			ProductId: productIdString,
		})

		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, productIdString, res.Product.Id)
	})

	t.Run("GetProductFailed", func(t *testing.T) {
		productRepo := &mocks.ProductRepo{}
		subscrRepo := &mocks.SubscriptionRepo{}

		productIdString := "d28aa4b9-ec42-4018-b221-0a6024d7aa57"

		productRepo.On("Get", mock.Anything).Return(nil, errors.New("Failed to retrieve product with id")).Once()

		s := service.NewProductSubscriptionService(productRepo, subscrRepo)
		res, err := s.GetProduct(context.TODO(), &pb.ProductGetRequest{
			ProductId: productIdString,
		})

		assert.Error(t, err)
		assert.Nil(t, res)
	})

}

func TestService_ListProduct(t *testing.T) {
	t.Run("ListProductAll", func(t *testing.T) {
		productRepo := &mocks.ProductRepo{}
		subscrRepo := &mocks.SubscriptionRepo{}

		productRepo.On("List", mock.Anything).Return([]db.Product{db.Product{
			Name:        "Shoe",
			Description: "Custom shoe",
			Price:       200.0,
		}}, nil).Once()
		s := service.NewProductSubscriptionService(productRepo, subscrRepo)

		res, err := s.ListProduct(context.TODO(), &pb.ProductListRequest{})
		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, 1, len(res.Products))
	})

	t.Run("ListProductWithType", func(t *testing.T) {
		productRepo := &mocks.ProductRepo{}
		subscrRepo := &mocks.SubscriptionRepo{}

		productRepo.On("List", mock.Anything).Return([]db.Product{db.Product{
			Name:        "Shoe",
			Description: "Custom shoe",
			Price:       200.0,
		}}, nil).Once()
		s := service.NewProductSubscriptionService(productRepo, subscrRepo)

		res, err := s.ListProduct(context.TODO(), &pb.ProductListRequest{
			ProductType: pb.ProductType_DIGITAL,
		})
		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, 1, len(res.Products))
	})
}

func TestService_UpdateProduct(t *testing.T) {
	t.Run("UpdateProductSuccess", func(t *testing.T) {
		productRepo := &mocks.ProductRepo{}
		subscrRepo := &mocks.SubscriptionRepo{}

		productIdString := "d28aa4b9-ec42-4018-b221-0a6024d7aa57"
		productId, _ := uuid.FromString(productIdString)

		productRepo.On("Update", mock.Anything, mock.Anything).Return(&db.Product{
			Name:        "Shoe",
			Description: "Custom shoe",
			Price:       200.0,
			Id:          productId,
		}, nil).Once()

		s := service.NewProductSubscriptionService(productRepo, subscrRepo)

		res, err := s.UpdateProduct(context.TODO(), &pb.ProductUpdateRequest{
			ProductId: productIdString,
			Product: &pb.ProductDetails{
				Name:        "New Shoe",
				Description: "New Custom Shoe",
				Price:       39.0,
				Attribute: &pb.ProductAttribute{
					FileSize: 899,
				},
			},
		})

		assert.NoError(t, err)
		assert.NotNil(t, res)

		assert.Equal(t, productIdString, res.Product.Id)
		productRepo.AssertExpectations(t)
	})
}

func TestService_DeleteProduct(t *testing.T) {
	t.Run("DeleteProductSuccess", func(t *testing.T) {
		productRepo := &mocks.ProductRepo{}
		subscrRepo := &mocks.SubscriptionRepo{}

		productIdString := "d28aa4b9-ec42-4018-b221-0a6024d7aa57"

		productRepo.On("Delete", mock.Anything).Return(nil).Once()

		s := service.NewProductSubscriptionService(productRepo, subscrRepo)

		res, err := s.DeleteProduct(context.TODO(), &pb.ProductDeleteRequest{
			ProductId: productIdString,
		})

		assert.NoError(t, err)
		assert.NotNil(t, res)
		productRepo.AssertExpectations(t)
	})
}

func TestService_CreateSubscription(t *testing.T) {
	t.Run("CreateSubscriptionSuccess", func(t *testing.T) {
		productRepo := &mocks.ProductRepo{}
		subscrRepo := &mocks.SubscriptionRepo{}

		productIdString := "d28aa4b9-ec42-4018-b221-0a6024d7aa57"
		productId, _ := uuid.FromString(productIdString)

		subscrRepo.On("Create", mock.Anything).Return(&db.Subscription{
			Id:        uuid.NewV4(),
			ProductId: productId,
			PlanName:  "Monthly",
			Duration:  30,
		}, nil).Once()

		s := service.NewProductSubscriptionService(productRepo, subscrRepo)

		res, err := s.CreateSubscription(context.TODO(), &pb.SubscriptionCreateRequest{
			ProductId: productIdString,
			PlanName:  "Monthly",
			Duration:  30,
		})
		assert.NoError(t, err)
		assert.NotNil(t, res)
		subscrRepo.AssertExpectations(t)
	})
}

func TestService_GetSubscription(t *testing.T) {
	t.Run("GetSubscriptionSuccess", func(t *testing.T) {
		productRepo := &mocks.ProductRepo{}
		subscrRepo := &mocks.SubscriptionRepo{}

		subId := uuid.NewV4()

		productIdString := "d28aa4b9-ec42-4018-b221-0a6024d7aa57"
		productId, _ := uuid.FromString(productIdString)

		subscrRepo.On("Get", mock.Anything).Return(&db.Subscription{
			Id:        subId,
			ProductId: productId,
			PlanName:  "Monthly",
			Duration:  30,
		}, nil).Once()

		s := service.NewProductSubscriptionService(productRepo, subscrRepo)

		res, err := s.GetSubscription(context.TODO(), &pb.SubscriptionGetRequest{
			SubscriptionId: subId.String(),
		})

		assert.NoError(t, err)
		assert.NotNil(t, res)

		assert.Equal(t, subId.String(), res.Subscription.Id)

		subscrRepo.AssertExpectations(t)
	})
}

func TestService_ListSubscription(t *testing.T) {
	t.Run("ListSubscriptionSuccess", func(t *testing.T) {
		productRepo := &mocks.ProductRepo{}
		subscrRepo := &mocks.SubscriptionRepo{}

		subId := uuid.NewV4()

		productIdString := "d28aa4b9-ec42-4018-b221-0a6024d7aa57"
		productId, _ := uuid.FromString(productIdString)

		subscrRepo.On("List", mock.Anything).Return([]db.Subscription{db.Subscription{
			Id:        subId,
			ProductId: productId,
			PlanName:  "Monthly",
			Duration:  30,
		}}, nil).Once()

		s := service.NewProductSubscriptionService(productRepo, subscrRepo)

		res, err := s.ListSubscription(context.TODO(), &pb.SubscriptionListRequest{
			ProductId: productIdString,
		})

		assert.NoError(t, err)
		assert.NotNil(t, res)

		assert.Equal(t, 1, len(res.Subscriptions))
	})
}

func TestService_UpdateSubscription(t *testing.T) {
	t.Run("UpdateSubscriptionSuccess", func(t *testing.T) {
		productRepo := &mocks.ProductRepo{}
		subscrRepo := &mocks.SubscriptionRepo{}

		subId := uuid.NewV4()

		productIdString := "d28aa4b9-ec42-4018-b221-0a6024d7aa57"
		productId, _ := uuid.FromString(productIdString)

		subscrRepo.On("Update", mock.Anything, mock.Anything).Return(&db.Subscription{
			Id:        subId,
			ProductId: productId,
			PlanName:  "Monthly",
			Duration:  30,
		}, nil).Once()

		s := service.NewProductSubscriptionService(productRepo, subscrRepo)

		res, err := s.UpdateSubscription(context.TODO(), &pb.SubscriptionUpdateRequest{
			SubscriptionId: subId.String(),
			Subscription: &pb.SubscriptionDetail{
				PlanName: "New Monthly Sub",
				Duration: 60,
				Price:    800.0,
			},
		})
		assert.NoError(t, err)
		assert.NotNil(t, res)
		subscrRepo.AssertExpectations(t)
	})
}

func TestService_DeleteSubscription(t *testing.T) {
	t.Run("DeleteSubscriptionSuccess", func(t *testing.T) {
		productRepo := &mocks.ProductRepo{}
		subscrRepo := &mocks.SubscriptionRepo{}

		subId := uuid.NewV4()

		subscrRepo.On("Delete", mock.Anything).Return(nil).Once()
		s := service.NewProductSubscriptionService(productRepo, subscrRepo)

		res, err := s.DeleteSubscription(context.TODO(), &pb.SubscriptionDeleteRequest{
			SubscriptionId: subId.String(),
		})

		assert.NoError(t, err)
		assert.NotNil(t, res)
	})
}
