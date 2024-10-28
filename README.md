# REST API
- [Использованные библиотеки](#libs)
- [Конфигурация](#config)
- [Функциональность](#func)
- [Установка](#management)
- [Тестирование](#test)
- [Сваггер](#swagger)


# Использованные пакеты <a name="libs"/>

- Логгирование: [sirupsen/logrus](https://github.com/sirupsen/logrus)
- Миграции: [pressly/goose](https://github.com/pressly/goose)
- Работа с HTTP: [gin-gonic/gin](https://github.com/gin-gonic/gin)
- Тестирование: [stretchr/testify](https://github.com/stretchr/testify)
- Сваггер: [swaggo/swag](https://github.com/swaggo/swag)

# Конфигурация <a name="config"/>

Конфигурация вынесена в файл .env

# Функциональность <a name="func">

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
DELETE http://localhost:8000/delete/bfaa3ad1-2a8c-4c58-a7d6-4f3f913c37e6

- Изменение данных песни

Пример запроса:  
PATCH http://localhost:8000/update/bfaa3ad1-2a8c-4c58-a7d6-4f3f913c37e6  
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

# Установка <a name="management"/>

Из корневой дирректории запускаем команду make для сборки контейнера с приложением,
затем команду make up для старта сервиса(при этом создается структура БД с помощью миграции).

[Makefile](Makefile) 
  
# Тестирование <a name="test"/>

Для удобства в корневой каталог помещен файл server_test.http для тестирования сервиса (например с помощью расширения [REST Client](https://marketplace.visualstudio.com/items?itemName=humao.rest-client))

# Сваггер <a name="swagger"/>
[Swagger](docs/swagger.json)