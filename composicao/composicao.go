package main

type esporitvo interface {
	ligarTurbo()
}

type luxuoso interface {
	fazerBaliza()
}

type esportivoLuxuoso interface {
	esporitvo
	luxuoso
}
