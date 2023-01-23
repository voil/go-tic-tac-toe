package modules

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	game "main/src"
	widgets "main/src/widgets"
)

type TableElement struct {
	Taken     bool
	Type      string
	Container *fyne.Container
}

/**
 * Main struct of gameplay scene.
 */
type Gameplay struct {
	content *fyne.Container
	turn    string
	result  string
	table   map[string]TableElement
}

/**
 * @interface {GameplayInterface}
 */
type GameplayInterface interface {
	Initialize()
	ShowScene()
	GetResult() string
}

/**
 * Constructor of Gameplay struct.
 * @implements GameplayInterface
 */
func CreateGameplayScenes() GameplayInterface {
	return &Gameplay{
		turn: "o",
	}
}

/**
 * Public method for initialize Gameplay instance.
 * @public
 */
func (gp *Gameplay) Initialize() {
	gp.table = map[string]TableElement{}
	gp.__CreateScene()
}

/**
 * Public method show scene.
 * @public
 */
func (gp *Gameplay) ShowScene() {
	game.GameInstance.GetWindow().SetContent(gp.__CreateScene())
}

/**
 * Public method to get result of play.
 * @public
 * @return {string}
 */
func (gp *Gameplay) GetResult() string {
	return gp.result
}

/**
 * Private method for create table board.
 * @private
 */
func (gp *Gameplay) __CreateTableBoard() {
	table := [9]string{
		"top-left", "top", "top-right",
		"middle-left", "middle", "middle-right",
		"bottom-left", "bottom", "bottom-right",
	}

	for _, element := range table {
		var key string = element
		gp.table[key] = TableElement{
			Taken: false,
			Container: container.New(
				layout.NewGridLayoutWithColumns(1),
				container.NewWithoutLayout(
					widgets.NewImageButton("", func(this *widgets.ImageButton, image *canvas.Image) {
						gp.__handleClickElementBoard(key, image)
					}, fyne.NewSize(64, 64)),
				),
			),
		}
	}
}

/**
 * Private method for handle click on element on board.
 * @param {string} key
 * @param {*canvas.Image} image
 * @private
 */
func (gp *Gameplay) __handleClickElementBoard(key string, image *canvas.Image) {
	entry := gp.table[key]
	if !gp.table[key].Taken {
		entry.Taken = true
		entry.Type = gp.turn

		image.File = "./assets/" + gp.turn + ".jpg"
		image.Refresh()

		gp.table[key] = entry

		gp.__ChangeTurn()
		gp.__CheckIsEndGame()
	}
}

/**
 * Private method for change turn of player.
 * @private
 */
func (gp *Gameplay) __ChangeTurn() {
	if gp.turn == "o" {
		gp.turn = "x"
	} else {
		gp.turn = "o"
	}
}

/**
 * Private method for check is game end.
 * @private
 */
func (gp *Gameplay) __CheckIsEndGame() {
	table := [9]string{
		"top-left", "top", "top-right",
		"middle-left", "middle", "middle-right",
		"bottom-left", "bottom", "bottom-right",
	}

	var countTakenElements int = 0
	for _, element := range table {
		gp.__CheckIsPlayerWinner(element, "x")
		gp.__CheckIsPlayerWinner(element, "o")
		countTakenElements = gp.__CheckIsDraw(countTakenElements, element)

	}
}

/**
 * Private method for check is plarye winnier.
 * @param {string} element
 * @param {string} typePlayer
 * @private
 */
func (gp *Gameplay) __CheckIsPlayerWinner(element string, typePlayer string) {

	arrayOfSolutions := map[string][][3]string{
		"top-left": {
			{"top-left", "top", "top-right"},
			{"top-left", "middle-left", "bottom-left"},
			{"top-left", "middle", "bottom-right"},
		},
		"top": {
			{"top", "middle", "bottom"},
		},
		"top-right": {
			{"top-right", "middle-right", "bottom-right"},
		},
		"middle-left": {
			{"middle-left", "middle", "middle-right"},
		},
		"bottom-left": {
			{"bottom-left", "bottom", "bottom-right"},
			{"bottom-left", "middle", "top-right"},
		},
	}

	if arrayOfSolutions[element] == nil {
		return
	}

	for _, value := range arrayOfSolutions[element] {
		if gp.table[value[0]].Type == typePlayer &&
			gp.table[value[1]].Type == typePlayer &&
			gp.table[value[2]].Type == typePlayer {
			gp.result = strings.ToUpper(typePlayer)
			ResultsInstance.ShowScene()
		}
	}
}

/**
 * Private method for check is plarye winnier.
 * @param {int} result
 * @param {string} element
 * @return {int}
 * @private
 */
func (gp *Gameplay) __CheckIsDraw(result int, element string) int {
	if gp.table[element].Taken {
		result += 1
		if result == 9 {
			gp.result = "draw"
			ResultsInstance.ShowScene()
			return result
		}
	}

	return result
}

/**
 * Private method to create scene container.
 * @private
 * @return {*fyne.Container}
 */
func (gp *Gameplay) __CreateScene() *fyne.Container {
	gp.__CreateTableBoard()

	button := widget.NewButton("", func() {})
	button.Resize(fyne.NewSize(64, 64))

	return container.New(
		layout.NewMaxLayout(),
		container.New(
			layout.NewGridLayoutWithRows(3),
			container.New(
				layout.NewGridLayoutWithColumns(3),
				gp.table["top-left"].Container,
				gp.table["top"].Container,
				gp.table["top-right"].Container,
			),
			container.New(
				layout.NewGridLayoutWithColumns(3),
				gp.table["middle-left"].Container,
				gp.table["middle"].Container,
				gp.table["middle-right"].Container,
			),
			container.New(
				layout.NewGridLayoutWithColumns(3),
				gp.table["bottom-left"].Container,
				gp.table["bottom"].Container,
				gp.table["bottom-right"].Container,
			),
		),
	)
}

var GameplayInstance GameplayInterface = CreateGameplayScenes()
