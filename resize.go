package main

import (
	"errors"
	"github.com/disintegration/imaging"
	"log"
	"math"
)

const (
	Standard = iota
	Percentage
	Width
	Height
)

type target struct {
	requestType int
	images      []string
	per         float64
	w           int
	h           int
}

func Resize(images []string, per float64, w, h int) error {
	t := &target{Standard, images, per, w, h}
	msg := "-- Resized by the specified Width and Height --"
	if per == 1.0 {
		if w == 0 {
			if h == 0 {
				return errors.New("Required argument is missing.")
			} else {
				t.requestType = Height
				msg = "-- Resized by the specified Height --"
			}
		} else {
			if h == 0 {
				t.requestType = Width
				msg = "-- Resized by the specified Width --"
			}
		}
	} else {
		t.requestType = Percentage
		msg = "-- Resized by the specified Percentage --"
	}
	log.Println(msg)
	t.resize()
	return nil
}

func (t *target) resize() {
	for _, name := range t.images {
		img, err := imaging.Open(name)
		if err == nil {
			size := img.Bounds().Size()
			switch t.requestType {
			case Standard:
				img = imaging.Resize(img, t.w, t.h, imaging.Lanczos)
			case Percentage:
				w := int(math.Ceil(float64(size.X) * t.per))
				img = imaging.Resize(img, w, 0, imaging.Lanczos)
			case Width:
				img = imaging.Resize(img, t.w, 0, imaging.Lanczos)
			case Height:
				img = imaging.Resize(img, 0, t.h, imaging.Lanczos)
			}
			err = imaging.Save(img, name)
		}
		if err != nil {
			log.Println("Failed to '" + name + "' of processing: " + err.Error())
		}
	}
}
