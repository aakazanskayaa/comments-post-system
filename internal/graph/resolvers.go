package graph

// THIS CODE WILL BE UPDATED WITH SCHEMA CHANGES. PREVIOUS IMPLEMENTATION FOR SCHEMA CHANGES WILL BE KEPT IN THE COMMENT SECTION. IMPLEMENTATION FOR UNCHANGED SCHEMA WILL BE KEPT.

import (
	"context"
	"time"

	model1 "github.com/aakazanskayaa/comments-post-system/internal/graph/model"
	"github.com/aakazanskayaa/comments-post-system/internal/model"
)

type Resolver struct{}

// Mutation: CreatePost - создаёт новый пост
func (r *mutationResolver) CreatePost(ctx context.Context, title string, content string, author string, commentsAllowed bool) (*model1.Post, error) {
	newPost := &model.Post{
		ID:              "2", // В реальной системе это должно быть сгенерировано, например, в базе данных
		Title:           title,
		Content:         content,
		Author:          author,
		CommentsAllowed: commentsAllowed,
		CreatedAt:       time.Now().Format(time.RFC3339),
	}
	return newPost, nil
}

// Mutation: AddComment - добавляет новый комментарий
func (r *mutationResolver) AddComment(ctx context.Context, postID string, parentID *string, author string, content string) (*model1.Comment, error) {
	newComment := &model.Comment{
		ID:        "1", // В реальной системе это должно быть сгенерировано
		PostID:    postID,
		ParentID:  parentID,
		Author:    author,
		Content:   content,
		CreatedAt: time.Now().Format(time.RFC3339),
	}
	return newComment, nil
}

// Query: Posts - возвращает список постов
func (r *queryResolver) Posts(ctx context.Context) ([]*model1.Post, error) {
	return []*model.Post{
		{
			ID:              "1",
			Title:           "First Post",
			Content:         "This is the first post content.",
			Author:          "Anna",
			CommentsAllowed: true,
			CreatedAt:       time.Now().Format(time.RFC3339),
		},
	}, nil
}

// Query: Post - возвращает пост по ID
func (r *queryResolver) Post(ctx context.Context, id string) (*model1.Post, error) {
	return &model.Post{
		ID:              id,
		Title:           "Example Post",
		Content:         "This is an example post.",
		Author:          "Anna",
		CommentsAllowed: true,
		CreatedAt:       time.Now().Format(time.RFC3339),
	}, nil
}

// Subscription: CommentAdded - подписка на добавление комментария
func (r *subscriptionResolver) CommentAdded(ctx context.Context, postID string) (<-chan *model1.Comment, error) {
	commentChannel := make(chan *model.Comment)

	// Пример: отправка тестового комментария через канал
	go func() {
		defer close(commentChannel)
		commentChannel <- &model.Comment{
			ID:        "1",
			PostID:    postID,
			ParentID:  nil,
			Author:    "John",
			Content:   "This is a new comment.",
			CreatedAt: time.Now().Format(time.RFC3339),
		}
	}()

	return commentChannel, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Subscription returns SubscriptionResolver implementation.
func (r *Resolver) Subscription() SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
/*
	type Resolver struct{}
*/
