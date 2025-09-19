Веб-сервер, на котором реализован планировщик задач, который хранит их в базе данных. 
Задания повышенной сложности не выполнялись
Инструкция по запуску кода локально: 
  go build -o app.exe ./cmd/main.go
  ./app.exe
TODO_PORT = 7540
TODO_DBFILE = "../scheduler.db"

Тесты:
1. Тест сервера
go test -run ^TestApp$ ./tests 
2. Тест БД
go test -run ^TestDB$ ./tests
3. Тест функции NextDate
go test -run ^TestNextDate$ ./tests
4. Тест добавления задачи
go test -run ^TestAddTask$ ./tests
5. Тест получения ближайших задач
go test -run ^TestTasks$ ./tests
6. Тест редактирования задач
go test -run ^TestEditTask$ ./tests
7. Тест удаления задач
go test -run ^TestDelTask$ ./tests
8. Полный тест приложения с поддержкой аутентификации
go test ./tests
