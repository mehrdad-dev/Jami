package screens

import (
	"fmt"
	"github.com/mehrdad-dev/Jami/pkg/sound"
	"os"

	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"github.com/dbatbold/beep"
)

// SettingsScreen return setting canvas
func SettingsScreen(a fyne.App, win fyne.Window) fyne.CanvasObject {
	music := beep.NewMusic("")
	buttonDownload := widget.NewButtonWithIcon("Download files", theme.CheckButtonCheckedIcon(),
		func() {
			prog := dialog.NewProgressInfinite("Download and Saving Files",
				"Please wait...", win)

			go func() {
				err := sound.DownloadVoiceFiles(music, os.Stdout, []string{})
				if err != nil {
					fmt.Println("ERR in download files: ", err)
				}
				prog.Hide()
			}()

			prog.Show()
		})

	return widget.NewVBox(
		widget.NewLabelWithStyle("Change Theme", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		widget.NewSelect([]string{"Dark Theme", "Light Theme"}, func(s string) {
			if s == "Dark Theme" {
				a.Settings().SetTheme(theme.DarkTheme())
			} else {
				a.Settings().SetTheme(theme.LightTheme())
			}
		}),
		widget.NewLabelWithStyle("Download piano and violin sounds", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		widget.NewLabel("For download press button and wait"),
		buttonDownload,
	)
}
