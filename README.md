# Сервис подсчёта арифметических выражений

Данный веб-сервис работает на порту 8080 и вычисляет арифметические выражения, которые пользователь отправляет по HTTP в формате JSON . 
При этом результат выражения сервис отправляет также в формате JSON.

## Запуск

1. Установить GO 1.23
2. Клонировать репозиторий:
```bash
   git clone https://github.com/serjpdev/WebServiceCalculation.git
```
3. Переход в директорию проекта
```bash
   cd ./WebServiceCalculation
```
4. Запуск ПО
```bash
   go run ./cmd/main.go
```
5. В случае необходимости можно сменить порт, на котором работает приложение. Это реализовано через переменную окружения PORT. 
Пример (для Linux):
```bash  
export PORT=80 
``` 
После этого запустить приложение (пункт 4).
## Работа с ПО
### У сервиса 1 endpoint с url-ом /api/v1/calculate. Пользователь отправляет на этот url POST-запрос:
```bash
   curl --location 'localhost:8080/api/v1/calculate' \
   --header 'Content-Type: application/json' \
   --data '{"expression": "1+1"}'
```

### Разрешённые символы
- Знаки операций (только бинарные): `+`, `-`, `*`, `/`
- Знаки приоритизации: `(`, `)`
- Рациональные числа (целые, либо через `.`)
### Ответы сервера 
   1. При валидном запросе сервер отвечает с кодом 200 и ответом:
```
     {"result": "результат выражения"}
```    
   2. Если входные данные не соответствуют требованиям приложения — например, кроме цифр и разрешённых операций пользователь ввёл символ английского алфавита. Также подобная ошибка отдается в случае обращения к несуществующему endpoint или иным методом запроса HTTP. То сервер вернет ошибку с кодом 422 и ответ:
```
      {"error": "Expression is not valid"}
```
   3. В случае какой-либо иной ошибки, сервер отвечает с кодом 500 и сообщением:
```
      {"error": "Internal server error"}
```
## Логирование данных 
Организовано логирование в поток вывода и ошибок.
## Автоматические тесты (запуск из директории проекта)
```bash
go test ./... 
```
## Ручное тестирование Linux curl:
1. Отправка валидных данных: 
```bash
curl --location 'localhost:8080/api/v1/calculate' \
   --header 'Content-Type: application/json' \
   --data '{
   "expression": "2+2*2"
   }'
```
2. Отправка невалидных данных:
```bash
curl --location 'localhost:8080/api/v1/calculate' \
   --header 'Content-Type: application/json' \
   --data '{
   "expression": "2+2*2)"
   }'
```
3. Отправка невалидного метода HTTP:
```bash
curl --location --request GET 'localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{"expression": "1+1"}'
```
4. Отправка на невалидный endpoint:
```bash
curl --location 'localhost:8080/api/v1/' \
--header 'Content-Type: application/json' \
--data '{"expression": "1+1"}'
```
## Тестирование с помощью Postman
    В директории postmanTests внутри проекта находится файл для импорта конфига в Postman. 