package graph

import (
	"context"
	"time"

	"github.com/aakazanskayaa/comments-post-system/db"
	"github.com/aakazanskayaa/comments-post-system/internal/graph/model"
	"github.com/google/uuid"
)

// Resolver - главный резолвер
type Resolver struct{}

// Query Resolver (Запросы)
type queryResolver struct{ *Resolver }

// Mutation Resolver (Мутации)
type mutationResolver struct{ *Resolver }

// Subscription Resolver (Подписки)
type subscriptionResolver struct{ *Resolver }

// ✅ Получение всех постов
func (r *queryResolver) Posts(ctx context.Context) ([]*model.Post, error) {
	return db.DB.GetAllPosts()
}

// ✅ Получение поста по ID
func (r *queryResolver) Post(ctx context.Context, id string) (*model.Post, error) {
	return db.DB.GetPostByID(id)
}

// ✅ Создание нового поста
func (r *mutationResolver) CreatePost(ctx context.Context, title, content, author string, commentsAllowed bool) (*model.Post, error) {
	post := &model.Post{
		ID:              uuid.New().String(),
		Title:           title,
		Content:         content,
		Author:          author,
		CommentsAllowed: commentsAllowed,
		CreatedAt:       time.Now().Format(time.RFC3339),
	}

	err := db.DB.CreatePost(post)
	if err != nil {
		return nil, err
	}
	return post, nil
}

// ✅ Добавление комментария
func (r *mutationResolver) AddComment(ctx context.Context, postID string, parentID *string, author, content string) (*model.Comment, error) {
	post, err := db.DB.GetPostByID(postID)
	if err != nil || post == nil {
		return nil, err
	}

	if !post.CommentsAllowed {
		return nil, err
	}

	comment := &model.Comment{
		ID:        uuid.New().String(),
		PostID:    postID,
		ParentID:  parentID,
		Author:    author,
		Content:   content,
		CreatedAt: time.Now().Format(time.RFC3339),
	}

	err = db.DB.CreateComment(comment)
	if err != nil {
		return nil, err
	}

	return comment, nil
}

// ✅ Получение комментариев к посту с поддержкой пагинации
func (r *queryResolver) Comments(ctx context.Context, postID string, limit, offset int) ([]*model.Comment, error) {
	return db.DB.GetCommentsByPostID(postID, limit, offset)
}

// ✅ Поддержка подписки на новые комментарии (GraphQL Subscriptions)
func (r *subscriptionResolver) CommentAdded(ctx context.Context, postID string) (<-chan *model.Comment, error) {
	commentChan := make(chan *model.Comment, 1)

	go func() {
		for {
			time.Sleep(2 * time.Second)
			comment := &model.Comment{
				ID:        uuid.New().String(),
				PostID:    postID,
				Author:    "Auto-generated",
				Content:   "This is a generated comment!",
				CreatedAt: time.Now().Format(time.RFC3339),
			}
			commentChan <- comment
		}
	}()

	return commentChan, nil
}

// ✅ Подключение к `generated.go`
func (r *Resolver) Mutation() MutationResolver         { return &mutationResolver{r} }
func (r *Resolver) Query() QueryResolver               { return &queryResolver{r} }
func (r *Resolver) Subscription() SubscriptionResolver { return &subscriptionResolver{r} }
