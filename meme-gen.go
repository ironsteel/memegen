package main

import (
	"bytes"
	"encoding/base64"
	"html/template"
	"image"
	"image/jpeg"
	"log"
	"net/http"
	"strconv"
)

var ImageTemplate string = `<!DOCTYPE html>
<html lang="en"><head></head>
<body><img src="data:image/jpg;base64,{{.Image}}"></body>`

func writeImageWithTemplate(w http.ResponseWriter, img *image.Image) {
	var err error
	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, *img, nil); err != nil {
		w.Write([]byte("unable to encode image"))
		return
	}

	str := base64.StdEncoding.EncodeToString(buffer.Bytes())
	var tmpl *template.Template
	if tmpl, err = template.New("image").Parse(ImageTemplate); err != nil {
		w.Write([]byte("Error parsing template"))
		return
	}

	data := map[string]interface{}{"Image": str}
	if err = tmpl.Execute(w, data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("unable to execute template"))
	}
}

func memeHandler(w http.ResponseWriter, r *http.Request) {
	size := r.URL.Query().Get("size")
	topText := r.URL.Query().Get("top")
	bottomText := r.URL.Query().Get("bottom")
	meme := loadPng("images/yno.png")
	f := loadFont("font/impact.ttf")

	s, err := strconv.ParseFloat(size, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	result, err := drawText(f, topText, bottomText, meme, 75, s)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	var img image.Image = result
	writeImageWithTemplate(w, &img)
}

func main() {
	http.HandleFunc("/meme", memeHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
