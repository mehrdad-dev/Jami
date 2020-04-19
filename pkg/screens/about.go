package screens

import (
	"net/url"

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
	logo := canvas.NewImageFromFile("/home/mehrdad/Pictures/piano.png")
	logo.SetMinSize(fyne.NewSize(212, 212))

	return widget.NewVBox(
		widget.NewHBox(layout.NewSpacer(), logo, layout.NewSpacer()),
		widget.NewLabelWithStyle("Giano", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		widget.NewLabelWithStyle("A Piano Keyboard Writen In GO", fyne.TextAlignCenter, fyne.TextStyle{Bold: false}),

		widget.NewHBox(layout.NewSpacer(),
			widget.NewHyperlink("github", parseURL("https://github.com/mehrdad-dev")),
			widget.NewLabel("-"),
			widget.NewHyperlink("conntact", parseURL("http://www.mehrdad-dev.ir/contact-me/")),
			layout.NewSpacer(),
		),
		widget.NewLabelWithStyle("Giano v1.1", fyne.TextAlignCenter, fyne.TextStyle{Bold: false}),
	)
}
