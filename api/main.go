package main

import (
	"fmt"

	noxgo "github.com/Auric-Trade-Collective/nox-go"
)
import "C"

func main() {
	fmt.Println("Spaunchbop")
}

//export NoxMain
func NoxMain() {
	nox := noxgo.InitNox()
	SetupRoutes(nox)
}
