package main

import (
	"bufio"
	"context"
	"image"
	"image/color"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/vitorduarte/coloryeezy/paintimage"
)

func main() {
	var availableTweetMessages []string
	twitterClient := createTwitterClient()
	outputImagePath := "./img/output.png"
	ctx := context.Background()
	delay := time.Second * 5
	startTime, err := time.Parse(
		"2006-01-02 15:04:05 -07",
		"2021-08-12 18:19:00 -03")
	if err != nil {
		log.Println(err)
	}

	rand.Seed(time.Now().UnixNano())
	for range cron(ctx, startTime, delay) {
		err := generateNewYeezyImage(outputImagePath)
		if err != nil {
			log.Println(err)
		}

		err = twitterClient.PostImage(outputImagePath, getTweetMessage(&availableTweetMessages))
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

func getTweetMessage(availableTweetMessages *[]string) (tweetMessage string) {
	if len(*availableTweetMessages) == 0 {
		var err error
		*availableTweetMessages, err = openFileAsSliceOfRows("tweets.txt")
		if err != nil {
			return
		}
	}

	messageIndex := rand.Intn(len(*availableTweetMessages))
	tweetMessage = (*availableTweetMessages)[messageIndex]
	*availableTweetMessages = append((*availableTweetMessages)[:messageIndex], (*availableTweetMessages)[messageIndex+1:]...)
	return
}

// Function to read a file and return list for each line
func openFileAsSliceOfRows(filename string) (rows []string, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rows = append(rows, scanner.Text())
	}
	return
}
