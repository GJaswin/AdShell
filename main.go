package main

import (
	"fmt"
	"github.com/rivo/tview"
)

var app *tview.Application
var pages = tview.NewPages()
var answers []int

func main() {
	app = tview.NewApplication()

	home := Home()
	quiz := Quiz()

	pages.AddPage("Home", home, true, true)
	pages.AddPage("Quiz", quiz, true, true)

	if err := app.SetRoot(pages.SwitchToPage("Home"), true).Run(); err != nil {
		panic(err)
	}

	fmt.Printf("Thanks for Playing! Your Score: %d / 20 \n", Scorer(answers))
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
	Scorer(answers)
	QuitApp()
}
