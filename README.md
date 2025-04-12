# Распределённый маркетплейс на микросервисах
В разработке. 

Планируется реализовать аналог маркетплейса как Ozon, Wildberries, Avito

## Независимые микросервисы

[Микросервис-авторизации](https://github.com/KlimenkoKayot/avito-go/tree/main/services/auth)

- - -

## Что я буду делать в первую очередь?

1. JWT для микросервиса [авторизации](https://github.com/KlimenkoKayot/avito-go/tree/main/services/auth)
2. Централизованный api-gateway
3. Микросервис profile
4. Kafka или Redis + адаптер к ним для гибкости

Обеспечить покрытие тестами
Обеспечить CI/CD
Обеспечить удобный деплой

## Запуск контейнера postgres для теста
```
sudo docker stop $(sudo docker ps -a -q)
sudo docker rm $(sudo docker ps -a -q)
docker run --name test-pg -e POSTGRES_PASSWORD=password -p 5432:5432 -d postgres:15
```
