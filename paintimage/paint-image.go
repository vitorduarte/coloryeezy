package paintimage

import (
	"encoding/json"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

type Painter struct {
	Config PaintConfig
}

type PaintConfig struct {
	Template string   `json:"template"`
	Masks    []string `json:"masks"`
}

func NewPainter(configPath string) (p Painter, err error) {
	config, err := ReadConfigFile(configPath)
	if err != nil {
		fmt.Println(err)
	}

	return Painter{Config: config}, nil
}

func ReadConfigFile(path string) (config PaintConfig, err error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		return
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		return
	}

	return
}

func (p *Painter) Paint(outputPath string) {
	rand.Seed(time.Now().UnixNano())
	canva, err := openTemplateImage(p.Config.Template)
	if err != nil {
		fmt.Println("Could not open template image", err)
		return
	}

	for _, mask := range p.Config.Masks {
		err = fillImageWithColorWithMask(canva, mask, generateRandomColor())
		if err != nil {
			fmt.Println("Could not fill image")
			return
		}
	}

	saveImage(canva, outputPath)

}

func openTemplateImage(path string) (canva *image.RGBA, err error) {
	templateImage, err := openImage(path)
	if err != nil {
		return
	}

	canva = image.NewRGBA(templateImage.Bounds())
	draw.Draw(canva, canva.Bounds(), templateImage, image.ZP, draw.Src)
	return
}

func openImage(path string) (image.Image, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func fillImageWithColorWithMask(canva *image.RGBA, maskPath string, color color.RGBA) (err error) {
	fill := &image.Uniform{color}

	mask, err := openImage(maskPath)
	if err != nil {
		return
	}

	draw.DrawMask(canva, canva.Bounds(), fill, image.ZP, mask, image.ZP, draw.Over)
	return
}

func generateRandomColor() color.RGBA {
	return color.RGBA{uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), 255}
}

func saveImage(img image.Image, outputPath string) (err error) {
	f, err := os.Create(outputPath)
	if err != nil {
		return
	}
	defer f.Close()

	err = png.Encode(f, img)
	if err != nil {
		return
	}

	return nil
}
