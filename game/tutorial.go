package game

import (
	"image"
	"image/color"
	"time"

	"code.rocket9labs.com/tslocum/etk"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/leonelquinteros/gotext"
)

type tutorialWidget struct {
	*etk.Frame
	grid      *etk.Grid
	page      int
	lastClick time.Time
}

func NewTutorialWidget() *tutorialWidget {
	w := &tutorialWidget{
		Frame: etk.NewFrame(),
		grid:  etk.NewGrid(),
	}
	w.Frame.SetPositionChildren(true)
	w.Frame.AddChild(w.grid)

	w.setPage(0)

	return w
}

func (w *tutorialWidget) hide() {
	game.lobby.showCreateGame = false
	game.setRoot(listGamesFrame)
	setViewBoard(false)
	game.Board.gameState.PlayerNumber = 0
	w.grid.Clear()
}

func (w *tutorialWidget) dialogText(message string) *tutorialDialog {
	t := etk.NewText(message)
	t.SetBackgroundColor(bufferBackgroundColor)
	return &tutorialDialog{
		Text: t,
		handler: func() {
			w.setPage(w.page + 1)
		},
	}
}

func (w *tutorialWidget) newTutorialBox() *tutorialBox {
	return &tutorialBox{
		Box:     etk.NewBox(),
		handler: w.hide,
	}
}

func (w *tutorialWidget) setPage(page int) {
	if time.Since(w.lastClick) < 250*time.Millisecond {
		return
	}
	w.lastClick = time.Now()
	w.page = page
	w.grid.Clear()

	var title string
	var message string
	switch w.page {
	case 0:
		title = gotext.Get("Tutorial")
		message = gotext.Get("Welcome to the guided tutorial. Click anywhere outside of this message box to close the tutorial. Click anywhere inside of this message box to view the next page.")
	case 1:
		title = gotext.Get("Matches List")
		message = gotext.Get("This screen lists the matches that are currently available. A few bots are always available to play against.")
	case 2:
		game.lobby.showCreateGame = true
		game.setRoot(createGameFrame)
		etk.SetFocus(game.lobby.createGameName)
		title = gotext.Get("Create Match")
		message = gotext.Get("Create a match if you would like to play against someone else. Backgammon and acey-deucey games are supported.")
	case 3:
		game.lobby.showCreateGame = false
		game.setRoot(listGamesFrame)
		game.Board.gameState.PlayerNumber = 1
		if game.needLayoutBoard {
			game.layoutBoard()
		}
		setViewBoard(true)
		title = gotext.Get("Board")
		message = gotext.Get("You have the black checkers. You can move a checker by either clicking it or dragging it.")
	case 4:
		title = gotext.Get("Bearing Off")
		message = gotext.Get("Double click a checker to bear it off. Bear off all 15 checkers to win.")
	case 5:
		title = gotext.Get("Good Luck, Have Fun")
		message = gotext.Get("This concludes the tutorial. Join the community via Matrix/Discord/IRC at %s", "bgammon.org/community")
	case 6:
		w.hide()
		return
	}
	message = title + "\n\n" + message

	w.grid.SetColumnSizes(-1, -1, -1, -1)
	w.grid.SetRowSizes(-1, -1, -1, -1)
	w.grid.AddChildAt(w.newTutorialBox(), 0, 0, 4, 1)
	w.grid.AddChildAt(w.newTutorialBox(), 0, 1, 1, 2)
	w.grid.AddChildAt(w.dialogText(message), 1, 1, 2, 2)
	w.grid.AddChildAt(w.newTutorialBox(), 3, 1, 1, 2)
	w.grid.AddChildAt(w.newTutorialBox(), 0, 3, 4, 1)
}

type tutorialDialog struct {
	*etk.Text
	handler func()
}

func (d *tutorialDialog) Draw(screen *ebiten.Image) error {
	screen.SubImage(d.Text.Rect().Inset(-2)).(*ebiten.Image).Fill(color.RGBA{0, 0, 0, 255})
	return d.Text.Draw(screen)
}

func (d *tutorialDialog) HandleMouse(cursor image.Point, pressed bool, clicked bool) (handled bool, err error) {
	if !cursor.In(d.Rect()) || !clicked {
		return false, nil
	}
	d.handler()
	return true, nil
}

type tutorialBox struct {
	*etk.Box
	handler func()
}

func (b *tutorialBox) HandleMouse(cursor image.Point, pressed bool, clicked bool) (handled bool, err error) {
	if !cursor.In(b.Rect()) || !clicked {
		return false, nil
	}
	b.handler()
	return true, nil
}