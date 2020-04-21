package screens

import (
	"fmt"
	"net/url"
	"path/filepath"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func parseURL(urlStr string) *url.URL {
	link, err := url.Parse(urlStr)
	if err != nil {
		fyne.LogError("Could not parse URL", err)
	}

	return link
}

func AboutScreen(a fyne.App) fyne.CanvasObject {
	filePath, err := filepath.Abs("../Jami/about.png")
	if err != nil {
		fmt.Println("ERR in read about.png ", err)
	}

	logo := canvas.NewImageFromFile(filePath)
	logo.SetMinSize(fyne.NewSize(390, 212))

	return widget.NewVBox(
		widget.NewHBox(layout.NewSpacer(), logo, layout.NewSpacer()),
		widget.NewLabelWithStyle("Jami is not just a musical instrument", fyne.TextAlignCenter, fyne.TextStyle{Bold: false}),

		widget.NewHBox(layout.NewSpacer(),
			widget.NewHyperlink("github", parseURL("https://github.com/mehrdad-dev")),
			widget.NewLabel("-"),
			widget.NewHyperlink("contact", parseURL("http://www.mehrdad-dev.ir/contact-me/")),
			layout.NewSpacer(),
		),
	)
}
