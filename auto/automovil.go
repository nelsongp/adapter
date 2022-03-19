package auto

import "fmt"

type Automovil struct {
	Marca     string
	Model     string
	Encendido bool
}

func (a *Automovil) Encender() {
	if a.Encendido {
		fmt.Println("Ya esta encendido")
		return
	}

	fmt.Println("Encendido!")
}

func (a *Automovil) Acelerar() {
	fmt.Println("Acelerando")
}
