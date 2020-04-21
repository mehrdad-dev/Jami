package main

import (
	"os"
	"piano/pkg/screens"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

const preferenceCurrentTab = "currentTab"

func main() {
	a := app.NewWithID("com.jami.demo")
	a.SetIcon(theme.FyneLogo())

	w := a.NewWindow("Jami")
	w.SetMainMenu(fyne.NewMainMenu(
		fyne.NewMenu("Contorol",
			fyne.NewMenuItem("Paste", func() { w.Clipboard() }),
			fyne.NewMenuItem("Force Quit", func() { os.Exit(1) }),
		),
		fyne.NewMenu(
			"View",
			fyne.NewMenuItem("Small Screen", func() { w.SetFullScreen(false) }),
			fyne.NewMenuItem("Full Screen", func() { w.SetFullScreen(true) }),
		),
	),
	)
	w.SetMaster()

	tabs := widget.NewTabContainer(
		widget.NewTabItemWithIcon("Home", theme.HomeIcon(), screens.HomeScreen(w)),
		widget.NewTabItemWithIcon("Learn", theme.HelpIcon(), screens.LearnScreen()),
		widget.NewTabItemWithIcon("Settings", theme.SettingsIcon(), screens.SettingsScreen(a, w)),
		widget.NewTabItemWithIcon("About", theme.InfoIcon(), screens.AboutScreen(a)))

	tabs.SetTabLocation(widget.TabLocationLeading)
	tabs.SelectTabIndex(a.Preferences().Int(preferenceCurrentTab))
	w.SetContent(tabs)
	w.SetFullScreen(true)
	w.SetPadded(true)

	// w.Resize(fyne.NewSize(700, 400))
	// w.SetFixedSize(true)

	w.ShowAndRun()
	a.Preferences().SetInt(preferenceCurrentTab, tabs.CurrentTabIndex())
}