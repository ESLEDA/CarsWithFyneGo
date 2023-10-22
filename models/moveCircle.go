package models

import (
	"fmt"
	"math/rand"
	"time"

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

func MoveCircle1() {
	rand.Seed(time.Now().UnixNano()) // semilla aleatoria 
	speed := rand.Intn(20) + 1       // velocidad aleatoria entre 1 y 20
	for Circle1.Position().X < Finishh-Radiuss {
		time.Sleep(time.Millisecond * time.Duration(speed))     // esperar un tiempo proporcional a la velocidad
		Circle1.Move(Circle1.Position().Add(fyne.NewPos(1, 0))) // mover el círculo un pixel a la derecha
	}
	CheckWinner(1) // comprobar si el círculo del carril 1 ha ganado
	//el tres puedo mandar un str y debo cambiar linea 82 x un str
}

func MoveCircle2() {
	rand.Seed(time.Now().UnixNano()) // s aleatoria 
	vdRandom := rand.Intn(20) + 1       // velocidad aleatoria entre 1 y 20
	for Circle2.Position().X < Finishh-Radiuss {
		time.Sleep(time.Millisecond * time.Duration(vdRandom))     // esperar un tiempo proporcional a la velocidad
		Circle2.Move(Circle2.Position().Add(fyne.NewPos(1, 0))) // mover el círculo un pixel a la derecha
	}
	CheckWinner(2) // comprobar si el círculo del carril 2 ha ganado
}

//pendiente add a 3rd car
//buscar imagen 

func MoveCircle3() {
	rand.Seed(time.Now().UnixNano()) // s aleatoria 
	vdRandom := rand.Intn(20) + 1       // velocidad aleatoria entre 1 y 20
	for Circle3.Position().X < Finishh-Radiuss {
		time.Sleep(time.Millisecond * time.Duration(vdRandom))     // esperar un tiempo proporcional a la velocidad
		Circle3.Move(Circle3.Position().Add(fyne.NewPos(1, 0))) // mover el círculo un pixel a la derecha
	}
	CheckWinner(3) // comprobar si el círculo del carril 3 ha ganado
	//el tres puedo mandar un str y debo cambiar linea 82 x un str
}

func CheckWinner(lane int) {
	if Labell.Text == "" { // si la etiqueta está vacía, significa que nadie ha ganado todavía
		Labell.SetText(fmt.Sprintf("¡El círculo del carril %d ha ganado!", lane)) // mostrar el ganador en la etiqueta
		Btn.Enable()                                                              // habilitar el botón
	}
}
