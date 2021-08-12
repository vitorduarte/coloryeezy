package main

import (
	"context"
	"image"
	"image/color"
	"log"
	"time"

	"github.com/vitorduarte/coloryeezy/paintimage"
)

func main() {
	ctx := context.Background()
	delay := time.Second * 5
	startTime, err := time.Parse(
		"2006-01-02 15:04:05 -07",
		"2021-08-12 18:19:00 -03")
	if err != nil {
		log.Println(err)
	}

	for range cron(ctx, startTime, delay) {
		err := generateNewYeezyImage("./img/output.png")
		if err != nil {
			log.Println(err)
		}
		log.Println("Created new yeezy")
	}

}

func generateNewYeezyImage(filename string) (err error) {
	painter, err := paintimage.NewPainter("config.json")
	if err != nil {
		return
	}

	outImage, err := painter.Paint()
	if err != nil {
		return
	}

	writeTodayDateOnImage(outImage)
	if err != nil {
		return
	}

	paintimage.SaveImage(outImage, filename)
	return
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
