package models

import (
	"fmt"
	
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

const (
	Widthh  = 600 // ancho del canvas
	Heightt = 200 // alto del canvas
	Radiuss = 20  // radio de los círculos
	Gapp    = 40  // espacio entre los carriles
	Startt  = 50  // posición de la salida
	Finishh = 550 // posición de la meta
)

var (
	Circle1 *canvas.Image  // carro 1
	Circle2 *canvas.Image  // carro 2
	Circle3 *canvas.Image  // carro3 2
	Btn     *widget.Button // botón para iniciar la carrera
	Labell  *widget.Label  // etiqueta para mostrar el ganador
	
)

//la gorutine del main no cuenta
func StartRace() {
	Btn.Disable()      // deshabilitar el botón
	Labell.SetText("") // borrar el texto de la etiqueta
	ResetPositions()   // reiniciar las posiciones de los círculos
	go MoveCircle1()   // mover el carro 1 en una gorutina
	go MoveCircle2()   // mover el carro 2 en otra gorutina
	go MoveCircle3()	// mover el carro 3 en otra gorutina
}

func ResetPositions() {
	Circle1.Move(fyne.NewPos(Startt, Heightt/4-Radiuss))   // posición inicial del carril 1
	Circle2.Move(fyne.NewPos(Startt, Heightt*3/4-Radiuss)) // posición inicial del carril 2
	Circle3.Move(fyne.NewPos(Startt, Heightt/2-Radiuss))   // posición inicial del carril 3
}





//pendiente add a 3rd car
//buscar imagen 
 


func CheckWinner(lane int) {
	if Labell.Text == "" { // si la etiqueta está vacía, significa que nadie ha ganado todavía
		Labell.SetText(fmt.Sprintf("¡El círculo del carril %d ha ganado!", lane)) // mostrar el ganador en la etiqueta
		Btn.Enable()                                                              // habilitar el botón
	}
}
