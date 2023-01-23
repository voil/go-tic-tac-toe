package modules

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	game "main/src"
)

/**
 * Main struct of splacscreen scene.
 */
type SplashScreen struct {
	content *fyne.Container
}

/**
 * @interface {SplashScreenInterface}
 */
type SplashScreenInterface interface {
	Initialize()
	ShowScene()
}

/**
 * Constructor of SplashScreen struct.
 * @implements SplashScreenInterface
 */
func CreateSplashScreenScenes() SplashScreenInterface {
	return &SplashScreen{}
}

/**
 * Public method for initialize splashscreen instance.
 * @public
 */
func (s *SplashScreen) Initialize() {
	s.__CreateScene()
}

/**
 * Public method show scene.
 * @public
 */
func (s *SplashScreen) ShowScene() {
	game.GameInstance.GetWindow().SetContent(s.__CreateScene())
}

/**
 * Private method to create start button.
 * @private
 * @return {*widget.Button}
 */
func (s *SplashScreen) __CreateStartButton() *widget.Button {
	button := widget.NewButton("Start new game", func() {
		GameplayInstance.ShowScene()
	})
	button.Resize(fyne.NewSize(150, 50))

	return button
}

/**
 * Private method to create scene container.
 * @private
 * @return {*fyne.Container}
 */
func (s *SplashScreen) __CreateScene() *fyne.Container {
	return container.New(
		layout.NewMaxLayout(),
		container.New(
			layout.NewCenterLayout(),
			s.__CreateStartButton(),
		),
	)
}

var SplashScreenInstance SplashScreenInterface = CreateSplashScreenScenes()
