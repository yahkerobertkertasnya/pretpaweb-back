package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.35

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/yahkerobertkertasnya/preweb/graph/model"
	"github.com/yahkerobertkertasnya/preweb/helper"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, inputUser model.NewUser) (*model.User, error) {
	user := &model.User{
		ID:        uuid.NewString(),
		Name:      inputUser.Name,
		Email:     inputUser.Email,
		Username:  inputUser.Username,
		Password:  inputUser.Password,
		Dob:       nil,
		Content:   nil,
		Profile:   nil,
		CreatedAt: nil,
		Followers: nil,
		Following: nil,
	}

	return user, r.DB.Save(user).Error
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, id string, inputUser model.NewUser) (*model.User, error) {
	var user *model.User

	if err := r.DB.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}

	user.Name = inputUser.Name
	user.Email = inputUser.Email
	user.Username = inputUser.Username
	user.Password = inputUser.Password

	return user, r.DB.Save(&user).Error
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (*model.User, error) {
	var user *model.User

	if err := r.DB.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return user, r.DB.Delete(&user).Error
}

// AuthenticateUser is the resolver for the authenticateUser field.
func (r *mutationResolver) AuthenticateUser(ctx context.Context, loginUser model.LoginUser) (string, error) {
	var user *model.User

	err := r.DB.Where("email = ? AND password = ?", loginUser.Email, loginUser.Password).First(&user).Error

	if err != nil {
		return "", err
	}

	ctx = context.WithValue(ctx, "user", user)

	return helper.CreateJWT(user.ID)
}

// FollowUser is the resolver for the followUser field.
func (r *mutationResolver) FollowUser(ctx context.Context, inputFollow model.NewFollow) (*model.Follow, error) {
	panic(fmt.Errorf("not implemented: FollowUser - followUser"))
}

// GetUser is the resolver for the getUser field.
func (r *queryResolver) GetUser(ctx context.Context, id string) (*model.User, error) {
	var user *model.User

	return user, r.DB.First(&user, "id = ?", id).Error
}

// GetAllUsers is the resolver for the getAllUsers field.
func (r *queryResolver) GetAllUsers(ctx context.Context) ([]*model.User, error) {
	var users []*model.User
	return users, r.DB.Find(&users).Error
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
