package models

import (
    
    "math/rand"
    "time"

    "fyne.io/fyne/v2"
    
)

func MoveCircle1() {
    rand.Seed(time.Now().UnixNano()) // semilla aleatoria 
    speed := rand.Intn(20) + 1       // velocidad aleatoria entre 1 y 20
    for Circle1.Position().X < Finishh-Radiuss {
        time.Sleep(time.Millisecond * time.Duration(speed))     // esperar un tiempo proporcional a la velocidad
        Circle1.Move(Circle1.Position().Add(fyne.NewPos(1, 0))) // mover el círculo un pixel a la derecha
    }
    CheckWinner(1) // comprobar si el círculo del carril 1 ha ganado
}
