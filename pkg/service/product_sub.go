package service

import (
	"context"
	"encoding/json"

	pb "github.com/prodsub/pb/gen"
	"github.com/prodsub/pkg/db"
	"github.com/prodsub/pkg/util"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/datatypes"
)

type ProductSubscriptionService struct {
	productRepo      db.ProductRepo
	subscriptionRepo db.SubscriptionRepo
	pb.UnimplementedProductServiceServer
	pb.UnimplementedSubscriptionServiceServer
}

func NewProductSubscriptionService(productRepo db.ProductRepo, subscriptionRepo db.SubscriptionRepo) *ProductSubscriptionService {

	return &ProductSubscriptionService{
		productRepo:      productRepo,
		subscriptionRepo: subscriptionRepo,
	}
}

func (p *ProductSubscriptionService) CreateProduct(ctx context.Context, req *pb.ProductCreateRequest) (*pb.ProductCreateResponse, error) {

	attr := &util.ProductAttribute{}

	m := protojson.MarshalOptions{
		AllowPartial: true,
	}
	jData, err := m.Marshal(req.ProductAttribute)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to unmarshal ProductAttribute. Error %s", err.Error())
	}

	err = json.Unmarshal(jData, attr)

	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to unmarshal ProductAttribute. Error %s", err.Error())
	}

	// ensure the right attributes where passed based on the product type
	if attr != nil && req.ProductType == pb.ProductType_PHYSICAL {
		if attr.Weight == 0 || attr.Dimensions == "" {
			return nil, status.Errorf(codes.InvalidArgument, "invalid data, missing weight/dimensions attribute. Error %v", attr)
		}
	}

	if attr != nil && req.ProductType == pb.ProductType_DIGITAL {
		if attr.FileSize == 0 || attr.DownloadLink == "" {
			return nil, status.Errorf(codes.InvalidArgument, "invalid data, missing file_size/download_link attribute. Error %v", attr)
		}
	}

	if attr != nil && req.ProductType == pb.ProductType_SUBSCRIPTION {
		if attr.SubscriptionPeriod == "" || attr.RenewalPrice == 0 {
			return nil, status.Errorf(codes.InvalidArgument, "invalid data, missing subscription_period/renewal_price attribute. Error %v", attr)
		}
	}

	product := &db.Product{
		Name:             req.Name,
		Type:             db.ProductType(req.ProductType),
		Description:      req.Description,
		Price:            req.Price,
		ProductAttribute: datatypes.JSON(jData),
	}

	product, err = p.productRepo.Create(product)

	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.ProductCreateResponse{
		Product: dbProductToPbProduct(product),
	}

	return res, nil

}

func (p *ProductSubscriptionService) GetProduct(ctx context.Context, req *pb.ProductGetRequest) (*pb.ProductGetResponse, error) {
	productId, err := uuid.FromString(req.ProductId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid format of product_id uuid. Error %s", err.Error())
	}
	product, err := p.productRepo.Get(productId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &pb.ProductGetResponse{
		Product: dbProductToPbProduct(product),
	}, nil
}

func (p *ProductSubscriptionService) UpdateProduct(ctx context.Context, req *pb.ProductUpdateRequest) (*pb.ProductUpdateResponse, error) {
	productId, err := uuid.FromString(req.ProductId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid format of product_id uuid. Error %s", err.Error())
	}

	m := protojson.MarshalOptions{
		AllowPartial: true,
	}
	jData, err := m.Marshal(req.Product.Attribute)

	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to unmarshal ProductAttribute. Error %s", err.Error())
	}
	updateArgs := db.ProductUpdateRequest{
		Name:        req.Product.Name,
		Description: req.Product.Description,
		Price:       req.Product.Price,
		Attribute:   jData,
	}

	product, err := p.productRepo.Update(productId, updateArgs)

	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.ProductUpdateResponse{
		Product: dbProductToPbProduct(product),
	}, nil

}

func (p *ProductSubscriptionService) ListProduct(ctx context.Context, req *pb.ProductListRequest) (*pb.ProductListResponse, error) {

	list, err := p.productRepo.List(req.ProductType.String())
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.ProductListResponse{
		Products: dbProductsToPbProducts(list),
	}, nil
}

func (p *ProductSubscriptionService) DeleteProduct(ctx context.Context, req *pb.ProductDeleteRequest) (*pb.ProductDeleteResponse, error) {
	productId, err := uuid.FromString(req.ProductId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid format of product_id uuid. Error %s", err.Error())
	}

	err = p.productRepo.Delete(productId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.ProductDeleteResponse{}, nil
}

func (p *ProductSubscriptionService) CreateSubscription(ctx context.Context, req *pb.SubscriptionCreateRequest) (*pb.SubscriptionCreateResponse, error) {
	productId, err := uuid.FromString(req.ProductId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid format of product_id uuid. Error %s", err.Error())
	}

	subscr := &db.Subscription{
		ProductId: productId,
		PlanName:  req.PlanName,
		Price:     req.Price,
		Duration:  req.Duration,
	}

	subscr, err = p.subscriptionRepo.Create(subscr)

	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.SubscriptionCreateResponse{
		Subscription: dbSubscriptionToPbSubscription(subscr),
	}, nil
}

func (p *ProductSubscriptionService) GetSubscription(ctx context.Context, req *pb.SubscriptionGetRequest) (*pb.SubscriptionGetResponse, error) {
	subId, err := uuid.FromString(req.SubscriptionId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid format of subscription_id uuid. Error %s", err.Error())
	}

	subscr, err := p.subscriptionRepo.Get(subId)

	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.SubscriptionGetResponse{
		Subscription: dbSubscriptionToPbSubscription(subscr),
	}, nil
}

func (p *ProductSubscriptionService) UpdateSubscription(ctx context.Context, req *pb.SubscriptionUpdateRequest) (*pb.SubscriptionUpdateResponse, error) {
	subId, err := uuid.FromString(req.SubscriptionId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid format of subscription_id uuid. Error %s", err.Error())
	}

	updateArgs := db.SubscriptionUpdateRequest{
		PlanName: req.Subscription.PlanName,
		Price:    req.Subscription.Price,
		Duration: req.Subscription.Duration,
	}

	subscr, err := p.subscriptionRepo.Update(subId, updateArgs)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.SubscriptionUpdateResponse{
		Subscription: dbSubscriptionToPbSubscription(subscr),
	}, nil
}

func (p *ProductSubscriptionService) DeleteSubscription(ctx context.Context, req *pb.SubscriptionDeleteRequest) (*pb.SubscriptionDeleteResponse, error) {
	subId, err := uuid.FromString(req.SubscriptionId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid format of subscription_id uuid. Error %s", err.Error())
	}

	err = p.subscriptionRepo.Delete(subId)

	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.SubscriptionDeleteResponse{}, nil
}

func (p *ProductSubscriptionService) ListSubscription(ctx context.Context, req *pb.SubscriptionListRequest) (*pb.SubscriptionListResponse, error) {
	productId, err := uuid.FromString(req.ProductId)

	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid format of product_id uuid. Error %s", err.Error())
	}

	list, err := p.subscriptionRepo.List(productId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.SubscriptionListResponse{
		Subscriptions: dbSubscriptionsToPbSubscriptions(list),
	}, nil
}

func dbProductToPbProduct(product *db.Product) *pb.Product {
	prod := &pb.Product{
		Id:          product.Id.String(),
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		CreateAt:    timestamppb.New(product.CreatedAt),
		UpdateAt:    timestamppb.New(product.UpdatedAt),
		ProductType: pb.ProductType(product.Type),
	}

	attr := &pb.ProductAttribute{}

	m := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: true,
	}
	err := m.Unmarshal([]byte(product.ProductAttribute.String()), attr)

	if err == nil {
		prod.ProductAttribute = attr
	}

	return prod
}

func dbProductsToPbProducts(products []db.Product) []*pb.Product {
	res := []*pb.Product{}

	for _, p := range products {
		res = append(res, dbProductToPbProduct(&p))
	}

	return res
}

func dbSubscriptionToPbSubscription(subscription *db.Subscription) *pb.Subscription {
	subscr := &pb.Subscription{
		Id:        subscription.Id.String(),
		PlanName:  subscription.PlanName,
		Price:     subscription.Price,
		Duration:  subscription.Duration,
		ProductId: subscription.ProductId.String(),
	}

	return subscr
}

func dbSubscriptionsToPbSubscriptions(subscriptions []db.Subscription) []*pb.Subscription {
	res := []*pb.Subscription{}

	for _, s := range subscriptions {
		res = append(res, dbSubscriptionToPbSubscription(&s))
	}

	return res
}
