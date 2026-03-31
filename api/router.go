package main

import (
	"C"

	"github.com/Auric-Trade-Collective/The-Grid/api/controllers"
	noxgo "github.com/Auric-Trade-Collective/nox-go"
)

func SetupRoutes(nox *noxgo.Nox) {
	core := &controllers.CoreController{}

	nox.CreateGet("/health", core.Health)
}
