package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/rivo/tview"
)

var app *tview.Application
var pages = tview.NewPages()
var answers []int
var name string
var score int

func main() {
	app = tview.NewApplication()

	home := Home()
	quiz := Quiz()

	pages.AddPage("Home", home, true, true)
	pages.AddPage("Quiz", quiz, true, true)

	if err := app.SetRoot(pages.SwitchToPage("Home"), true).Run(); err != nil {
		panic(err)
	}

	if len(answers) >= 20 {
		fmt.Printf("Thanks for Playing! Your Score: %d / 20 \n", score)
	} else if len(answers) < 20 {
		fmt.Printf("Quiz Incomplete! %d / 20 questions attempted\n", len(answers))
	} else {
		return
	}
}

func QuitApp() {
	if app != nil {
		app.Stop()
	}
}

func StartQuiz() {
	if pages != nil {
		pages.SwitchToPage("Quiz")

	} else {
		error := fmt.Errorf("failed to start Quiz")
		fmt.Println(error)
	}
}

func SubmitQuiz() {

	payload := map[string]any{"name": name, "answers": answers}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling answers: %v", err)
		QuitApp()
		return
	}

	resp, err := http.Post("http://localhost:8080/score", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("Error submitting score to backend: %v", err)
		QuitApp()
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		var result map[string]int
		json.NewDecoder(resp.Body).Decode(&result)
		score = result["score"]
	} else {
		log.Printf("Error: Received non-OK status from backend: %d", resp.StatusCode)
	}
	QuitApp()
}
