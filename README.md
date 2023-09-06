# backend-trainee-assignment
# Тестовое задание в рамках отбора на стажировку в AvitoTech

Проект **Backend trainee assignment**  позволяет хранить пользователя и сегменты, в которых он состоит (создание, изменение, удаление сегментов, а также добавление и удаление пользователей в сегмент)

### Для запуска приложения:

```
make build && make run
```

Если приложение запускается впервые, необходимо применить миграции к базе данных:

```
make migrate
```

### Для авторизации используется Basic Auth: 
    1. `login`: admin
    2. `password`: qwerty
   
### Эндпойнты
    1. POST `/api/user/` - POST-запрос для создание пользователя
       - Принимает в body JSON: {"name":"username"}
       - В ответе выдаёт id созданного пользователя
    2. GET `/api/user/{userId}` - GET-запрос для получения сегментов пользователя
       - userId - это Id нужного пользователя
    3. POST `/api/segments/` - POST-запрос на создание сегмента 
       - Принимает в body JSON: {"name":"segment_name"}
       - В ответе выдаёт id созданного сегмента
    4. POST `/api/segments/{userId}` - POST-запрос на добавление сегментов пользователю 
       - Принмает в body JSON: {"segments":["segment_name1", "segment_name2"]}
       - В ответе выдаёт список Id которые были добавлены пользователю 
    5. DELETE `/api/segments/{userId}` - DELETE-запрос на удаление сегментов у пользователя
        - Принмает в body JSON: {"segments":["segment_name1", "segment_name2"]}
    6. DELETE `/api/segments/` - DELETE-запрос на удаление сегмента из таблицы сегментов.
        - Принимает в body JSON: {"name":"segment_name"}

