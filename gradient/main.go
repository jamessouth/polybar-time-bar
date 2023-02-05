package main

import (
	"bytes"
	"flag"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"

	"github.com/go-playground/colors"
)

func main() {
	prefix := flag.String("prefix", "", "prefix before #hex color. escape quotes")
	row := flag.Int("row", 0, "image row to use")
	suffix := flag.String("suffix", "", "suffix after #hex color")
	flag.Parse()

	imgfile := flag.Arg(0)
	if imgfile == "" {
		log.Fatal("no img file given")
	}

	data, err := os.ReadFile(imgfile)
	if err != nil {
		log.Fatal(err)
	}

	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		log.Fatal(err)
	}

	bnds := img.Bounds()

	for x := bnds.Min.X; x < bnds.Max.X; x++ {
		fmt.Println(*prefix + colors.FromStdColor(img.At(x, *row)).ToHEX().String() + *suffix)
	}
}
