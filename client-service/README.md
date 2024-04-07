# Клиентский сервис

# Описание
Перед вами микросервис, который отвечает за клиентскую информацию в сервисе.
На его плечах лежат такие задачи, как:
- создание пользователей
- удаление
- поиск

# Документация
## Запуск микросервиса

Для запуска микросервиса необходимо расположить конфигурирующий yaml файл с секретами по следующему пути:

```
./internal/config/flags/values_local.yaml
```
Либо указать путь к файлу с секретами можно при помощи env переменной `ENV_FILE_NAME`. 
Путь желательно использовать абсолютный.

Пример файла конфигурации выглядит так:
```yaml
env:
  - name: db_dsn
    value: postgres:///

  - name: listen_port
    value: 80
  - name: server_ip
    value: 127.0.0.1
```

### Описание переменных конфигурации

`db_dsn` -- ссылка для подключения к базе данных.

## API

Команда для запуска сервиса в докере:
```bash
docker compose -f "client-service/docker-compose.yaml" up -d --build
```

Поддерживаются следующие методы:
- CreateUser (POST)
- SearchUserByUserName (GET)
- GetClientByID (GET)
- DeleteUserByID (DELETE)

После локального запуска, будет доступен swagger, где можно опробовать почти каждый метод.
Swagger http://<your_IP>:<your_port>/docs/index.html

### CreateUser (POST)
Post запрос для создания пользователя

Пример запроса:
```bash
curl -X 'POST' \
  'http://localhost:8080/api/person/CreateUser' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "is_driver": true,
  "password": "string",
  "user_name": "another_string"
}'
```

Пример ответа:
```json
{
  "id": "65e44dae-61e5-4589-abac-ee71c2f01fac",
  "user_name": "another_string",
  "driver": true
}
```

### SearchUserByUserName (GET)
Get запрос, который выдает пользователей, у которых совпал substring в имени пользователя

Пример запроса:
```bash
curl -X 'GET' \
  'http://localhost:8080/api/person/SearchUserByUserName?user_name_in=string' \
  -H 'accept: application/json'
```

Пример ответа:
```json

[
  {
    "driver": true,
    "id": "string",
    "user_name": "string"
  }
]

```
### GetClientByID (GET)
Get запрос, который находит пользователя по его UUID
Пример запроса:
```bash
curl -X 'GET' \
  'http://localhost:8080/api/person/GetClientByID?id=65e44dae-61e5-4589-abac-ee71c2f01fac' \
  -H 'accept: application/json'
```

Пример ответа:
```json
{
  "driver": true,
  "id": "string",
  "user_name": "string"
}

```


### DeleteUserByID (DELETE)
Delete запрос, который удаляет пользователя по его UUID. Причем, повторив данную операцию, будет выдаваться ошибка, сигнализирующая, что пользователя больше нет.

Пример запроса:
```bash
curl -X 'DELETE' \
  'http://localhost:8080/api/person/DeleteUserByID?id=65e44dae-61e5-4589-abac-ee71c2f01fac' \
  -H 'accept: application/json'
```

Пример ответа:
```json
{
  "driver": true,
  "id": "string",
  "user_name": "string"
}

```