package screens

import (
	"jami/pkg/sound"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func createButton(win fyne.Window, textArray, soundSTR string) fyne.CanvasObject {
	if textArray == "" {
		return layout.NewSpacer()
	}
	button := widget.NewButton(textArray, func() {
		sound.PlayNotes(textArray)
		win.Canvas().SetOnTypedRune(func(r rune) {
			sound.PlayNotes(soundSTR + string(r))
		})
	})
	return button
}

// KeyboardScreen create Keyboard screen
func KeyboardScreen(win fyne.Window) fyne.CanvasObject {
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

	return fyne.NewContainerWithLayout(layout.NewGridLayout(1),
		soundSelector,
		fyne.NewContainerWithLayout(layout.NewGridLayout(9)),
		fyne.NewContainerWithLayout(layout.NewGridLayout(9)),
		fyne.NewContainerWithLayout(layout.NewGridLayout(9)),
		fyne.NewContainerWithLayout(layout.NewGridLayout(11),
			createButton(win, "2", soundSTR),
			createButton(win, "3", soundSTR),
			createButton(win, "", soundSTR),
			createButton(win, "5", soundSTR),
			createButton(win, "6", soundSTR),
			createButton(win, "7", soundSTR),
			createButton(win, "", soundSTR),
			createButton(win, "9", soundSTR),
			createButton(win, "0", soundSTR),
			createButton(win, "", soundSTR),
			createButton(win, "=", soundSTR),
		),
		fyne.NewContainerWithLayout(layout.NewGridLayout(12),
			createButton(win, "q", soundSTR),
			createButton(win, "w", soundSTR),
			createButton(win, "e", soundSTR),
			createButton(win, "r", soundSTR),
			createButton(win, "t", soundSTR),
			createButton(win, "y", soundSTR),
			createButton(win, "u", soundSTR),
			createButton(win, "i", soundSTR),
			createButton(win, "o", soundSTR),
			createButton(win, "p", soundSTR),
			createButton(win, "[", soundSTR),
			createButton(win, "]", soundSTR),
		),
		fyne.NewContainerWithLayout(layout.NewGridLayout(9),
			createButton(win, "a", soundSTR),
			createButton(win, "s", soundSTR),
			createButton(win, "", soundSTR),
			createButton(win, "f", soundSTR),
			createButton(win, "g", soundSTR),
			createButton(win, "", soundSTR),
			createButton(win, "j", soundSTR),
			createButton(win, "k", soundSTR),
			createButton(win, "l", soundSTR),
		),
		fyne.NewContainerWithLayout(layout.NewGridLayout(9),
			createButton(win, "z", soundSTR),
			createButton(win, "x", soundSTR),
			createButton(win, "c", soundSTR),
			createButton(win, "v", soundSTR),
			createButton(win, "b", soundSTR),
			createButton(win, "n", soundSTR),
			createButton(win, "m", soundSTR),
			createButton(win, ",", soundSTR),
			createButton(win, ".", soundSTR),
		),
		fyne.NewContainerWithLayout(layout.NewGridLayout(9)),
	)
}
