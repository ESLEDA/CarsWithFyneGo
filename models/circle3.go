package models

import (
    
    "math/rand"
    "time"

    "fyne.io/fyne/v2"
    
)

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