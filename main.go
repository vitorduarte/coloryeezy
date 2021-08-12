package main

import (
	"fmt"
	"image"
	"image/color"
	"time"

	"github.com/vitorduarte/coloryeezy/paintimage"
)

func main() {
	painter, err := paintimage.NewPainter("config.json")
	if err != nil {
		fmt.Println(err)
	}

	outImage, err := painter.Paint()
	if err != nil {
		fmt.Println(err)
	}

	writeTodayDateOnImage(outImage)
	if err != nil {
		fmt.Println(err)
	}

	paintimage.SaveImage(outImage, "output.png")
}

func writeTodayDateOnImage(canva *image.RGBA) (err error) {
	textLeft := paintimage.TextConfig{
		FontPath: "./fonts/font.ttf",
		Text:     getTodayDate(),
		Size:     63,
		Color:    color.RGBA{0, 0, 0, 255},
		Location: image.Pt(58, 1125),
	}

	err = paintimage.WriteTextOnImage(canva, textLeft)
	if err != nil {
		return
	}

	textRight := textLeft
	textRight.Location.X = 901
	err = paintimage.WriteTextOnImage(canva, textRight)
	if err != nil {
		return
	}

	return
}

func getTodayDate() string {
	return time.Now().Format("01.02")
}
