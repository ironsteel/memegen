package main

import (
	"image"
	"image/png"
	"io/ioutil"
	"log"
	"os"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
)

func loadPng(filename string) image.Image {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer f.Close()
	img, err := png.Decode(f)
	if err != nil {
		log.Fatalf(err.Error())
	}
	return img
}

func loadFont(filename string) *truetype.Font {
	// Read the font data.
	fontBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	f, err := freetype.ParseFont(fontBytes)
	if err != nil {
		log.Fatal(err)
	}
	return f
}
