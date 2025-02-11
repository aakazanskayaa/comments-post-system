package graph_test

import (
	"context"
	"os"
	"testing"

	//"time"

	"github.com/aakazanskayaa/comments-post-system/db"
	"github.com/aakazanskayaa/comments-post-system/internal/graph"

	//"github.com/aakazanskayaa/comments-post-system/internal/graph/model"
	//"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func setupTestDB() {
	db.DB = db.NewMemoryStorage() // или db.NewPostgresStorage(dsn), если тестируешь с PostgreSQL
}

func TestMain(m *testing.M) {
	setupTestDB() // Инициализация хранилища перед запуском тестов
	os.Exit(m.Run())
}

// Тест создания поста
func TestCreatePost(t *testing.T) {
	// Создаем in-memory хранилище перед тестом
	db.DB = db.NewMemoryStorage()

	resolver := &graph.Resolver{} // Используем указатель, чтобы передавать структуру по ссылке
	ctx := context.Background()

	newPost, err := resolver.Mutation().CreatePost(ctx, "Test Title", "Test Content", "Test Author", true)

	assert.NoError(t, err, "Ошибка при создании поста")
	assert.NotNil(t, newPost, "Созданный пост не должен быть nil")
	assert.Equal(t, "Test Title", newPost.Title)
	assert.Equal(t, "Test Content", newPost.Content)
	assert.Equal(t, "Test Author", newPost.Author)
	assert.Equal(t, true, newPost.CommentsAllowed)
}

// Тест добавления комментария
func TestAddComment(t *testing.T) {
	db.DB = db.NewMemoryStorage() // Инициализируем хранилище

	resolver := &graph.Resolver{}
	ctx := context.Background()

	// Создаем тестовый пост перед добавлением комментария
	newPost, err := resolver.Mutation().CreatePost(ctx, "Test Title", "Test Content", "Test Author", true)
	assert.NoError(t, err, "Ошибка при создании поста")
	assert.NotNil(t, newPost, "Созданный пост не должен быть nil")

	// Теперь создаем комментарий к только что созданному посту
	newComment, err := resolver.Mutation().AddComment(ctx, newPost.ID, nil, "Comment Author", "Test Comment")

	assert.NoError(t, err, "Ошибка при добавлении комментария")
	assert.NotNil(t, newComment, "Созданный комментарий не должен быть nil")
	assert.Equal(t, newPost.ID, newComment.PostID, "Комментарий должен быть привязан к правильному посту")
	assert.Equal(t, "Comment Author", newComment.Author)
	assert.Equal(t, "Test Comment", newComment.Content)
}

// Тест получения поста с комментариями
func TestGetPostWithComments(t *testing.T) {
	db.DB = db.NewMemoryStorage() // Инициализируем хранилище

	resolver := &graph.Resolver{}
	ctx := context.Background()

	// Создаем тестовый пост
	newPost, err := resolver.Mutation().CreatePost(ctx, "Test Title", "Test Content", "Test Author", true)
	assert.NoError(t, err, "Ошибка при создании поста")
	assert.NotNil(t, newPost, "Созданный пост не должен быть nil")

	// Добавляем комментарий к этому посту
	newComment, err := resolver.Mutation().AddComment(ctx, newPost.ID, nil, "Comment Author", "Test Comment")
	assert.NoError(t, err, "Ошибка при добавлении комментария")
	assert.NotNil(t, newComment, "Созданный комментарий не должен быть nil")

	// Получаем комментарии к этому посту
	comments, err := resolver.Query().Comments(ctx, newPost.ID, 10, 0)
	assert.NoError(t, err, "Ошибка при получении комментариев")
	assert.NotNil(t, comments, "Список комментариев не должен быть nil")
	assert.Equal(t, 1, len(comments), "Должен быть 1 комментарий")
	assert.Equal(t, newComment.ID, comments[0].ID, "ID комментария должен совпадать")
}

// Тест создания поста без комментариев
func TestCreatePostWithoutComments(t *testing.T) {
	resolver := &graph.Resolver{}
	ctx := context.Background()

	newPost, err := resolver.Mutation().CreatePost(ctx, "No Comments Post", "No Comments Content", "Test Author", false)
	assert.NoError(t, err, "Ошибка при создании поста")
	assert.NotNil(t, newPost, "Созданный пост не должен быть nil")
	assert.Equal(t, false, newPost.CommentsAllowed)
}

// Тест добавления комментария к посту, где они запрещены
func TestAddCommentToPostWithDisabledComments(t *testing.T) {
	db.DB = db.NewMemoryStorage() // Инициализируем хранилище

	resolver := &graph.Resolver{}
	ctx := context.Background()

	// Создаем пост, в котором запрещены комментарии
	newPost, err := resolver.Mutation().CreatePost(ctx, "No Comments Post", "Content", "Test Author", false)
	assert.NoError(t, err, "Ошибка при создании поста")
	assert.NotNil(t, newPost, "Созданный пост не должен быть nil")

	// Пытаемся добавить комментарий и ожидаем ошибку
	newComment, err := resolver.Mutation().AddComment(ctx, newPost.ID, nil, "Comment Author", "Test Comment")
	assert.Error(t, err, "Ожидалась ошибка при попытке добавить комментарий к посту с отключенными комментариями")
	assert.Nil(t, newComment, "Комментарий не должен быть создан")
	assert.Equal(t, "комментарии к этому посту запрещены", err.Error(), "Текст ошибки должен совпадать")
}

/* func TestFullProcess(t *testing.T) {
	db.DB = db.NewMemoryStorage() // Инициализируем хранилище

	resolver := &graph.Resolver{}
	ctx := context.Background()

	// 1️⃣ Создаем пост
	newPost, err := resolver.Mutation().CreatePost(ctx, "Integration Test Post", "Test Content", "Test Author", true)
	assert.NoError(t, err, "Ошибка при создании поста")
	assert.NotNil(t, newPost, "Созданный пост не должен быть nil")

	// 2️⃣ Добавляем комментарий к посту
	newComment, err := resolver.Mutation().AddComment(ctx, newPost.ID, nil, "Test Commenter", "This is a test comment.")
	assert.NoError(t, err, "Ошибка при добавлении комментария")
	assert.NotNil(t, newComment, "Созданный комментарий не должен быть nil")

	// 3️⃣ Ждем 100 мс (чтобы эмулировать задержку обработки в памяти)
	time.Sleep(100 * time.Millisecond)

	// 4️⃣ Получаем пост с комментариями
	retrievedPost, err := resolver.Query().Post(ctx, newPost.ID)
	assert.NoError(t, err, "Ошибка при получении поста")
	assert.NotNil(t, retrievedPost, "Полученный пост не должен быть nil")
	assert.Equal(t, newPost.ID, retrievedPost.ID, "ID поста должен совпадать")

	// ✅ Логируем результат для отладки
	t.Logf("Полученные комментарии: %+v", retrievedPost.Comments)

	// 5️⃣ Проверяем, что комментарий сохранился
	if assert.NotNil(t, retrievedPost.Comments, "Список комментариев не должен быть nil") {
		assert.Equal(t, 1, len(retrievedPost.Comments), "Должен быть 1 комментарий")
		assert.Equal(t, newComment.ID, retrievedPost.Comments[0].ID, "ID комментария должен совпадать")
		assert.Equal(t, "This is a test comment.", retrievedPost.Comments[0].Content, "Текст комментария должен совпадать")
	}
} */
