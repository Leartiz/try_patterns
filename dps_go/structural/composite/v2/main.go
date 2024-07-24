package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

// -----------------------------------------------------------------------

// Composite!
type GameMap struct {
	rects []*canvas.Rectangle // <--- Leaf
}

func NewGameMap(sideRectCount int) (*GameMap, error) {
	if sideRectCount < 0 {
		return nil,
			fmt.Errorf("side rect count less than zero")
	}

	rects := make([]*canvas.Rectangle, 0,
		sideRectCount*sideRectCount)

	for i := 0; i < sideRectCount; i++ {
		rects = append(rects,
			canvas.NewRectangle(color.Black))
	}

	return &GameMap{
		rects: rects,
	}, nil
}

// CanvasObject
// -----------------------------------------------------------------------

func (gm *GameMap) MinSize() fyne.Size {
	return gm.rects[0].MinSize()
}

func (gm *GameMap) Move(position fyne.Position) {}

func (gm *GameMap) Position() fyne.Position {
	return fyne.Position{}
}

// -----------------------------------------------------------------------

func (gm *GameMap) Resize(size fyne.Size) {}

func (gm *GameMap) Size() fyne.Size {
	return fyne.Size{}
}

func (gm *GameMap) Hide() {
	for i := 0; i < len(gm.rects); i++ {
		gm.rects[i].Hide()
	}
}

func (gm *GameMap) Visible() bool {
	return gm.rects[0].Visible()
}

func (gm *GameMap) Show() {
	for i := 0; i < len(gm.rects); i++ {
		gm.rects[i].Show()
	}
}

func (gm *GameMap) Refresh() {
	for i := 0; i < len(gm.rects); i++ {
		gm.rects[i].Refresh()
	}
}

// main
// -----------------------------------------------------------------------

func main() {

	a := app.New()
	window := a.NewWindow("Grid!")

	// ***

	window.Resize(fyne.Size{
		Width:  512,
		Height: 512,
	})

	gridSideSize := 10
	grid := container.New(layout.NewGridLayout(gridSideSize))
	fmt.Printf("Grid: %v", len(grid.Objects))

	for i := 0; i < gridSideSize*gridSideSize; i++ {
		rect := canvas.NewRectangle(color.Black)
		grid.Add(rect)
	}

	// ***

	window.SetContent(grid)
	window.ShowAndRun()
}

// -----------------------------------------------------------------------
