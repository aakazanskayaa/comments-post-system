Система постов и комментариев (GraphQL + Go + PostgreSQL)
Этот проект реализует систему постов и комментариев
Поддерживает GraphQL, подписки, пагинацию, вложенные комментарии и работает как с in-memory хранилищем, так и с PostgreSQL.

Функциональные возможности
1. Просмотр списка постов
2. Просмотр поста и комментариев под ним 
3. Создание постов и комментариев
4. Запрет комментариев к посту (по желанию автора)  
5. Иерархические (вложенные) комментарии
6. Ограничение длины комментария до 2000 символов  
7. Пагинация комментариев
8. GraphQL Subscriptions — асинхронные обновления комментариев в реальном времени  
9. Выбор хранилища: In-Memory или PostgreSQL  
10. Unit-тестирование основного функционала


Технологии
- Язык: Go  
- GraphQL : gqlgen  
- База данных : PostgreSQL или In-Memory хранилище  
- Хранение зависимостей : go modules  
- Docker : для контейнеризации  
- Тестирование : go test, assert  


 Запуск проекта  
1. 
docker-compose up --build
http://localhost:8080
2. STORAGE_TYPE=memory go run main.go
http://localhost:8080 


Запросы:
Получить все посты
query {
  posts {
    id
    title
    content
    author
    commentsAllowed
    createdAt
  }
}


Создать новый пост 
mutation {
  createPost(
    title: "Грустный пост"
    content: "Мне жарко"
    author: "Aнна"
    commentsAllowed: true
  ) {
    id
    title
    content
  }
}


Получить конкретный пост 
query {
  post(id: "") {
    id
    title
    content
    author
    commentsAllowed
    createdAt
  }
}


Добавить комментарий к посту 
mutation {
  addComment(
    postId: ""
    parentId: null
    author: "Анна"
    content: "Действительно грустный пост"
  ) {
    id
    content
    author
    createdAt
  }
}


Подписка на новые посты 
subscription {
  commentAdded(postId: "") {
    id
    content
    author
    createdAt
  }
}

Получить все комментарии к посту
query GetComments($postID: ID!, $limit: Int!, $offset: Int!) {
  comments(postID: $postID, limit: $limit, offset: $offset) {
    id
    postId
    parentId
    author
    content
    createdAt
  }
}

(Добавить в variables)
{
  "postID": "",
  "limit": 10,
  "offset": 0
}


Тестирование:
go test -v ./internal/graph/...

Ожидаемый резултат:
=== RUN   TestCreatePost
--- PASS: TestCreatePost (0.00s)
=== RUN   TestAddComment
--- PASS: TestAddComment (0.00s)
=== RUN   TestGetPostWithComments
--- PASS: TestGetPostWithComments (0.00s)
=== RUN   TestCreatePostWithoutComments
--- PASS: TestCreatePostWithoutComments (0.00s)
=== RUN   TestAddCommentToPostWithDisabledComments
--- PASS: TestAddCommentToPostWithDisabledComments (0.00s)
=== RUN   TestFullProcess
    resolvers_test.go:149: Полученные комментарии: [0x14000101080]
--- PASS: TestFullProcess (0.10s)
=== RUN   TestPaginationComments
--- PASS: TestPaginationComments (0.10s)
=== RUN   TestNestedComments
--- PASS: TestNestedComments (0.00s)
=== RUN   TestCommentSubscription
--- PASS: TestCommentSubscription (2.00s)
PASS
ok   
