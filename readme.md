# Сервис авторизации через GitHub

Этот сервис предоставляет API для авторизации пользователей через GitHub OAuth2. Пользователи могут войти в систему используя свой GitHub аккаунт, после чего их данные сохраняются в базе данных PostgreSQL.

## Описание сущности GitHubUser

Основная сущность сервиса - `GitHubUser`, которая содержит следующие поля:
* `ID` - Уникальный идентификатор пользователя в нашей системе (UUID)
* `GitHubID` - ID пользователя в системе GitHub
* `Login` - Логин пользователя в GitHub
* `Email` - Email пользователя
* `CreatedAt` - Время создания записи
* `UpdatedAt` - Время последнего обновления

## Используемые технологии

* **Язык программирования:** Go
* **База данных:** PostgreSQL
* **HTTP-роутер:** Gin-gonic
* **ORM:** GORM
* **OAuth2:** golang.org/x/oauth2
* **Генерация документации:** Swaggo

## Запуск приложения

### Требования

* Установленный Go (версия 1.18 или выше)
* Установленный PostgreSQL
* Зарегистрированное OAuth приложение на GitHub
* Настроенный конфигурационный файл

### Команды Makefile

* **`run-dev`**: Запускает приложение в режиме разработки
  * Использует config.yaml для конфигурации
  * Устанавливает переменную окружения `ENV_MODE` в `development`
* **`swag`**: Генерирует документацию API с помощью Swaggo
  * `--parseDependency`: Включает парсинг зависимостей
  * `--parseInternal`: Включает парсинг внутренних пакетов
  * `--parseDepth 3`: Устанавливает глубину парсинга зависимостей
  * `-g cmd/service/main.go`: Указывает главный файл приложения

### Запуск в режиме разработки

1. **Настройте GitHub OAuth приложение:**
   * Перейдите в GitHub -> Settings -> Developer settings -> OAuth Apps
   * Создайте новое приложение
   * Укажите Callback URL: http://localhost:8080/api/v1/auth/github/callback
   * Скопируйте Client ID и Client Secret

2. **Настройте конфигурационный файл:**
   * Создайте файл `config.yaml` в корне проекта
   * Заполните его по следующему шаблону:
   ```yaml
   Port: 8080
   Host: localhost
   DB:
       Host: "localhost"
       Port: 5432
       Username: "postgres"
       Name: "auth_db"
       Password: "your-password"
       SSLMode: "disable"
   github:
     client_id: "your-github-client-id"
     client_secret: "your-github-client-secret"
     redirect_url: "http://localhost:8080/api/v1/auth/github/callback"
   ```

3. **Запустите приложение:**
   ```bash
   make run-dev
   ```
   Сервис будет доступен по адресу http://localhost:8080

## Документация API

Документация API генерируется автоматически с помощью Swaggo и доступна по адресу `http://localhost:8080/swagger/index.html` после запуска приложения.

### Основные эндпоинты

* **GET /api/v1/** - Домашняя страница с кнопкой входа через GitHub
* **GET /api/v1/login** - Начало процесса авторизации через GitHub
* **GET /api/v1/auth/github/callback** - Обработка ответа от GitHub после авторизации

## Конфигурация

Конфигурация приложения осуществляется через файл `config.yaml`:

```yaml
# Основные настройки сервиса
Port: 8080              # Порт, на котором запускается сервис
Host: localhost         # Хост сервиса

# Настройки базы данных
DB:
    Host: localhost     # Хост базы данных
    Port: 5432         # Порт базы данных
    Username: postgres  # Пользователь базы данных
    Name: auth_db      # Название базы данных
    Password: password  # Пароль для базы данных
    SSLMode: disable   # Режим SSL

# Настройки GitHub OAuth
github:
    client_id: ""      # Client ID от GitHub OAuth App
    client_secret: ""  # Client Secret от GitHub OAuth App
    redirect_url: ""   # URL для callback после авторизации
```

## Внесение изменений

При изменении кода необходимо:

1. Обновить структуры данных в коде, если необходимо
2. Обновить миграции базы данных, если изменилась модель данных
3. Обновить документацию API, выполнив команду `make swag`

## Примечание
* Убедитесь, что все необходимые конфигурации установлены перед запуском приложения
* Никогда не коммитьте реальные значения client_secret в репозиторий
* При развертывании в production-среде используйте HTTPS для redirect_url
* Рекомендуется использовать переменные окружения для хранения чувствительных данных