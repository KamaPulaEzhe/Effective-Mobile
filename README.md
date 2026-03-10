# Effective-Mobile Subscription Service
REST сервис для управления онлайн подписками пользователей. Тестовое задание.

## Стек
- Go + Gin
- PostgreSQL
- Docker Compose
- Swagger

## Запуск
```bash
docker compose up --build
```
Сервис поднимется на `http://localhost:8000`

Swagger: `http://localhost:8000/swagger/index.html`

## Эндпоинты
```
POST   /api/subscriptions/           - создать подписку
GET    /api/subscriptions/:id        - получить подписки пользователя
PATCH  /api/subscriptions/:id        - обновить подписку
DELETE /api/subscriptions/:id        - удалить подписку
GET    /api/subscriptions/total-cost - сумма подписок за период
```

## Примеры

Создать подписку:
```json
POST /api/subscriptions/
{
    "service_name": "Yandex Plus",
    "price": 400,
    "user_id": "60601fee-2bf1-4721-ae6f-7636e79a0cba",
    "start_date": "07-2025"
}
```

Получить все подписки пользователя:
```
GET /api/subscriptions/60601fee-2bf1-4721-ae6f-7636e79a0cba
```

Получить подписку по сервису:
```
GET /api/subscriptions/60601fee-2bf1-4721-ae6f-7636e79a0cba?name=Netflix
```

Обновить подписку (все поля опциональны, но минимум одно):
```json
PATCH /api/subscriptions/c5b462f6-0493-4f83-bd77-91f8786fa23c
{
    "price": 599
}
```
```json
PATCH /api/subscriptions/c5b462f6-0493-4f83-bd77-91f8786fa23c
{
    "service_name": "Yandex Plus",
    "price": 299,
    "end_date": "12-2025"
}
```

Удалить подписку:
```
DELETE /api/subscriptions/60601fee-2bf1-4721-ae6f-7636e79a0cba?name=Netflix
```

Сумма за период:
```
GET /api/subscriptions/total-cost?user_id=60601fee-2bf1-4721-ae6f-7636e79a0cba&start_date=01-2025&end_date=12-2025
```

Сумма за выбранный период с фильтрацией по id пользователя и названию подписки:
```
GET /api/subscriptions/total-cost?user_id=60601fee-2bf1-4721-ae6f-7636e79a0cba&start_date=01-2025&end_date=12-2025&service_name=Netflix
```