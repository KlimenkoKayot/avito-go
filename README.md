# avito-go
 В разработке...

## Запуск контейнера PostgreSQL для теста
```
docker run --name test-pg -e POSTGRES_PASSWORD=password -p 5432:5432 -d postgres:15
```