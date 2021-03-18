package screens

import (
	"fmt"
	"github.com/mehrdad-dev/Jami/pkg/sound"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func makeHome(win fyne.Window) fyne.Widget {
	var soundSTR string

	// Create sound selector
	soundSelector := widget.NewSelect([]string{"Piano", "Violin", "System"}, func(s string) {
		if s == "Piano" {
			soundSTR = "VP SA9 SR9\n"
		} else if s == "Violin" {
			soundSTR = "VV SA9 SR9\n"
		} else {
			soundSTR = "VD SA9 SR9\n"
		}
	})

	// Create input form
	notesInput := widget.NewMultiLineEntry()
	notesInput.SetPlaceHolder("Fill with notes")
	form := &widget.Form{
		OnSubmit: func() {
			sound.PlayNotes(soundSTR + notesInput.Text)
		},
	}
	form.Append("Notes", notesInput)

	buttonSave := widget.NewButtonWithIcon("Save notes sound in .wav file", theme.DocumentSaveIcon(),
		func() {
			sound.SaveNotes(soundSTR + notesInput.Text)
		})

	return widget.NewGroupWithScroller("Home Page",
		widget.NewLabelWithStyle("Select Instrument", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		soundSelector,
		buttonSave,
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
	return fyne.NewContainerWithLayout(layout.NewAdaptiveGridLayout(1),
		makeHome(win),
		// widget.NewTabContainer(
		// 	widget.NewTabItem("Notes", makeNoteTab(win)),
		// 	// widget.NewTabItem("Keyboard", makeKeyboardTab()),
		// ),
	)
}
