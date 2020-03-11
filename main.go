package main

import (
	"log"

	"github.com/juanjosegongi/automatas-celulares/models"
)

func handleError(err error) {
	log.Fatal(err)
}

func main() {
	rule := 50
	width := 5001
	height := 5001

	universe := models.NewUniverse(rule, width, height)
	for index := 0; index < width; index++ {
		universe.AddCell()
	}

	for index := 0; index < height; index++ {
		universe.Update()
		universe.DrawRow()
	}

	err := universe.Save()
	if err != nil {
		handleError(err)
	}
}
