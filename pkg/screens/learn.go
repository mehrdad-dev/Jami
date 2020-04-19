package screens

import (
	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

// LearnScreen creat scroll screen
func LearnScreen() fyne.CanvasObject {
	image := canvas.NewImageFromFile("/home/mehrdad/Documents/Photoshop/PROJECT/eeteraz3.jpg")
	horiz := widget.NewVScrollContainer(image)

	return fyne.NewContainerWithLayout(layout.NewAdaptiveGridLayout(2),
		fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, nil, nil, nil), horiz),
	)
}
