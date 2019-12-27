package main

// import "fmt"
import "github.com/carolina-colagrossi/daily-theme-telegram-bot/pkg/journal"

func main() {
	_, err := journal.NewJournal("test.db")
	if err != nil {
		panic(err)
	}
}
