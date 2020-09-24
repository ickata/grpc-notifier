package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type Instance struct {
	db *sql.DB
}

func Init() *Instance {
	os.Remove("notifier.db")
	file, err := os.Create("notifier.db")
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()

	sqliteDatabase, _ := sql.Open("sqlite3", "./notifier.db")
	createTable(sqliteDatabase)

	return &Instance{db: sqliteDatabase}
}

func createTable(db *sql.DB) {
	SQL := `CREATE TABLE notifications (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"title" TEXT,
		"message" TEXT,
		"date" DATETIME		
	);`
	
	statement, err := db.Prepare(SQL)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
}

func (db *Instance) ScheduleNotification(title string, message string, date uint32) {
	SQL := `INSERT INTO notifications(title, message, date) VALUES (?, ?, ?)`
	statement, err := db.db.Prepare(SQL)
	
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(title, message, date)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

type Notification struct {
	Title string
	Message string
	Date string
}

func (db *Instance) GetNotifications() []Notification {
	row, err := db.db.Query("SELECT * FROM notifications ORDER BY date")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()

	var notifications []Notification

	for row.Next() {
		var id int
		var title string
		var message string
		var date string
		row.Scan(&id, &title, &message, &date)
		notification := Notification{
			Title: title,
			Message: message,
			Date: date,
		}
		notifications = append(notifications, notification)
	}

	return notifications
}

func (db *Instance) GetCurrentNotifications() []Notification {
	row, err := db.db.Query("SELECT * FROM notifications WHERE date BETWEEN STRFTIME('%s', 'now') AND STRFTIME('%s', 'now', '+5 second')")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()

	var notifications []Notification

	for row.Next() {
		var id int
		var title string
		var message string
		var date string
		row.Scan(&id, &title, &message, &date)
		notification := Notification{
			Title: title,
			Message: message,
			Date: date,
		}
		notifications = append(notifications, notification)
	}

	return notifications
}
