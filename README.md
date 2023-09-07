
# Создание веб-сервера с использованием библиотеки net/http для обработки HTTP-запросов. Это приложение также использует HTTP Basic Authentication для защиты определенных эндпоинтов.

***
### Требования
Для запуска проекта вам понадобятся следующие компоненты:

- Go 1.20
- mkcert

***

1. Склонируйте репозиторий на вашу локальную машину:

   ```bash
   git clone https://github.com/Jhnvlglmlbrt/basic-auth

2. Нужно установить mkcert:

 [Ссылка](https://github.com/FiloSottile/mkcert )

3. Создать папку certs и сертификаты:

    ```bash
    mkdir certs && cd certs
    mkcert localhost

3. Перейдите в директорию проекта:

   ```bash
   cd basic-auth

3. Установить переменную окружения SSL_CERT_FILE, указывая путь к SSL-сертификату

    ```bash
    export SSL_CERT_FILE= путь/basic-auth/certs/ваше_название_сертификата

5. Запустите сервер с авторизационными данными:

    ```bash
    AUTH_USERNAME=логин AUTH_PASSWORD=пароль go run .

***

### Запросы с помощью Curl, чтобы убедиться, что проверки аутентификации работают правильно

    
    ```bash
    curl -i https://localhost:4000/unprotected

    curl -i https://localhost:4000/protected

    curl -i -u логин:пароль https://localhost:4000/protected

    curl -i -u логин:пароль https://localhost:4000/protected
    ```
