package main

import (
	"bufio"
	"flag"
	"fmt"
	"image/png"
	"log"
	"os"
)

func main() {
	flag.Parse()
	meme := loadPng("images/yno.png")
	f := loadFont("font/impact.ttf")

	result, err := drawText(f, "HELLO", "WORLD", meme, 75, 48)
	if err != nil {
		log.Fatal(err)
	}

	// Save that RGBA image to disk.
	outFile, err := os.Create("out.png")
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()
	b := bufio.NewWriter(outFile)
	err = png.Encode(b, result)
	if err != nil {
		log.Fatal(err)
	}
	err = b.Flush()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Wrote out.png OK.")
}
