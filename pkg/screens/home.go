package screens

import (
	"fmt"
	"piano/pkg/sound"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func makeNoteTab(win fyne.Window) fyne.Widget {
	var soundSTR string

	// Creat sound selector
	soundSelector := widget.NewSelect([]string{"Piano", "Violin", "System"}, func(s string) {
		if s == "Piano" {
			soundSTR = "VP SA9 SR9\n"
		} else if s == "Violin" {
			soundSTR = "VV SA9 SR9\n"
		} else {
			soundSTR = "VD SA9 SR9\n"
		}
	})

	win.Canvas().SetOnTypedRune(func(r rune) {
		sound.PlayNotes(soundSTR + string(r))
		fmt.Println(string(r))
	})

	// Create input form
	notesInput := widget.NewMultiLineEntry()
	notesInput.SetPlaceHolder("Fill with notes")
	form := &widget.Form{
		OnSubmit: func() {
			fmt.Println("Notes:", notesInput.Text)
			sound.PlayNotes(soundSTR + notesInput.Text)
		},
	}
	form.Append("Notes", notesInput)
	return widget.NewVBox(
		widget.NewLabelWithStyle("Select Instrument", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		soundSelector,
		form,
	)
}

func makeKeyboardTab() fyne.Widget {
	list2 := widget.NewHBox()

	for i := 1; i <= 8; i++ {
		index := i
		button := widget.NewButton(fmt.Sprintf("Button %d", index), func() {
			fmt.Println("Tapped", index)
		})
		button.Resize(fyne.NewSize(50, 50))
		list2.Append(button)
	}
	return widget.NewVBox(
		widget.NewLabelWithStyle("\nEnter notes from keyboard", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		list2,
	)
}

// HomeScreen create home screen
func HomeScreen(win fyne.Window) fyne.CanvasObject {
	return fyne.NewContainerWithLayout(layout.NewGridLayout(),
		widget.NewTabContainer(
			widget.NewTabItem("Notes", makeNoteTab(win)),
			// widget.NewTabItem("Keyboard", makeKeyboardTab()),
		),
	)
}
