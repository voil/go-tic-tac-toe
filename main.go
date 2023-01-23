package main

import (
	game "main/src"
	modules "main/src/modules"
)

func main() {
	game.GameInstance.Initialize()
	modules.SplashScreenInstance.Initialize()
	modules.GameplayInstance.Initialize()

	modules.SplashScreenInstance.ShowScene()
	game.GameInstance.StartGame()
}
