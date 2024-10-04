# effective-mobile-task
- [Условие задачи](#task)
- [Использованные библиотеки](#libs)
- [Конфигурация](#config)
- [Функциональность](#func)
- [Установка](#management)
- [Тестирование](#test)

# Условие задачи <a name="task"/>

Реализация онлайн библиотеки песен 🎶

Необходимо реализовать следующее

1. Выставить rest методы
Получение данных библиотеки с фильтрацией по всем полям и пагинацией
Получение текста песни с пагинацией по куплетам
Удаление песни
Изменение данных песни
Добавление новой песни в формате

JSON

{
 "group": "Muse",
 "song": "Supermassive Black Hole"
}


2. При добавлении сделать запрос в АПИ, описанного сваггером

openapi: 3.0.3
info:
  title: Music info
  version: 0.0.1
paths:
  /info:
    get:
      parameters:
        - name: group
          in: query
          required: true
          schema:
            type: string
        - name: song
          in: query
          required: true
          schema:
            type: string
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/SongDetail'
        '400':
          description: Bad request
        '500':
          description: Internal server error
components:
  schemas:
    SongDetail:
      required:
        - releaseDate
        - text
        - link
      type: object
      properties:
        releaseDate:
          type: string
          example: 16.07.2006
        text:
          type: string
          example: Ooh baby, don't you know I suffer?\nOoh baby, can you hear me moan?\nYou caught me under false pretenses\nHow long before you let me go?\n\nOoh\nYou set my soul alight\nOoh\nYou set my soul alight
        link:
          type: string
          example: https://www.youtube.com/watch?v=Xsp3_a-PMTw


3. Обогащенную информацию положить в БД postgres (структура БД должна быть создана путем миграций при старте сервиса)
4. Покрыть код debug- и info-логами
5. Вынести конфигурационные данные в .env-файл
6. Сгенерировать сваггер на реализованное АПИ


# Использованные пакеты <a name="libs"/>

- Логгирование: (https://github.com/sirupsen/logrus)
- Миграции: [pressly/goose](https://github.com/pressly/goose)
- Работа с HTTP: [gin-gonic/gin](https://github.com/gin-gonic/gin)
- Тестирование: [stretchr/testify](https://github.com/stretchr/testify)

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

Makefile:  
build:
		go build -v ./cmd/app

test:
		go test -v -timeout 30s ./...

up:
	docker-compose up --force-recreate

down:
	docker-compose down
	docker image rm music-lib-app

.DEFAULT_GOAL := build

Из корневой дирректории запускаем команду make для сборки контейнера с приложением,
затем команду make up для старта сервиса(при этом создается структура БД с помощью миграции).
  
# Тестирование <a name="test"/>

Для удобства в корневой каталог помещен файл server_test.http для тестирования сервиса