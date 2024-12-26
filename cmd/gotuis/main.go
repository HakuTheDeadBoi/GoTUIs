package main

import (
	"time"

	"github.com/hakuthedeadboi/GoTUIs/pkg/graphics"
	"github.com/nsf/termbox-go"
)

func main() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()

	sstate, err := graphics.InitScreenState()
	if err != nil {
		panic(err)
	}

	if err := graphics.Render(sstate); err != nil {
		panic(err)
	}

	time.Sleep(time.Second * 10)
}
