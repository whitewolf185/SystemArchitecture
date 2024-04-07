# Сервис маршрутов и попутчиков

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
    value: mongodb://<user_name>:<password>@<host>:<port>

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
docker compose -f "domain-service/docker-compose.yaml" up -d --build
```

Поддерживаются следующие методы:
- GetCompanionInfo (GET)
- GetRouteInfo (GET)
- CreateRoute (POST)
- CreateCompanion (POST)
- DeleteRoute (DELETE)
- DeleteCompanion (DELETE)

После локального запуска, будет доступен swagger, где можно опробовать почти каждый метод.
Swagger http://<your_IP>:<your_port>/docs/index.html

### GetCompanionInfo (GET)
Получение списка попутчиков по фильтрам.

Пример запроса:
```bash
curl -X 'GET' \
  'http://localhost:8080/api/domain/GetCompanionInfo?client_id=fdskj&destination=destination' \
  -H 'accept: application/json'
```

Пример ответа:
```json
[
  {
    "client_id": "string",
    "destination": "string"
  }
]
```

### GetCompanionInfo (GET)
Получение списка маршрутов по фильтрам. 
_Один пользователь может иметь только один маршрут, но несколько пунктов назначения._

Пример запроса:
```bash
curl -X 'GET' \
  'http://localhost:8080/api/domain/GetRouteInfo?one_of_path=Moscow' \
  -H 'accept: application/json'
```

Пример ответа:
```json
[
  {
    "client_id": "b69bf55e-e044-4f11-8c4c-546e242eaca0",
    "path": [
      "Moscow",
      "Piter"
    ]
  }
]
```

### CreateRoute (POST)
Post запрос, который создает маршрут

Пример запроса:
```bash
curl -X 'POST' \
  'http://localhost:8080/api/domain/CreateRoute' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "client_id": "82c44825-3cb0-4c61-bb6e-4c781f28d597",
  "path": [
    "Moscow", "Piter", "Новосибирск"
  ]
}'
```

Пример ответа:
```json
{
  "client_id": "82c44825-3cb0-4c61-bb6e-4c781f28d597",
  "path": [
    "Moscow",
    "Piter",
    "Новосибирск"
  ]
}

```
### CreateCompanion (POST)
Post запрос, который создает попутчика с назначением куда бы он хотел добраться

Пример запроса:
```bash
curl -X 'POST' \
  'http://localhost:8080/api/domain/CreateCompanion' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "client_id": "82c44825-3cb0-4c61-bb6e-4c781f28d597",
  "destination": "Moscow"
}'
```

Пример ответа:
```json
{
  "client_id": "82c44825-3cb0-4c61-bb6e-4c781f28d597",
  "destination": "Moscow"
}
```


### DeleteRoute (DELETE)
Delete запрос, который удаляет сущность маршрута у пользователя.

Пример запроса:
```bash
curl -X 'DELETE' \
  'http://localhost:8080/api/domain/DeleteCompanion?client_id=82c44825-3cb0-4c61-bb6e-4c781f28d597' \
  -H 'accept: application/json'
```

Пример ответа:
В этом случае надо смотреть на коды ошибки. 
- Если 200, то все получилось; 
- если 404, то нечего удалять

### DeleteCompanion (DELETE)
Delete запрос, который удаляет сущность попутчика у пользователя.

Пример запроса:
```bash
curl -X 'DELETE' \
  'http://localhost:8080/api/domain/DeleteRoute?client_id=82c44825-3cb0-4c61-bb6e-4c781f28d597' \
  -H 'accept: application/json'
```

Пример ответа:
В этом случае надо смотреть на коды ошибки. 
- Если 200, то все получилось; 
- если 404, то нечего удалять