package main

import (
	"github.com/Parside01/dev11/internal/repository"
	service "github.com/Parside01/dev11/internal/service"
	"github.com/Parside01/dev11/internal/transport"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

/*
=== HTTP server ===

Реализовать HTTP сервер для работы с календарем. В рамках задания необходимо работать строго со стандартной HTTP библиотекой.
В рамках задания необходимо:
	1. Реализовать вспомогательные функции для сериализации объектов доменной области в JSON.
	2. Реализовать вспомогательные функции для парсинга и валидации параметров методов /create_event и /update_event.
	3. Реализовать HTTP обработчики для каждого из методов API, используя вспомогательные функции и объекты доменной области.
	4. Реализовать middleware для логирования запросов
Методы API: POST /create_event POST /update_event POST /delete_event GET /events_for_day GET /events_for_week GET /events_for_month
Параметры передаются в виде www-url-form-encoded (т.е. обычные user_id=3&date=2019-09-09).
В GET методах параметры передаются через queryString, в POST через тело запроса.
В результате каждого запроса должен возвращаться JSON документ содержащий либо {result: ...} в случае успешного выполнения метода,
либо {error: ...} в случае ошибки бизнес-логики.

В рамках задачи необходимо:
	1. Реализовать все методы.
	2. Бизнес логика НЕ должна зависеть от кода HTTP сервера.
	3. В случае ошибки бизнес-логики сервер должен возвращать HTTP 503. В случае ошибки входных данных (невалидный int например) сервер должен возвращать HTTP 400. В случае остальных ошибок сервер должен возвращать HTTP 500. Web-сервер должен запускаться на порту указанном в конфиге и выводить в лог каждый обработанный запрос.
	4. Код должен проходить проверки go vet и golint.
*/

func init() {
	_ = godotenv.Load(".env")
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	repo := repository.NewUserRepository()
	s := service.NewEventService(repo)
	u := service.NewUserService(repo)

	http.Handle("/create_user", transport.LogMiddleware(http.HandlerFunc(transport.NewUserCreateHandler(u).CreateUser)))
	http.Handle("/create_event", transport.LogMiddleware(http.HandlerFunc(transport.NewCreateHandler(s).CreateEvent)))
	http.Handle("/update_event", transport.LogMiddleware(http.HandlerFunc(transport.NewUpdateHandler(s).UpdateEvent)))
	http.Handle("/delete_event", transport.LogMiddleware(http.HandlerFunc(transport.NewDeleteEventHandler(s).DeleteEvent)))

	http.Handle("/events_for_day", transport.LogMiddleware(http.HandlerFunc(transport.NewEventsForDayHandler(s).EventsForDay)))
	http.Handle("/events_for_week", transport.LogMiddleware(http.HandlerFunc(transport.NewEventsForWeekHandler(s).EventsForWeek)))
	http.Handle("/events_for_month", transport.LogMiddleware(http.HandlerFunc(transport.NewEventsForMonthHandler(s).EventsForMonth)))

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
