# REST API
- [Использованные библиотеки](#libs)
- [Конфигурация](#config)
- [Функциональность](#func)
- [Интеграция с API](#API)
- [Установка](#management)
- [Тестирование](#test)
- [Сваггер](#swagger)


# Использованные пакеты <a name="libs"/>

- Логгирование: [sirupsen/logrus](https://github.com/sirupsen/logrus)
- Миграции: [pressly/goose](https://github.com/pressly/goose)
- Работа с HTTP: [gin-gonic/gin](https://github.com/gin-gonic/gin)
- Тестирование: [stretchr/testify](https://github.com/stretchr/testify)
- Сваггер: [swaggo/swag](https://github.com/swaggo/swag)
- Работа с БД Postgres: database/sql

# Конфигурация <a name="config"/>

Конфигурация вынесена в файл .env

# Функциональность <a name="func"/>

- Получение данных библиотеки с фильтрацией по всем полям и пагинацией

Пример запроса:  
POST http://localhost:8000/songs  
Content-Type: application/json

{  
    "page": 1,  
    "per_page": 8,  
    "group": "Muse"  
}

- Удаление песни

Пример запроса:  
DELETE http://localhost:8000/delete/1

- Изменение данных песни

Пример запроса:  
PATCH http://localhost:8000/update/2
Content-Type: application/json

{  
"group": "Muse",  
"song": "Supermassive Black Hole"  
}

- Добавление новой песни в формате

Пример запроса:  
POST http://localhost:8000/add  
Content-Type: application/json  

  
{  
"group": "Muse",  
"song": "Supermassive Black Hole"  
}

- Получение текста песни с пагинацией по куплетам

Пример запроса:
POST http://localhost:8000/songtext/1
Content-Type: application/json

{
    "page": 1,
    "per_page": 8,
    "id": 1
}

# Интеграция с внешним API <a name="API"/>

При добавлении песни делается запрос к внешнему API по адресу http://localhost:8088/info
для обогащения информации о песне.

# Установка <a name="management"/>

Из корневой дирректории запускаем команду make для сборки контейнера с приложением,
затем команду make up для старта сервиса(при этом создается структура БД с помощью миграции).

[Makefile](Makefile) 
  
# Тестирование <a name="test"/>

Для удобства в корневой каталог помещен файл server_test.http для тестирования сервиса (например с помощью расширения [REST Client](https://marketplace.visualstudio.com/items?itemName=humao.rest-client))

# Сваггер <a name="swagger"/>
[Swagger](docs/swagger.json)