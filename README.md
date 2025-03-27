# Распределённый маркетплейс на микросервисах
В разработке. 

Планируется реализовать аналог маркетплейса как Ozon, Wildberries, Avito

- - -

## Что я буду делать в первую очередь?

1. Поменять временные логи на middleware
2. Написать mok unit-тесты, интеграционные тесты
3. Исправить неудобную струкутуру директорий, чтобы каждый микросервис реализовал полноценный список директорий (internal, cmd,..)
4. Общую точку входа в API (API-шлюз)

## Запуск контейнера PostgreSQL для теста
```
sudo docker stop $(sudo docker ps -a -q)
sudo docker rm $(sudo docker ps -a -q)
docker run --name test-pg -e POSTGRES_PASSWORD=password -p 5432:5432 -d postgres:15
```
