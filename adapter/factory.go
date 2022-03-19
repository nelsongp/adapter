package adapter

import (
	"github.com/nelsongp/adapter/auto"
	"github.com/nelsongp/adapter/bici"
)

func Factory(s string) Adapter {
	switch s {
	case "automovil":
		d := auto.Automovil{}
		return &AutomovilAdapter{&d}
	case "bicicleta":
		b := bici.Bicicleta{}
		return &BicicletaAdapter{&b}
	}
	return nil
}
