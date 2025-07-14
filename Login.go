package main

import (
	"log"
	"os"

	"github.com/rivo/tview"
)

func Home() *tview.Flex {
	HomeContainer := tview.NewFlex().SetDirection(0)
	formContainer := tview.NewFlex().SetDirection(0)

	form := tview.NewForm().
		AddInputField("Enter your Name", "", 20, nil, nil).
		AddButton("Start", func() {
			StartQuiz()
		}).
		AddButton("Quit", func() {
			QuitApp()
		}).SetItemPadding(2)



	formContainer.AddItem(form, 8, 1, true)

	greeter := tview.NewTextView()
	greeterText, err := os.ReadFile("asciiGreeter.txt")

	if err != nil {
		log.Println(err)
	}

	greeter.SetText(string(greeterText))

	HomeContainer.AddItem(greeter, 6, 0, false)
	HomeContainer.AddItem(formContainer, 20, 0, true).SetBorder(true).SetTitle("GoQuiz - Login").SetBorderPadding(1, 1, 1, 1)

	return HomeContainer

}
