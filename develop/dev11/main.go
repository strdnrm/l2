package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

type Event struct {
	UserId int
	Date   time.Time
}

var store Store

func main() {
	port := flag.String("port", ":8080", "starting at port")
	flag.Parse()

	CreateDB("localhost", "5432", "admin", "password", "events")

	http.Handle("/", loggingMiddleware(http.DefaultServeMux))

	http.HandleFunc("/create_event", createEventHandler)
	http.HandleFunc("/update_event", updateEventHandler)
	http.HandleFunc("/delete_event", deleteEventHandler)
	http.HandleFunc("/events_for_day", eventsForDayHandler)
	http.HandleFunc("/events_for_week", eventsForWeekHandler)
	http.HandleFunc("/events_for_month", eventsForMonthHandler)

	fmt.Printf("Starting web server at port %s\n", *port)
	log.Fatal(http.ListenAndServe(*port, nil))
}

func sendJSONResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func parseUserid(data string) (int, error) {
	str, err := strconv.Atoi(data)
	if err != nil {
		return 0, err
	}
	return str, nil
}

func parseDate(data string) (time.Time, error) {
	date, err := time.Parse("2006-01-02", data)
	if err != nil {
		return date, err
	}
	return date, nil
}

func createEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	userID, err := parseUserid(r.FormValue("user_id"))
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	date, err := parseDate(r.FormValue("date"))
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	event := Event{
		UserId: userID,
		Date:   date,
	}

	err = CreateEvent(event)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	sendJSONResponse(w, map[string]string{"result": "Event created"})
}

func updateEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	userID, err := parseUserid(r.FormValue("user_id"))
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	date, err := parseDate(r.FormValue("date"))
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	event := Event{
		UserId: userID,
		Date:   date,
	}

	err = UpdateEvent(event)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	sendJSONResponse(w, map[string]string{"result": "Event created"})
}

func deleteEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	userID, err := parseUserid(r.FormValue("user_id"))
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	date, err := parseDate(r.FormValue("date"))
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	event := Event{
		UserId: userID,
		Date:   date,
	}

	err = DeleteEvent(event)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	sendJSONResponse(w, map[string]string{"result": "Event created"})
}

func eventsForDayHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	userID, err := parseUserid(r.URL.Query().Get("user_id"))
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	data, err := GetEventsByDay(userID)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	sendJSONResponse(w, data)

}

func eventsForWeekHandler(w http.ResponseWriter, r *http.Request) {
	// Реализация метода /events_for_week
}

func eventsForMonthHandler(w http.ResponseWriter, r *http.Request) {
	// Реализация метода /events_for_month
}

type Store struct {
	db *sql.DB
}

func CreateDB(host, port, user, password, dbname string) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	store.db = db
}

func CreateEvent(event Event) error {
	_, err := store.db.Exec("INSERT INTO events (user_id, date) VALUES (?, ?)", event.UserId, event.Date)
	if err != nil {
		return err
	}
	return nil
}

func UpdateEvent(event Event) error {
	_, err := store.db.Exec("UPDATE events SET date = ? WHERE user_id = ?", event.Date, event.UserId)
	if err != nil {
		return err
	}
	return nil
}

func DeleteEvent(event Event) error {
	_, err := store.db.Exec("DELETE FROM events WHERE date = ? AND user_id = ?", event.Date, event.UserId)
	if err != nil {
		return err
	}
	return nil
}

func GetEventsByDay(userId int) ([]Event, error) {
	rows, err := store.db.Query("SELECT user_id, date FROM events WHERE date == ? AND user_id = ?", time.Now().Format("01-02-2006"), userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	events := []Event{}
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.UserId, &event.Date)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func GetEventsByWeek(userId int) ([]Event, error) {
	rows, err := store.db.Query("SELECT user_id, date FROM events WHERE date >= ? AND date < ? AND user_id = ?",
		time.Now().Format("01-02-2006"), time.Now().AddDate(0, 0, 7).Format("01-02-2006"), userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	events := []Event{}
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.UserId, &event.Date)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func GetEventsByMonth(userId int) ([]Event, error) {
	rows, err := store.db.Query("SELECT user_id, date FROM events WHERE date >= ? AND date < ? AND user_id = ?",
		time.Now().Format("01-02-2006"), time.Now().AddDate(0, 1, 0).Format("01-02-2006"), userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	events := []Event{}
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.UserId, &event.Date)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}
