package main

import (
	"github.com/mehrdad-dev/Jami/pkg/screens"
	"os"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func main() {
	a := app.NewWithID("com.jami.app")
	a.SetIcon(theme.FyneLogo())

	w := a.NewWindow("Jami")
	w.SetMainMenu(fyne.NewMainMenu(
		fyne.NewMenu("Control",
			fyne.NewMenuItem("Force Quit", func() { os.Exit(1) }),
		),
		fyne.NewMenu(
			"View",
			fyne.NewMenuItem("Full Screen", func() { w.SetFullScreen(true) }),
			fyne.NewMenuItem("Small Screen", func() { w.SetFullScreen(false) }),
		),
	),
	)
	w.SetMaster()

	tabs := widget.NewTabContainer(
		widget.NewTabItemWithIcon("Home", theme.HomeIcon(), screens.HomeScreen(w)),
		widget.NewTabItemWithIcon("Keyboard", theme.MediaPlayIcon(), screens.KeyboardScreen(w)),
		widget.NewTabItemWithIcon("Settings", theme.SettingsIcon(), screens.SettingsScreen(a, w)),
		widget.NewTabItemWithIcon("About", theme.InfoIcon(), screens.AboutScreen(a)))

	tabs.SetTabLocation(widget.TabLocationBottom)
	w.SetContent(tabs)
	w.Resize(fyne.NewSize(800, 500))
	// w.SetFixedSize(true)

	w.ShowAndRun()
}
