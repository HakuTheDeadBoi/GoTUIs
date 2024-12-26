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

package graphics

import (
	"errors"

	"github.com/hakuthedeadboi/GoTUIs/pkg/utils"
	"github.com/nsf/termbox-go"
)

const (
	MinWidthFullMode  = 66       // Minimal terminal width
	MinHeightFullMode = 26       // Minimal terminal height
	VLine             = '\u2503' // Boc drawing characters
	HLine             = '\u2501'
	TLCorner          = '\u250F'
	TRCorner          = '\u2513'
	BLCorner          = '\u2517'
	BRCorner          = '\u251B'
	BoxFGColor        = termbox.ColorWhite // Used termbox color constants
	BoxBGColor        = termbox.ColorBlack
	MainBoxLabel      = "GoTUIs"
	GameBoxWidth      = 24 // GameBox width
	GameBoxHeight     = 22 // GameBox height
	GameBoxLabel      = "GAME"
	NextBoxWidth      = 14
	NextBoxHeight     = 8
	NextBoxLabel      = "NEXT"
	LevelBoxWidth     = 14
	LevelBoxHeight    = 5
	LevelBoxLabel     = "LEVEL"
	ControlsBoxWidth  = 24
	ControlsBoxHeight = 18
	ControlsBoxLabel  = "CONTROLS"
)

var ControlsChart = [16]string{
	"UP", "ROTATE",
	"LEFT", "MOVE LEFT",
	"RIGHT", "MOVE RIGHT",
	"DOWN", "FALL",
	"SPACE", "INSTANT DROP",
	"P", "(UN)PAUSE",
	"R", "RESTART",
	"Q/ESC", "QUIT",
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
	scrState := getScreenState()
	if err := checkScreenSize(scrState); err != nil {
		return nil, err
	}

	return scrState, nil
}

func Render(scrState *ScreenState) error {
	// main box
	if err := drawMainBox(scrState); err != nil {
		return err
	}

	// game box
	if err := drawGameBox(scrState); err != nil {
		return err
	}

	// controls box
	if err := drawControlsBox(scrState); err != nil {
		return err
	}

	// next box
	if err := drawNextBox(scrState); err != nil {
		return err
	}

	// level box
	if err := drawLevelBox(scrState); err != nil {
		return err
	}

	if err := termbox.Flush(); err != nil {
		return err
	}

	return nil
}

func checkScreenSize(scrState *ScreenState) error {
	if (*scrState).Width < MinWidthFullMode || (*scrState).Height < MinHeightFullMode {
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
		LevelBoxOffsetRow: ScreenH - ((ScreenH - GameBoxHeight) / 2) - LevelBoxHeight,
		LevelBoxOffsetCol: ((ScreenW - GameBoxWidth) / 2) + GameBoxWidth,
	}
}

func drawBox(row, col, height, width int, label string, scrState *ScreenState) error {
	if len(label) > width-2 {
		return errors.New("Label too long.")
	}

	if row < 0 || col < 0 || row+height-1 >= (*scrState).Height || col+width-1 >= (*scrState).Width {
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

func drawMainBox(scrState *ScreenState) error {
	if err := drawBox(0, 0, (*scrState).Height, (*scrState).Width, MainBoxLabel, scrState); err != nil {
		return err
	}

	return nil
}

func drawGameBox(scrState *ScreenState) error {
	if err := drawBox(scrState.GameBoxOffsetRow, scrState.GameBoxOffsetCol, GameBoxHeight, GameBoxWidth, GameBoxLabel, scrState); err != nil {
		return err
	}

	return nil
}

func drawControlsBox(scrState *ScreenState) error {
	if err := drawBox(scrState.CtrlBoxOffsetRow, scrState.CtrlBoxOffsetCol, ControlsBoxHeight, ControlsBoxWidth, ControlsBoxLabel, scrState); err != nil {
		return err
	}

	ctrlContentOffsetRow := scrState.CtrlBoxOffsetRow + 1
	ctrlContentOffsetCol := scrState.CtrlBoxOffsetCol + 1

	for i := 0; i < len(ControlsChart); i++ {
		var text string
		var err error

		if i%2 == 0 {
			text, err = utils.PadRight(ControlsChart[i], ControlsBoxWidth-2)
		} else {
			text, err = utils.PadLeft(ControlsChart[i], ControlsBoxWidth-2)
		}

		if err != nil {
			return err
		}

		err = drawText(ctrlContentOffsetRow+i, ctrlContentOffsetCol, text, scrState)

		if err != nil {
			return err
		}
	}

	return nil
}

func drawNextBox(scrState *ScreenState) error {
	if err := drawBox(scrState.NextBoxOffsetRow, scrState.NextBoxOffsetCol, NextBoxHeight, NextBoxWidth, NextBoxLabel, scrState); err != nil {
		return err
	}

	return nil
}

func drawLevelBox(scrState *ScreenState) error {
	if err := drawBox(scrState.LevelBoxOffsetRow, scrState.LevelBoxOffsetCol, LevelBoxHeight, LevelBoxWidth, LevelBoxLabel, scrState); err != nil {
		return err
	}

	return nil
}

func drawText(row int, col int, text string, scrState *ScreenState) error {
	if row+len(text) >= (*scrState).Width {
		return errors.New("Text reached outside the box.")
	}

	for i := 0; i < len(text); i++ {
		termbox.SetCell(col+i, row, rune(text[i]), BoxFGColor, BoxBGColor)
	}

	return nil
}
