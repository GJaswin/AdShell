package main

import (
	"github.com/rivo/tview"
	"os"
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

	greeterText, _ := os.ReadFile("asciiGreeter.txt")

	greeter := tview.NewTextView()
	greeter.SetText(string(greeterText))

	HomeContainer.AddItem(greeter, 6, 0, false)
	HomeContainer.AddItem(formContainer, 20, 0, true).SetBorder(true).SetTitle("GoQuiz - Login")

	return HomeContainer

}
