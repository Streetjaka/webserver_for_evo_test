# webserver_for_evo_test
Проверить работу сервера можно по адресу http://91.223.180.76:8080/

Запуск локального сервера.

Можно запустить скомпилированный файл main.exe (при работе на Windows)/main (при работе на Ubuntu) который запустит сервер на локальной машине с подключением к БД на хостинге и решением задания.
Порт сервера 8080. Если сервер не запустился или выдал ошибку/вылетела программа - остановить в диспетчере служб ISS веб-сервер.
Можно запустить исходник используя golang. Для этого установить golang по ссылке https://go.dev/doc/install
Или установить через команду sudo snap install go --classic (Ubuntu)
и запустить скрипт main.go командой go run main.go (в той же папке должны быть все файлы и папки из репозитория.
Перед запуском разрешить порт 8080 ($ sudo ufw allow 8080) - Для Linux
Сервер запускается по адресу localhost:8080
Для работы сервера нужно интернет-соединение так как база данных находится на хостинге.
