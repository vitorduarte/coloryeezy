package main

import (
	"fmt"

	"github.com/vitorduarte/coloryeezy/paintimage"
)

func main() {
	painter, err := paintimage.NewPainter("config.json")
	if err != nil {
		fmt.Println(err)
	}
	painter.Paint("output.png")
}
