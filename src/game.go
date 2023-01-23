package src

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

type Player struct {
	Active bool
	Image  string
}

/**
 * Main struct of game.
 */
type Game struct {
	name     string
	instance fyne.App
	window   fyne.Window
	players  map[string]Player
}

/**
 * @interface {GameInterface}
 */
type GameInterface interface {
	StartGame()
	Initialize()
	GetWindow() fyne.Window
}

/**
 * Constructor of Game struct.
 * @implements GameInterface
 */
func CreateGame() GameInterface {
	return &Game{"Tic Tac Toe", nil, nil, nil}
}

/**
 * Public method for initialize game instance.
 * @public
 */
func (g *Game) Initialize() {
	g.players = map[string]Player{}

	g.__CreateApplication()
	g.__CreateWindow()
}

/**
 * Public method to start game and show window.
 * @public
 */
func (g *Game) StartGame() {
	g.window.ShowAndRun()
}

/**
 * Public method to get window instance.
 * @public
 * @return {fyne.Window}
 */
func (g *Game) GetWindow() fyne.Window {
	return g.window
}

/**
 * Private method to create main application instance.
 * @private
 */
func (g *Game) __CreateApplication() {
	g.instance = app.New()
}

/**
 * Private method to create main window instance.
 * @private
 */
func (g *Game) __CreateWindow() {
	g.window = g.instance.NewWindow(g.name)
	g.window.CenterOnScreen()
	g.window.Resize(fyne.NewSize(212, 212))
}

var GameInstance GameInterface = CreateGame()
