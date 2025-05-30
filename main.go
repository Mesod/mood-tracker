package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/enescakir/emoji"
	_ "github.com/mattn/go-sqlite3"
)

// Version will be set during build time
var Version = "dev"

type Mood struct {
	ID        int
	Mood      string
	Timestamp time.Time
}

var moods = []string{"happy", "sad", "neutral", "angry", "excited"}
var moodsEmoji = map[string]emoji.Emoji{
	"happy":   emoji.SmilingFace,
	"sad":     emoji.PensiveFace,
	"neutral": emoji.NeutralFace,
	"angry":   emoji.AngryFace,
	"excited": emoji.StarStruck,
}

const DB_PATH = "mood.db"

var DB *sql.DB

func initDB() {
	var err error
	DB, err = sql.Open("sqlite3", DB_PATH)
	if err != nil {
		log.Fatal(err)
	}

	_, err = DB.Exec(`CREATE TABLE IF NOT EXISTS moods (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		mood TEXT NOT NULL,
		timestamp DATETIME DEFAULT CURRENT_TIMESTAMP
	)`)
	if err != nil {
		log.Fatal(err)
	}

}

func main() {
	initDB()
	defer DB.Close()

	if len(os.Args) < 2 {
		fmt.Println("Usage: mood-tracker <command> [arguments]")
		fmt.Println("Commands: record <mood>, timeline")
		fmt.Printf("\nProudly made in %v \nversion: %s\n", emoji.FlagForIran, Version)
		return
	}

	switch os.Args[1] {
	case "record":
		recordMood(DB, os.Args[2:])
	case "timeline":
		showTimeline(DB)
	default:
		fmt.Println("Unknown command.")
	}
}

func recordMood(db *sql.DB, args []string) {
	if len(args) < 1 {
		fmt.Println("Please specify a mood:", moods)
		return
	}
	mood := args[0]
	valid := false
	for _, m := range moods {
		if m == mood {
			valid = true
			break
		}
	}
	if !valid {
		fmt.Println("Invalid mood. Valid options are:", moods)
		return
	}
	_, err := db.Exec("INSERT INTO moods (mood, timestamp) VALUES (?, ?)", mood, time.Now())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Mood recorded:", mood)
}

func showTimeline(db *sql.DB) {
	rows, err := db.Query("SELECT id, mood, timestamp FROM moods ORDER BY timestamp ASC")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	fmt.Println("Mood Timeline:")
	user_moods := []Mood{}
	for rows.Next() {
		var mood Mood
		if err := rows.Scan(&mood.ID, &mood.Mood, &mood.Timestamp); err != nil {
			log.Fatal(err)
		}
		user_moods = append(user_moods, mood)
	}
	for _, m := range user_moods {
		fmt.Println(moodsEmoji[m.Mood])
	}
}
