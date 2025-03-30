# Распределённый маркетплейс на микросервисах
В разработке. 

Планируется реализовать аналог маркетплейса как Ozon, Wildberries, Avito

- - -

## Важные коммиты

[Коммит-про-фундаментальную-структуру](https://github.com/KlimenkoKayot/avito-go/commit/641f8baf1d8c5819a98b4e80ec7f9a0d5e876a31)

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
