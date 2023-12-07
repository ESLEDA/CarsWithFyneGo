package models
import (
    
    "math/rand"
    "time"

    "fyne.io/fyne/v2"
    
)

func MoveCircle2() {
	rand.Seed(time.Now().UnixNano()) // s aleatoria 
	vdRandom := rand.Intn(20) + 1       // velocidad aleatoria entre 1 y 20
	for Circle2.Position().X < Finishh-Radiuss {
		time.Sleep(time.Millisecond * time.Duration(vdRandom))     // esperar un tiempo proporcional a la velocidad
		Circle2.Move(Circle2.Position().Add(fyne.NewPos(1, 0))) // mover el círculo un pixel a la derecha
	}
	CheckWinner(2) // comprobar si el círculo del carril 2 ha ganado
}