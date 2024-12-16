# Сервис подсчёта арифметических выражений

Данный веб-сервис вычисляет арифметические выражения, которые пользователь отправляет по HTTP в формате JSON. 
При этом результат выражения сервис отправляет также в формате JSON.

## Установка
1. Клонирование репозитория:

```git clone github```
2. Переход в директорию

```cd ./calc_project/cmd```
3. Запуск ПО

```go run ./calc_project/cmd```

## Работа с ПО
### У сервиса 1 endpoint с url-ом /api/v1/calculate. Пользователь отправляет на этот url POST-запрос со следующими вариантами:

1. {"expression": "выражение, которое ввёл пользователь"}
    ```curl --location 'localhost/api/v1/calculate' \
   --header 'Content-Type: application/json' \
   --data '{
   "expression": "2+2*2"
   }
 При валидном запросе должен быть ответ следующего вида:

    {
    "result": "результат выражения"
    }


## Manual Tests:
1. Send data: curl --location 'localhost/api/v1/calculate' \
   --header 'Content-Type: application/json' \
   --data '{
   "expression": "2+2*2"
   }
