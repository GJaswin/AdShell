package main

import (
	"github.com/rivo/tview"
)

func SubmitPage() *tview.Flex {
	SubmitPageContainer := tview.NewFlex().SetDirection(0)

	prompt := tview.NewTextView()
	prompt.SetText("Submit Quiz?")


	SubmitPageForm := tview.NewForm().
		AddButton("Submit", func() {
			SubmitQuiz()
		}).
		AddButton("Exit", func() {
			QuitApp()
		}).SetItemPadding(1)

	SubmitPageContainer.AddItem(prompt, 1, 0, false).
		AddItem(SubmitPageForm, 0, 1, true)

		SubmitPageContainer.SetBorder(true).SetTitle("Submit").SetBorderPadding(1,1,1,1)

	return SubmitPageContainer

}
