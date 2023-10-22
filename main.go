package main

import ( 
	
	"image/color"
	"myGame4/models"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var (
	line1   *canvas.Line   // línea que separa los carriles
	line2   *canvas.Line   // línea que marca la salida
	line3   *canvas.Line   // línea que marca la meta
	line4  *canvas.Line
)

func main() {
	a := app.New()
	w := a.NewWindow("Carrera de círculos")

	models.Circle1 = canvas.NewImageFromFile("assets/carBlackConcurrente.png") //  imagen
	models.Circle1.Move(fyne.NewPos(models.Startt, models.Heightt/4-models.Radiuss))		// posición inicial del carril 1
	models.Circle1.Resize(fyne.NewSize(models.Radiuss*2, models.Radiuss*2))	// tamaño del círculo
	
	models.Circle2 = canvas.NewImageFromFile("assets/carWhiteConcurrente.png") //  imagen
	models.Circle2.Move(fyne.NewPos(models.Startt, models.Heightt*3/4-models.Radiuss))		// posición inicial del carril 1
	models.Circle2.Resize(fyne.NewSize(models.Radiuss*2, models.Radiuss*2))	// tamaño del círculo
	//pendiente add a 3rd car, llamar de moveCircle, search image car
	
	models.Circle3 = canvas.NewImageFromFile("assets/carCafeConcurrente.png") //  imagen
	models.Circle3.Move(fyne.NewPos(models.Startt, models.Heightt/2-models.Radiuss))
	models.Circle3.Resize(fyne.NewSize(models.Radiuss*2, models.Radiuss*2))

	models.Btn = widget.NewButton("Iniciar carrera", models.StartRace) // botón para iniciar la carrera
	models.Btn.Resize(fyne.NewSize(150, 50))                     // tamaño del botón
	models.Btn.Move(fyne.NewPos(models.Widthh/2-75, models.Heightt+models.Gapp))         // posición del botón

	models.Labell = widget.NewLabel("") // etiqueta para mostrar el ganador
	models.Labell.Resize(fyne.NewSize(200, 50))                     // tamaño de la etiqueta
	models.Labell.Move(fyne.NewPos(models.Widthh/2-100, models.Heightt+models.Gapp+models.Btn.Size().Height))         // posición de la etiqueta

	//pendiente add lines start and finish
	line1 = canvas.NewLine(color.RGBA{R:0,G:0,B:0,A:255})          // línea separa carril blanca
	line1.StrokeWidth = 5                                   // grosor de la línea
	line1.Move(fyne.NewPos(0, models.Heightt/3))                    // posición de la línea
	line1.Resize(fyne.NewSize(models.Widthh, 0))                    // tamaño de la línea

	line2 = canvas.NewLine(color.Black)          // línea negra inicial
	line2.StrokeWidth = 5                                   // grosor de la línea
	line2.Move(fyne.NewPos(models.Startt, 0))                       // posición de la línea
	line2.Resize(fyne.NewSize(0, models.Heightt))                   // tamaño de la línea

	line3 = canvas.NewLine(color.Black)          // línea negra final
	line3.StrokeWidth = 5                                   // grosor de la línea
	line3.Move(fyne.NewPos(models.Finishh, 0))                      // posición de la línea
	line3.Resize(fyne.NewSize(0, models.Heightt))                   // tamaño de la línea

	line4 = canvas.NewLine(color.RGBA{R:0,G:0,B:0,A:255})          // línea separa carril blanca
	line4.StrokeWidth = 5                                   // grosor de la línea
	line4.Move(fyne.NewPos(0, models.Heightt/1.6))                    // posición de la línea
	line4.Resize(fyne.NewSize(models.Widthh, 0))                    // tamaño de la línea

	w.SetContent(container.NewWithoutLayout(models.Circle1, models.Circle2, models.Circle3 , models.Btn, models.Labell, line1, line2, line3, line4)) // contenedor sin layout para posicionar los elementos 

	w.Resize(fyne.NewSize(models.Widthh, models.Heightt+models.Gapp+models.Btn.Size().Height+models.Labell.Size().Height)) // tamaño de la ventana
	w.ShowAndRun()                                                                     // mostrar y ejecutar la ventana
}
