package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"main/common/db"
	"main/graph/generated"
	"main/graph/model"

	"go.mongodb.org/mongo-driver/bson"
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

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
type userResolver struct{ *Resolver }
