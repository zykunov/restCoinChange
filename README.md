REST сервис по нахождению разменов для разных сумм.
Сервис базируется на фреймворке echo (маршрутизация, логирование и сервер) 

Для запуска - 
go run .\cmd\currencyChange\main.go

Для запуска тестов 
1. cd .\internal\app\handlers\
2. go test

Структура файлов и папок - Standard Go Project Layout
/cmd/currencyChange/main.go - точка входа, подгрузка конфига, запуск сервера, роутинг, завершение работы сервера
/configs/.env - конфигурационный файл
/internal/app/handlers - Хэндлер. Тут находится сам алгоритм нахождения размена.
/models/ - модели для валидации входящего и исходящего json
