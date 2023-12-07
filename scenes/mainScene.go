package scenes

import (
	"image/color"
	"myGame4/models"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var (
	line1  *canvas.Line // línea que separa los carriles
	line2  *canvas.Line // línea que marca la salida
	line3  *canvas.Line // línea que marca la meta
	line4  *canvas.Line
)

func SetupGameElements() fyne.CanvasObject {
	models.Circle1 = canvas.NewImageFromFile("assets/carBlackConcurrente.png")
	models.Circle1.Move(fyne.NewPos(models.Startt, models.Heightt/4-models.Radiuss))
	models.Circle1.Resize(fyne.NewSize(models.Radiuss*2, models.Radiuss*2))

	models.Circle2 = canvas.NewImageFromFile("assets/carWhiteConcurrente.png")
	models.Circle2.Move(fyne.NewPos(models.Startt, models.Heightt*3/4-models.Radiuss))
	models.Circle2.Resize(fyne.NewSize(models.Radiuss*2, models.Radiuss*2))

	models.Circle3 = canvas.NewImageFromFile("assets/carCafeConcurrente.png")
	models.Circle3.Move(fyne.NewPos(models.Startt, models.Heightt/2-models.Radiuss))
	models.Circle3.Resize(fyne.NewSize(models.Radiuss*2, models.Radiuss*2))

	models.Btn = widget.NewButton("Iniciar carrera", models.StartRace)
	models.Btn.Resize(fyne.NewSize(150, 50))
	models.Btn.Move(fyne.NewPos(models.Widthh/2-75, models.Heightt+models.Gapp))

	models.Labell = widget.NewLabel("")
	models.Labell.Resize(fyne.NewSize(200, 50))
	models.Labell.Move(fyne.NewPos(models.Widthh/2-100, models.Heightt+models.Gapp+models.Btn.Size().Height))

	line1 = canvas.NewLine(color.RGBA{R: 0, G: 0, B: 0, A: 255})
	line1.StrokeWidth = 5
	line1.Move(fyne.NewPos(0, models.Heightt/3))
	line1.Resize(fyne.NewSize(models.Widthh, 0))

	line2 = canvas.NewLine(color.Black)
	line2.StrokeWidth = 5
	line2.Move(fyne.NewPos(models.Startt, 0))
	line2.Resize(fyne.NewSize(0, models.Heightt))

	line3 = canvas.NewLine(color.Black)
	line3.StrokeWidth = 5
	line3.Move(fyne.NewPos(models.Finishh, 0))
	line3.Resize(fyne.NewSize(0, models.Heightt))

	line4 = canvas.NewLine(color.RGBA{R: 0, G: 0, B: 0, A: 255})
	line4.StrokeWidth = 5
	line4.Move(fyne.NewPos(0, models.Heightt/1.6))
	line4.Resize(fyne.NewSize(models.Widthh, 0))

	return container.NewWithoutLayout(models.Circle1, models.Circle2, models.Circle3, models.Btn, models.Labell, line1, line2, line3, line4)
}
