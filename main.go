package main

import (
	"myGame4/scenes"
	"fyne.io/fyne/v2/app"
	"myGame4/models"
	"fyne.io/fyne/v2"
)

func main() {
	a := app.New()
	w := a.NewWindow("Carrera de círculos")

	gameElements := scenes.SetupGameElements()

	w.SetContent(gameElements)
	w.Resize(fyne.NewSize(models.Widthh, models.Heightt+models.Gapp+models.Btn.Size().Height+models.Labell.Size().Height))
	w.ShowAndRun()
}
