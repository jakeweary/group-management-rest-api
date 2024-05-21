# REST API сервис управления группами людей

## Сборка и запуск

> [!NOTE]
> Конфигурация производится через переменные окружения и/или через `.env` файл, пример файла предоставлен в репозитории.

### Полный запуск через Docker Compose

```sh
docker compose up -d
```

### Через системный компилятор Go

```sh
# запуск БД (опционально)
docker compose up postgres -d

# сборка и запуск одной командой
go run cmd/rest-api/main.go

# либо раздельно
go build -o rest-api cmd/rest-api/main.go
./rest-api
```

### Через Docker (без Compose)

```sh
docker build . -t rest-api
docker run --rm -p 8080:8080 \
  -e 'LISTEN_ADDRESS=:8080' \
  -e 'DATABASE_URL=...' \
  rest-api
```

## Эндпоинты

### Создать группу

`POST /group`

```jsonc
// пример запроса:
{
  "parent_group": 1, // id родительской группы (необязательное поле)
  "name": "Имя группы"
}

// пример ответа:
{
  "group_id": 1
}
```

### Обновить группу

`PUT /group/{id}`

```jsonc
// пример запроса:
{
  "parent_group": 1, // id родительской группы (необязательное поле)
  "name": "Имя группы"
}
```

### Удалить группу

`DELETE /group/{id}`

> [!CAUTION]
> При удалении группы так же удаляются все дочерние группы и пользователи.

### Получить группу и всех её пользователей

`GET /group/{id}`

| Query параметр | Описание                                  |
| -------------- | ----------------------------------------- |
| `subgroups=1`  | Показать пользователей во всех подгруппах |

```jsonc
// пример ответа:
{
  "group": {
    "id": 2,
    "parent_group": 1, // может отсутствовать
    "name": "Имя группы"
  },
  "users": [
    {
      "id": 1,
      "group": 2,
      "first_name": "Имя",
      "last_name": "Фамилия",
      "birth_year": 2000
    }
  ]
}
```

### Создать пользователя

`POST /user`

```jsonc
// пример запроса:
{
  "group": 1, // id родительской группы (обязательное поле)
  "first_name": "Имя",
  "last_name": "Фамилия",
  "birth_year": 2000
}

// пример ответа:
{
  "user_id": 1
}
```

### Обновить пользователя

`PUT /user/{id}`

```jsonc
// пример запроса:
{
  "group": 1, // id родительской группы (обязательное поле)
  "first_name": "Имя",
  "last_name": "Фамилия",
  "birth_year": 2000
}
```

### Удалить пользователя

`DELETE /user/{id}`

### Получить пользователя

`GET /user/{id}`

```jsonc
// пример ответа:
{
  "user": {
    "id": 1,
    "group": 1,
    "first_name": "Имя",
    "last_name": "Фамилия",
    "birth_year": 2000
  }
}
```

### Получить все группы

`GET /groups`

| Query параметр | Описание                                 |
| -------------- | ---------------------------------------- |
| `subgroups=1`  | Считать пользователей во всех подгруппах |

```jsonc
// пример ответа:
{
  "groups": [
    {
      "id": 1,
      "name": "Юзеры",
      "user_count": 3
    },
    {
      "id": 2,
      "parent_group": 1,
      "name": "Модеры",
      "user_count": 2
    },
    {
      "id": 3,
      "parent_group": 2,
      "name": "Админы",
      "user_count": 1
    }
  ]
}
```

## Ошибки

Если при выполнении запроса возникла ошибка, API вернет статус 400 и JSON с полем `error`, где будет код и сообщение ошибки.

При непредвиденной внутренней ошибке API вернет пустой ответ без JSON со статусом 500.

Пример:

```jsonc
{
  "error": {
    "code": 100,
    "message": "parent group doesn't exist"
  }
}
```

### Таблица ошибок

| Код | Ошибка                     |
| --- | -------------------------- |
| 100 | parent group doesn't exist |
| 101 | group doesn't exist        |
| 102 | user doesn't exist         |
| 200 |	invalid url params         |
| 201 |	invalid url query params   |
| 202 |	invalid request body       |
