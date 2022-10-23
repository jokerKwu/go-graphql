package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"main/common/db"
	"main/graph/generated"
	"main/graph/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	userID, _ := primitive.ObjectIDFromHex(id)
	findData := bson.D{{"_id", userID}}
	var user *model.User
	err := db.UserCollection.FindOne(ctx, findData).Decode(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Users is the resolver for the Users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	cur, err := db.UserCollection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	users := make([]*model.User, 0, cur.RemainingBatchLength())
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var curUser *model.User
		err := cur.Decode(&curUser)
		if err != nil {
			fmt.Println(err)
		}
		users = append(users, curUser)
	}
	return users, nil
}

// AddressBook is the resolver for the addressBook field.
func (r *queryResolver) AddressBook(ctx context.Context, id string) (*model.AddressBook, error) {
	aID, _ := primitive.ObjectIDFromHex(id)
	findData := bson.D{{"_id", aID}}
	var addressBook *model.AddressBook
	err := db.AddressBookCollection.FindOne(ctx, findData).Decode(&addressBook)
	if err != nil {
		return nil, err
	}
	return addressBook, nil
}

// AddressBooks is the resolver for the addressBooks field.
func (r *queryResolver) AddressBooks(ctx context.Context) ([]*model.AddressBook, error) {
	cur, err := db.AddressBookCollection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	addressBooks := make([]*model.AddressBook, 0, cur.RemainingBatchLength())
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var curAddressBook *model.AddressBook
		err := cur.Decode(&curAddressBook)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(curAddressBook)
		addressBooks = append(addressBooks, curAddressBook)
	}
	return addressBooks, nil
}

// Products is the resolver for the products field.
func (r *queryResolver) Products(ctx context.Context) ([]*model.ProductInfo, error) {
	cur, err := db.ProductCollection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	products := make([]*model.ProductInfo, 0, cur.RemainingBatchLength())
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var curProduct *model.ProductInfo
		err := cur.Decode(&curProduct)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(curProduct)
		products = append(products, curProduct)
	}
	return products, nil
}

// Product is the resolver for the product field.
func (r *queryResolver) Product(ctx context.Context, id string) (*model.ProductInfo, error) {
	productID, _ := primitive.ObjectIDFromHex(id)
	findData := bson.D{{"_id", productID}}
	var productDTO *model.ProductInfo
	err := db.ProductCollection.FindOne(ctx, findData).Decode(&productDTO)
	if err != nil {
		return nil, err
	}
	return productDTO, nil
}

// Service is the resolver for the service field.
func (r *queryResolver) Service(ctx context.Context, id string) (*model.Service, error) {
	sID, _ := primitive.ObjectIDFromHex(id)
	findData := bson.D{{"_id", sID}}
	var serviceDTO *model.Service
	err := db.ServiceCollection.FindOne(ctx, findData).Decode(&serviceDTO)
	if err != nil {
		return nil, err
	}
	return serviceDTO, nil
}

// Services is the resolver for the services field.xk
func (r *queryResolver) Services(ctx context.Context) ([]*model.Service, error) {
	panic(fmt.Errorf("not implemented: Services - services"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
