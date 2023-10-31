package main

import (
	"challenge/application/rest"
	"challenge/internal/dpi"
	infrastructures "challenge/internal/infrastructure"
)

func main() {
	env := infrastructures.NewEnv()
	env.LoadEnv()
	dpi.SetupDPI(env)
	rest.Setup(env)
}
