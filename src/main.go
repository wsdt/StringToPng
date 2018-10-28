package main

import (
	"fmt"
	"time"
	"math/rand"
	"sync"
)

/** © Riedl Kevin, 2018
My first GoLang program (don't like HelloWorld-programs),
so please bare with me :) */


/** CONSTANTS */
const USE_RANDOM_STR = true
// Length of random str generated (takes only effect if USE_RANDOM_STR true)
const RANDOM_STR_LENGTH = 30
// String_to_encode is only used if use_random_str is false!
const STRING_TO_ENCODE = "!ä"
/*Desired name of outputted image (saved in current folder)
Your image(s) will be saved as image_0.png, image_1.png ...*/
const PIC_OUTPUT_FILE = "image"
//Size of outputted image
const PIC_WIDTH = 200
const PIC_HEIGHT = 200
/*Encode desired string in multiple random colors (so structure of all the
additional outputted images will be the same, but the colors are randomized.

This variable should be bigger than 0! */
const ENCODING_VERSIONS = 3

/** Is used instead of STRING_TO_ENCODE in case
of USE_RANDOM_STR = false. */
const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ<>|µ,;.:-_#'+*~^°!§$%&/()=?`´1234567890ß{[]}²³"
func createRandStr(length int) string {
	defer fmt.Println("createRandStr: Generated random string; length =",length)

	rand.Seed(time.Now().Unix())

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

// To wait for other go-routines
var wg sync.WaitGroup
func main() {
	wg.Add(ENCODING_VERSIONS)

	var encodableStr string
	if (USE_RANDOM_STR) {
		encodableStr = createRandStr(RANDOM_STR_LENGTH)
	} else {
		encodableStr = STRING_TO_ENCODE
	}

	fmt.Println("Encoding string '"+encodableStr+"' into image(s).")

	for i:=0; i<ENCODING_VERSIONS;i++ {
		fmt.Println("main: Starting to encode/create new image -> No.",i)
		go saveImg(drawOnImg(encodableStr, createImg()), i)
	}

	//Wait before exiting (to prevent go routines to abort
	wg.Wait()
	fmt.Println("main: All images exported.")
}
