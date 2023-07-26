package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.35

import (
	"context"

	"github.com/google/uuid"
	"github.com/yahkerobertkertasnya/preweb/graph/model"
)

// CreateTweet is the resolver for the createTweet field.
func (r *mutationResolver) CreateTweet(ctx context.Context, inputTweet model.NewTweet) (*model.Tweet, error) {
	var user *model.User

	err := r.DB.First(&user, "id = ?", inputTweet.UserID).Error

	if err != nil {
		return nil, err
	}

	tweet := &model.Tweet{
		ID:      uuid.NewString(),
		UserID:  inputTweet.UserID,
		User:    user,
		Title:   inputTweet.Title,
		Content: inputTweet.Content,
	}

	return tweet, r.DB.Save(tweet).Error
}

// GetUserTweets is the resolver for the getUserTweets field.
func (r *queryResolver) GetUserTweets(ctx context.Context, id string) ([]*model.Tweet, error) {
	var tweets []*model.Tweet

	return tweets, r.DB.Find(&tweets, "user_id = ?", id).Preload("User").Find(&tweets).Error
}

// GetAllTweets is the resolver for the getAllTweets field.
func (r *queryResolver) GetAllTweets(ctx context.Context) ([]*model.Tweet, error) {
	var tweets []*model.Tweet

	return tweets, r.DB.Find(&tweets).Preload("User").Find(&tweets).Error
}
