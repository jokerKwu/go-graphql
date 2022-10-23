package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"main/common/db"
	"main/graph/generated"
	"main/graph/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	res, err := db.UserCollection.InsertOne(ctx, input)
	if err != nil {
		return nil, err
	}
	fmt.Println(res.InsertedID)
	result := &model.User{
		ID:    res.InsertedID.(primitive.ObjectID).String(),
		Name:  input.Name,
		Age:   input.Age,
		Phone: input.Phone,
	}
	return result, nil
}

// CreateAddressBook is the resolver for the createAddressBook field.
func (r *mutationResolver) CreateAddressBook(ctx context.Context, input model.NewAddressBook) (*model.AddressBook, error) {
	res, err := db.AddressBookCollection.InsertOne(ctx, input)
	if err != nil {
		return nil, err
	}
	fmt.Println(res.InsertedID)
	result := &model.AddressBook{
		ID: res.InsertedID.(primitive.ObjectID).String(),
		Address: &model.Address{
			Receiver: input.Address.Receiver,
			Address:  input.Address.Address,
			Detail:   input.Address.Detail,
			PinCode:  input.Address.PinCode,
			Tel:      input.Address.Tel,
			Request:  input.Address.Request,
		},
		IsDefault: input.IsDefault,
	}
	return result, nil
}

// CreateProduct is the resolver for the createProduct field.
func (r *mutationResolver) CreateProduct(ctx context.Context, input model.NewProductInfo) (*model.ProductInfo, error) {
	_, err := db.ProductCollection.InsertOne(ctx, input)
	if err != nil {
		return nil, err
	}
	result := &model.ProductInfo{
		Name:        input.Name,
		Description: input.Description,
		PerAmount:   input.PerAmount,
		ImgURL:      input.ImgURL,
	}
	return result, nil
}

// CreateService is the resolver for the createService field.
func (r *mutationResolver) CreateService(ctx context.Context, input model.NewService) (bool, error) {
	_, err := db.ServiceCollection.InsertOne(ctx, input)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
