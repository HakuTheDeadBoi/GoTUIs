// Package:		GoTUIs
// File:		graphics.go
// Description:	provides tools to render the screen
// Author:		Haku the Dead Boi
// Date:		2024-12-24
// Version:		unfinished
// License:		Do whatever you want licence
//
// Notes:
//				Haku, remember that you decided tp
//				represent block two chars wide and on char high

package GoTUIs

import (
	"errors"
	"github.com/nsf/termbox-go"
)

const (
	MinWidthFullMode  = 45       // Minimal terminal width
	MinHeightFullMode = 26       // Minimal terminal height
	VLine             = '\u2503' // Boc drawing characters
	HLine             = '\u2501'
	TLCorner          = '\u250F'
	TRCorner          = '\u2513'
	BLCorner          = '\u2517'
	BRCorner          = '\u251B'
	BoxFGColor        = termbox.ColorWhite // Used termbox color constants
	BoxBGColor        = termbox.ColorBlack
	GameBoxWidth      = 24 // GameBox width
	GameBoxHeight     = 22 // GameBox height
	NextBoxWidth      = 14
	NextBoxHeight     = 8
	LevelBoxWidth     = 14
	LevelBoxHeight    = 5
	ControlsBoxWidth  = 14
	ControlsBoxHeight = 18
)

var ControlsChart = map[string]string{
	"UP":    "ROTATE",
	"LEFT":  "MOVE LEFT",
	"RIGHT": "MOVE RIGHT",
	"DOWN":  "FALL",
	"SPACE": "INSTANT DROP",
	"P":     "(UN)PAUSE",
	"R":     "RESTART",
	"Q/ESC": "QUIT",
}

type ScreenState struct {
	Width             int
	Height            int
	GameBoxOffsetRow  int
	GameBoxOffsetCol  int
	CtrlBoxOffsetRow  int
	CtrlBoxOffsetCol  int
	NextBoxOffsetRow  int
	NextBoxOffsetCol  int
	LevelBoxOffsetRow int
	LevelBoxOffsetCol int
}

func InitScreenState() (*ScreenState, error) {
	sstate := getScreenState()
	if err := checkScreenSize(sstate); err != nil {
		return nil, err
	}

	return sstate, nil
}

func Render(sstate *ScreenState) error {
	// main box
	if err := drawBox(0, 0, (*sstate).Height, (*sstate).Width, "GoTUIs", sstate); err != nil {
		return err
	}

	// game box
	if err := drawBox(sstate.GameBoxOffsetRow, sstate.GameBoxOffsetCol, GameBoxHeight, GameBoxWidth, "Game", sstate); err != nil {
		return err
	}

	// controls box
	if err := drawBox(sstate.CtrlBoxOffsetRow, sstate.CtrlBoxOffsetCol, ControlsBoxHeight, ControlsBoxWidth, "Controls", sstate); err != nil {
		return err
	}

	// next box
	if err := drawBox(sstate.NextBoxOffsetRow, sstate.NextBoxOffsetCol, NextBoxHeight, NextBoxWidth, "Next", sstate); err != nil {
		return err
	}

	// level box
	if err := drawBox(sstate.LevelBoxOffsetRow, sstate.LevelBoxOffsetCol, LevelBoxHeight, LevelBoxWidth, "Level", sstate); err != nil {
		return err
	}

	if err := termbox.Flush(); err != nil {
		return err
	}

	return nil
}

func checkScreenSize(sstate *ScreenState) error {
	if (*sstate).Width < MinWidthFullMode || (*sstate).Height < MinHeightFullMode {
		return errors.New("Width and Height must be greater than or equal to MinWidth and MinHeight")
	}

	return nil
}

func getScreenState() *ScreenState {
	ScreenW, ScreenH := termbox.Size()
	return &ScreenState{
		Width:             ScreenW,
		Height:            ScreenH,
		GameBoxOffsetRow:  (ScreenH - GameBoxHeight) / 2,
		GameBoxOffsetCol:  (ScreenW - GameBoxWidth) / 2,
		CtrlBoxOffsetRow:  (ScreenH - GameBoxHeight) / 2,
		CtrlBoxOffsetCol:  ((ScreenW - GameBoxWidth) / 2) - ControlsBoxWidth,
		NextBoxOffsetRow:  (ScreenH - GameBoxHeight) / 2,
		NextBoxOffsetCol:  ((ScreenW - GameBoxWidth) / 2) + GameBoxWidth,
		LevelBoxOffsetRow: ScreenH - ((ScreenH-GameBoxHeight)/2 + 1) - LevelBoxHeight,
		LevelBoxOffsetCol: ((ScreenW - GameBoxWidth) / 2) + GameBoxWidth,
	}
}

func drawBox(row, col, height, width int, label string, sstate *ScreenState) error {
	if len(label) > width-2 {
		return errors.New("Label too long.")
	}

	if row < 0 || col < 0 || row+height-1 >= (*sstate).Height || col+width-1 >= (*sstate).Width {
		return errors.New("Box overflows the screen.")
	}

	for x := 0; x < width; x++ {
		if x == 0 {
			termbox.SetCell(col, row, TLCorner, BoxFGColor, BoxBGColor)
			termbox.SetCell(col, row+height-1, BLCorner, BoxFGColor, BoxBGColor)
		} else if x == width-1 {
			termbox.SetCell(col+x, row, TRCorner, BoxFGColor, BoxBGColor)
			termbox.SetCell(col+x, row+height-1, BRCorner, BoxFGColor, BoxBGColor)
		} else if x >= 1 && x <= len(label) {
			termbox.SetCell(col+x, row, rune(label[x-1]), BoxFGColor, BoxBGColor)
			termbox.SetCell(col+x, row+height-1, HLine, BoxFGColor, BoxBGColor)
		} else {
			termbox.SetCell(col+x, row, HLine, BoxFGColor, BoxBGColor)
			termbox.SetCell(col+x, row+height-1, HLine, BoxFGColor, BoxBGColor)
		}
	}

	for y := 1; y < height-1; y++ {
		termbox.SetCell(col, row+y, VLine, BoxFGColor, BoxBGColor)
		termbox.SetCell(col+width-1, row+y, VLine, BoxFGColor, BoxBGColor)
	}

	return nil
}
