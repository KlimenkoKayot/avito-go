# Распределённый маркетплейс на микросервисах
В разработке. 

Планируется реализовать аналог маркетплейса как Ozon, Wildberries, Avito

## Независимые микросервисы

[Микросервис-авторизации](https://github.com/KlimenkoKayot/avito-go/tree/main/services/auth)

- - -

## Что я буду делать в первую очередь?

1. Написать mok unit-тесты, интеграционные тесты
2. Общую точку входа в API (API-шлюз)
3. Написать market-микросервис

## Запуск контейнера PostgreSQL для теста
```
sudo docker stop $(sudo docker ps -a -q)
sudo docker rm $(sudo docker ps -a -q)
docker run --name test-pg -e POSTGRES_PASSWORD=password -p 5432:5432 -d postgres:15
```
