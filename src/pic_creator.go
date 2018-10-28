package main

/** © Riedl Kevin, 2018
My first GoLang program (don't like HelloWorld-programs),
so please bare with me :) */

import (
	"image"
	"image/color"
	"os"
	"image/png"
	"math/rand"
	"time"
	"strconv"
)

func strToUnicode(str string) []rune {
	return []rune(str)
}

var addSeed int64 = 0
func getRandNo() uint8 {
	addSeed++
	//Additional seed, so that even at same timestamp different color
	rand.Seed(time.Now().Unix()+addSeed)
	return uint8(rand.Intn(255))
}

func getRandColor() color.RGBA {
	/* Generate random no between 0-255 4 times to have a random color.*/
	return color.RGBA{getRandNo(),getRandNo(),getRandNo(),getRandNo()}
}

func createImg() *image.RGBA {
	upLeft := image.Point{0,0}
	lowRight := image.Point{PIC_WIDTH, PIC_HEIGHT}

	return image.NewRGBA(image.Rectangle{upLeft, lowRight})
}

func drawOnImg(strToEncode string, img *image.RGBA) (*image.RGBA) {
	//cyan := color.RGBA{100,200,200,0xff}

	strI := 0
	var currColor color.RGBA = getRandColor()
	var currLetter rune = rune(strToEncode[strI]) //get first letter
	for x := 0; x < PIC_WIDTH; x++ {
		for y := 0; y < PIC_HEIGHT; y++ {
			if (currLetter <= 0) {
				//get new color when currLeter has expired
				currColor = getRandColor()

				if strI+1 >= len(strToEncode) {
					strI = 0
				} else {
					strI++
				}
				currLetter = rune(strToEncode[strI])
			}
			img.Set(x,y,currColor)
			currLetter--
		}
	}
	return img
}

func saveImg(img *image.RGBA, index int) {
	f, _ := os.Create(PIC_OUTPUT_FILE+"_"+strconv.Itoa(index)+".png")
	png.Encode(f, img)

	// Indicate afterwards that pic has been exported
	wg.Done()
}
