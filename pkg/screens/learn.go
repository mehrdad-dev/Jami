package screens

import (
	"fmt"
	"path/filepath"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

// LearnScreen creat scroll screen
func LearnScreen() fyne.CanvasObject {
	filePath, err := filepath.Abs("../Jami/assets/keyboard.jpg")
	if err != nil {
		fmt.Println("ERR in read keyboard.jpg ", err)
	}

	image := canvas.NewImageFromFile(filePath)
	// image.SetMinSize(fyne.NewSize(590, 212))

	horiz := widget.NewHScrollContainer(image)

	return fyne.NewContainerWithLayout(layout.NewAdaptiveGridLayout(1),
		fyne.NewContainerWithLayout(layout.NewAdaptiveGridLayout(1), horiz),
	)
}
