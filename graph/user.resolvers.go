package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.35

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql"
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
		Password:  *inputUser.Password,
		CreatedAt: time.Now(),
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
	if inputUser.Password != nil {
		user.Password = *inputUser.Password
	}
	if inputUser.Biography != nil {
		user.Biography = inputUser.Biography
	}
	if inputUser.Location != nil {
		user.Location = inputUser.Location
	}
	if inputUser.Website != nil {
		user.Website = inputUser.Website
	}

	return user, r.DB.Save(&user).Error
}

// UpdateUserProfile is the resolver for the updateUserProfile field.
func (r *mutationResolver) UpdateUserProfile(ctx context.Context, photo graphql.Upload) (string, error) {
	var user *model.User

	userId := ctx.Value("UserID").(string)

	err := r.DB.First(&user, "id = ?", userId).Error

	if err != nil {
		return "", err
	}

	if user.Profile != nil {
		os.Remove("public/" + *user.Profile)
	}

	filename := uuid.NewString() + photo.Filename
	out, err := os.Create("public/images/" + filename)
	if err != nil {
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, photo.File)
	if err != nil {
		return "", err
	}

	src := "images/" + filename
	user.Profile = &src

	return src, r.DB.Save(&user).Error
}

// UpdateUserBackground is the resolver for the updateUserBackground field.
func (r *mutationResolver) UpdateUserBackground(ctx context.Context, photo graphql.Upload) (string, error) {
	var user *model.User

	userId := ctx.Value("UserID").(string)

	err := r.DB.First(&user, "id = ?", userId).Error

	if err != nil {
		return "", err
	}

	if user.Background != nil {
		os.Remove("public/" + *user.Background)
	}

	filename := uuid.NewString() + photo.Filename
	out, err := os.Create("public/images/" + filename)
	if err != nil {
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, photo.File)
	if err != nil {
		return "", err
	}

	src := "images/" + filename
	user.Background = &src

	return src, r.DB.Save(&user).Error
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
func (r *mutationResolver) AuthenticateUser(ctx context.Context, loginUser model.LoginUser) (*model.UserAuth, error) {
	var user *model.User

	err := r.DB.Where("email = ? AND password = ?", loginUser.Email, loginUser.Password).First(&user).Error

	if err != nil {
		return nil, err
	}

	ctx = context.WithValue(ctx, "userID", user.ID)

	tkn, err := helper.CreateJWT(user.ID)

	userAuth := &model.UserAuth{
		User:  user,
		Token: tkn,
	}

	return userAuth, nil
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

// GetUserFromUsername is the resolver for the getUserFromUsername field.
func (r *queryResolver) GetUserFromUsername(ctx context.Context, username string) (*model.User, error) {
	var user *model.User

	return user, r.DB.First(&user, "username = ?", username).Error
}

// GetAllUsers is the resolver for the getAllUsers field.
func (r *queryResolver) GetAllUsers(ctx context.Context) ([]*model.User, error) {
	var users []*model.User
	return users, r.DB.Find(&users).Error
}

// GetUserAuth is the resolver for the getUserAuth field.
func (r *queryResolver) GetUserAuth(ctx context.Context) (*model.User, error) {
	var user *model.User

	userId := ctx.Value("UserID").(string)

	err := r.DB.First(&user, "id = ?", userId).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
