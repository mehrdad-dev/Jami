package screens

import (
	"github.com/mehrdad-dev/Jami/pkg/sound"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func createButton(win fyne.Window, textArray string) fyne.CanvasObject {
	if textArray == "" {
		return layout.NewSpacer()
	}
	button := widget.NewButton(textArray, func() {
		sound.PlayNotes(textArray)
		win.Canvas().SetOnTypedRune(func(r rune) {
			sound.PlayNotes("VP SA9 SR9\n" + string(r))
		})
	})
	return button
}

// KeyboardScreen create Keyboard screen
func KeyboardScreen(win fyne.Window) fyne.CanvasObject {
	return fyne.NewContainerWithLayout(layout.NewGridLayout(1),
		fyne.NewContainerWithLayout(layout.NewGridLayout(9)),
		fyne.NewContainerWithLayout(layout.NewGridLayout(9)),
		fyne.NewContainerWithLayout(layout.NewGridLayout(9)),
		fyne.NewContainerWithLayout(layout.NewGridLayout(9)),
		fyne.NewContainerWithLayout(layout.NewGridLayout(11),
			createButton(win, "2"),
			createButton(win, "3"),
			createButton(win, ""),
			createButton(win, "5"),
			createButton(win, "6"),
			createButton(win, "7"),
			createButton(win, ""),
			createButton(win, "9"),
			createButton(win, "0"),
			createButton(win, ""),
			createButton(win, "="),
		),
		fyne.NewContainerWithLayout(layout.NewGridLayout(12),
			createButton(win, "q"),
			createButton(win, "w"),
			createButton(win, "e"),
			createButton(win, "r"),
			createButton(win, "t"),
			createButton(win, "y"),
			createButton(win, "u"),
			createButton(win, "i"),
			createButton(win, "o"),
			createButton(win, "p"),
			createButton(win, "["),
			createButton(win, "]"),
		),
		fyne.NewContainerWithLayout(layout.NewGridLayout(9),
			createButton(win, "a"),
			createButton(win, "s"),
			createButton(win, ""),
			createButton(win, "f"),
			createButton(win, "g"),
			createButton(win, ""),
			createButton(win, "j"),
			createButton(win, "k"),
			createButton(win, "l"),
		),
		fyne.NewContainerWithLayout(layout.NewGridLayout(9),
			createButton(win, "z"),
			createButton(win, "x"),
			createButton(win, "c"),
			createButton(win, "v"),
			createButton(win, "b"),
			createButton(win, "n"),
			createButton(win, "m"),
			createButton(win, ","),
			createButton(win, "."),
		),
		fyne.NewContainerWithLayout(layout.NewGridLayout(9)),
	)
}
