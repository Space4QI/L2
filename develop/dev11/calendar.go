package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

// Event представляет событие в календаре
type Event struct {
	UserID int       `json:"user_id"` // Идентификатор пользователя
	Date   time.Time `json:"date"`    // Дата события
	Title  string    `json:"title"`   // Заголовок события
}

// Хранилище для событий, где ключом является дата в формате "YYYY-MM-DD"
var events = make(map[string][]Event)

// Обработчик для создания нового события
func createEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		// Если метод не POST, возвращаем ошибку 405 Method Not Allowed
		http.Error(w, `{"error": "Invalid method"}`, http.StatusMethodNotAllowed)
		return
	}

	var e Event
	if err := parseForm(r, &e); err != nil {
		// Если произошла ошибка при парсинге формы, возвращаем ошибку 400 Bad Request
		http.Error(w, `{"error": "Invalid input"}`, http.StatusBadRequest)
		return
	}

	// Формируем ключ по дате события
	key := e.Date.Format("2006-01-02")
	// Добавляем событие в хранилище
	events[key] = append(events[key], e)
	// Возвращаем успешный ответ в формате JSON
	respondWithJSON(w, map[string]string{"result": "Event created"})
}

// Обработчик для обновления существующего события
func updateEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		// Если метод не POST, возвращаем ошибку 405 Method Not Allowed
		http.Error(w, `{"error": "Invalid method"}`, http.StatusMethodNotAllowed)
		return
	}

	var e Event
	if err := parseForm(r, &e); err != nil {
		// Если произошла ошибка при парсинге формы, возвращаем ошибку 400 Bad Request
		http.Error(w, `{"error": "Invalid input"}`, http.StatusBadRequest)
		return
	}

	// Формируем ключ по дате события
	key := e.Date.Format("2006-01-02")
	eventsForDay := events[key]
	// Ищем событие с таким же user_id для обновления
	for i, ev := range eventsForDay {
		if ev.UserID == e.UserID {
			eventsForDay[i] = e
			events[key] = eventsForDay
			// Возвращаем успешный ответ в формате JSON
			respondWithJSON(w, map[string]string{"result": "Event updated"})
			return
		}
	}
	// Если событие не найдено, возвращаем ошибку 404 Not Found
	http.Error(w, `{"error": "Event not found"}`, http.StatusNotFound)
}

// Обработчик для удаления события
func deleteEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		// Если метод не POST, возвращаем ошибку 405 Method Not Allowed
		http.Error(w, `{"error": "Invalid method"}`, http.StatusMethodNotAllowed)
		return
	}

	var e Event
	if err := parseForm(r, &e); err != nil {
		// Если произошла ошибка при парсинге формы, возвращаем ошибку 400 Bad Request
		http.Error(w, `{"error": "Invalid input"}`, http.StatusBadRequest)
		return
	}

	// Формируем ключ по дате события
	key := e.Date.Format("2006-01-02")
	eventsForDay := events[key]
	// Ищем событие с таким же user_id для удаления
	for i, ev := range eventsForDay {
		if ev.UserID == e.UserID {
			eventsForDay = append(eventsForDay[:i], eventsForDay[i+1:]...)
			events[key] = eventsForDay
			// Возвращаем успешный ответ в формате JSON
			respondWithJSON(w, map[string]string{"result": "Event deleted"})
			return
		}
	}
	// Если событие не найдено, возвращаем ошибку 404 Not Found
	http.Error(w, `{"error": "Event not found"}`, http.StatusNotFound)
}

// Обработчик для получения событий за конкретный день
func eventsForDayHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		// Если метод не GET, возвращаем ошибку 405 Method Not Allowed
		http.Error(w, `{"error": "Invalid method"}`, http.StatusMethodNotAllowed)
		return
	}

	dateStr := r.URL.Query().Get("date")
	if dateStr == "" {
		// Если не указан параметр date, возвращаем ошибку 400 Bad Request
		http.Error(w, `{"error": "Missing date parameter"}`, http.StatusBadRequest)
		return
	}

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		// Если произошла ошибка при парсинге даты, возвращаем ошибку 400 Bad Request
		http.Error(w, `{"error": "Invalid date format"}`, http.StatusBadRequest)
		return
	}

	// Формируем ключ по дате события
	key := date.Format("2006-01-02")
	// Возвращаем события за указанный день в формате JSON
	respondWithJSON(w, map[string][]Event{"events": events[key]})
}

// Обработчик для получения событий за неделю
func eventsForWeekHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		// Если метод не GET, возвращаем ошибку 405 Method Not Allowed
		http.Error(w, `{"error": "Invalid method"}`, http.StatusMethodNotAllowed)
		return
	}

	dateStr := r.URL.Query().Get("date")
	if dateStr == "" {
		// Если не указан параметр date, возвращаем ошибку 400 Bad Request
		http.Error(w, `{"error": "Missing date parameter"}`, http.StatusBadRequest)
		return
	}

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		// Если произошла ошибка при парсинге даты, возвращаем ошибку 400 Bad Request
		http.Error(w, `{"error": "Invalid date format"}`, http.StatusBadRequest)
		return
	}

	// Определяем начало недели и конец недели
	startOfWeek := startOfWeek(date)
	endOfWeek := startOfWeek.AddDate(0, 0, 7)

	var weeklyEvents []Event
	// Собираем события за каждый день недели
	for d := startOfWeek; d.Before(endOfWeek); d = d.AddDate(0, 0, 1) {
		key := d.Format("2006-01-02")
		weeklyEvents = append(weeklyEvents, events[key]...)
	}

	// Возвращаем события за неделю в формате JSON
	respondWithJSON(w, map[string][]Event{"events": weeklyEvents})
}

// Обработчик для получения событий за месяц
func eventsForMonthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		// Если метод не GET, возвращаем ошибку 405 Method Not Allowed
		http.Error(w, `{"error": "Invalid method"}`, http.StatusMethodNotAllowed)
		return
	}

	dateStr := r.URL.Query().Get("date")
	if dateStr == "" {
		// Если не указан параметр date, возвращаем ошибку 400 Bad Request
		http.Error(w, `{"error": "Missing date parameter"}`, http.StatusBadRequest)
		return
	}

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		// Если произошла ошибка при парсинге даты, возвращаем ошибку 400 Bad Request
		http.Error(w, `{"error": "Invalid date format"}`, http.StatusBadRequest)
		return
	}

	// Определяем начало месяца и конец месяца
	startOfMonth := time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, time.UTC)
	endOfMonth := startOfMonth.AddDate(0, 1, 0)

	var monthlyEvents []Event
	// Собираем события за каждый день месяца
	for d := startOfMonth; d.Before(endOfMonth); d = d.AddDate(0, 0, 1) {
		key := d.Format("2006-01-02")
		monthlyEvents = append(monthlyEvents, events[key]...)
	}

	// Возвращаем события за месяц в формате JSON
	respondWithJSON(w, map[string][]Event{"events": monthlyEvents})
}

// Функция для парсинга формы из запроса и заполнения структуры Event
func parseForm(r *http.Request, e *Event) error {
	if err := r.ParseForm(); err != nil {
		return err
	}

	userIDStr := r.FormValue("user_id")
	dateStr := r.FormValue("date")
	title := r.FormValue("title")

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		return err
	}

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return err
	}

	e.UserID = userID
	e.Date = date
	e.Title = title

	return nil
}

// Функция для отправки ответа в формате JSON
func respondWithJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		// Если произошла ошибка при кодировании JSON, возвращаем ошибку 500 Internal Server Error
		http.Error(w, `{"error": "Internal server error"}`, http.StatusInternalServerError)
	}
}

// Функция для определения начала недели по данной дате
func startOfWeek(date time.Time) time.Time {
	weekday := int(date.Weekday())
	if weekday == 0 {
		weekday = 7 // Если неделя начинается с воскресенья, то устанавливаем 7
	}
	return date.AddDate(0, 0, -weekday+1) // Вычисляем начало недели
}
