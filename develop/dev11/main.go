package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/create_event", logMiddleware(createEventHandler))
	mux.HandleFunc("/update_event", logMiddleware(updateEventHandler))
	mux.HandleFunc("/delete_event", logMiddleware(deleteEventHandler))
	mux.HandleFunc("/events_for_day", logMiddleware(eventsForDayHandler))
	mux.HandleFunc("/events_for_week", logMiddleware(eventsForWeekHandler))
	mux.HandleFunc("/events_for_month", logMiddleware(eventsForMonthHandler))

	port := "8080" // Порт можно изменить или получить из конфигурационного файла
	log.Printf("Starting server on port %s", port)
	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

func logMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	}
}
