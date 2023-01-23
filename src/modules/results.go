package modules

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	game "main/src"
)

/**
 * Main struct of splacscreen scene.
 */
type Results struct {
	content *fyne.Container
}

/**
 * @interface {ResultsInterface}
 */
type ResultsInterface interface {
	Initialize()
	ShowScene()
}

/**
 * Constructor of Results struct.
 * @implements ResultsInterface
 */
func CreateResultsScenes() ResultsInterface {
	return &Results{}
}

/**
 * Public method for initialize Results instance.
 * @public
 */
func (r *Results) Initialize() {
	r.__CreateScene()
}

/**
 * Public method show scene.
 * @public
 */
func (r *Results) ShowScene() {
	game.GameInstance.GetWindow().SetContent(r.__CreateScene())
}

/**
 * Private method to create start button.
 * @private
 * @return {*widget.Button}
 */
func (r *Results) __CreateStartButton() *widget.Button {
	button := widget.NewButton("Start new game", func() {
		GameplayInstance.ShowScene()
	})
	button.Resize(fyne.NewSize(150, 50))

	return button
}

/**
 * Private method to create resutl text.
 * @private
 * @return {*canvas.Text}
 */
func (r *Results) __CreateTextResult() *canvas.Text {
	label := ""
	switch {
	case GameplayInstance.GetResult() == "X",
		GameplayInstance.GetResult() == "O":
		label = "The winner is " + GameplayInstance.GetResult()
	default:
		label = "There is a draw"
	}

	text := canvas.NewText(label, color.Black)
	return text
}

/**
 * Private method to create scene container.
 * @private
 * @return {*fyne.Container}
 */
func (r *Results) __CreateScene() *fyne.Container {
	return container.New(
		layout.NewMaxLayout(),
		container.New(
			layout.NewGridLayoutWithRows(2),
			container.New(
				layout.NewCenterLayout(),
				r.__CreateTextResult(),
			),
			container.New(
				layout.NewCenterLayout(),
				r.__CreateStartButton(),
			),
		),
	)
}

var ResultsInstance ResultsInterface = CreateResultsScenes()
