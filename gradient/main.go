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
	flag.Usage = func() {
		o := flag.CommandLine.Output()
		fmt.Fprintf(o, "Usage of %s:\n\nCall with 0 - 3 of the following flags, then the filename (png or jpg).\nOutput will be quoted; if additional quotes are needed, escape them:\n\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Fprintf(o, "Examples:\n\n\t./getColors image.png\n\t\t\"#ff1212\"\n\t\t...\n\n\t./getColors -prefix %%{F -suffix } image.jpg\n\t\t\"%%{F#ff1200}\"\n\t\t...\n\n")
	}
	prefix := flag.String("prefix", "", "prefix before #hex color; default is empty string")
	row := flag.Int("row", 0, "row of your image to scan; default is 0")
	suffix := flag.String("suffix", "", "suffix after #hex color; default is empty string\n")
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
		fmt.Println("\"" + *prefix + colors.FromStdColor(img.At(x, *row)).ToHEX().String() + *suffix + "\"")
	}
}
