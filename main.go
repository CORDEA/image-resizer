package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
)

var (
	per = flag.Float64("p", 1.0, "Reduction percentage of the image")
	w   = flag.Int("w", 0, "Width of image")
	h   = flag.Int("h", 0, "Height of image")
)

func getImages(path string) []string {
	images := make([]string, 0)

	stat, err := os.Stat(path)
	if err != nil {
		return images
	}
	switch md := stat.Mode(); {
	case md.IsRegular():
		return append(images, path)
	case md.IsDir():
		files, err := ioutil.ReadDir(path)
		if err != nil {
			return images
		}
		for _, f := range files {
			images = append(images, f.Name())
		}
	}
	return images
}

func main() {
	flag.Parse()

	if flag.NArg() == 0 {
		log.Fatalln("Required argument is missing.")
	}

	path := flag.Arg(0)
	images := getImages(path)
	if err := Resize(images, *per, *w, *h); err != nil {
		log.Fatalln(err)
	}
}
