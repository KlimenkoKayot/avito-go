# Распределённый маркетплейс на микросервисах
 В разработке...

## Запуск контейнера PostgreSQL для теста
```
sudo docker stop $(sudo docker ps -a -q)
sudo docker rm $(sudo docker ps -a -q)
docker run --name test-pg -e POSTGRES_PASSWORD=password -p 5432:5432 -d postgres:15
```